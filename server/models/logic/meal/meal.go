package meal

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"fmt"
	"io/ioutil"
	"mime/multipart"
	"os"
	"path/filepath"
	"server/helpers"
	"strconv"
	"strings"
	"time"

	"github.com/astaxie/beego"

	"server/helpers/constant"
	structAPI "server/structs/api"
	structLogic "server/structs/logic"

	"github.com/dustin/go-humanize"
	"github.com/jung-kurt/gofpdf"
)

// GetMealRequestByID ...
func GetMealRequestByID(ID int64) (structLogic.GetMealRequest, error) {
	respGet, errGet := DBMeal.GetMealRequestByID(ID)
	if errGet != nil {
		helpers.CheckErr("Error get pending request @GetMealRequestByID - logicMeal", errGet)
	}

	return respGet, errGet
}

// PostMealRequest ...
func PostMealRequest(employeeNumber int64, req *structAPI.ReqMeal) (resp structAPI.RespMeal, err error) {
	user, errGetUser := DBUser.GetUserByEmployeeNumber(employeeNumber)
	if errGetUser != nil {
		helpers.CheckErr("Error get user @PostMealRequest - logicMeal", errGetUser)
		return resp, errGetUser
	}

	//Check if User is Supervisor
	if user.SupervisorID == 0 || user.SupervisorID == user.EmployeeNumber {
		user.SupervisorID = user.EmployeeNumber
	}
	supervisor, errGetSpv := DBUser.GetUserByEmployeeNumber(user.SupervisorID)
	if errGetSpv != nil {
		helpers.CheckErr("Error get supervisor @PostMealRequest - logicMeal", errGetSpv)
		return resp, errGetSpv
	}

	//parse request date
	reqDate, errParse := time.Parse(constant.FormatDateTime, req.RequestDate)
	if errParse != nil {
		helpers.CheckErr("Error parse time request date @PostMealRequest - logicMeal", errParse)
		return resp, errParse
	}

	insertDate := time.Now()
	insertDateFormatted := insertDate.Format(constant.FormatDateTime)
	reqStatus := constant.StatusPendingApproval

	ID, errInsert := DBMeal.PostMealRequest(
		user.EmployeeNumber,
		req.OtherRequestorID,
		supervisor.EmployeeNumber,
		req.Amount,
		req.BriefDescription,
		req.Notes,
		req.ReceiptUploadPath,
		reqDate,
		reqStatus,
		insertDate,
		insertDate,
	)
	if errInsert != nil {
		helpers.CheckErr("Error insert meal request @PostMealRequest - logicMeal", errInsert)
		return resp, errInsert
	}

	resp.ID = ID
	resp.RequestorID = user.EmployeeNumber
	resp.OtherRequestorID = req.OtherRequestorID
	resp.SupervisorID = supervisor.EmployeeNumber
	resp.Amount = req.Amount
	resp.BriefDescription = req.BriefDescription
	resp.Notes = req.Notes
	resp.ReceiptUploadPath = req.ReceiptUploadPath
	resp.RequestDate = req.RequestDate
	resp.Status = reqStatus
	resp.CreatedAt = insertDateFormatted
	resp.UpdatedAt = insertDateFormatted

	return resp, nil
}

// SendEmailToSupervisor ...
func SendEmailToSupervisor(respMeal structAPI.RespMeal) error {
	requestor, errGetEmployee := DBUser.GetUserByEmployeeNumber(respMeal.RequestorID)
	if errGetEmployee != nil {
		helpers.CheckErr("Failed get employee @SendEmailToSupervisor - logicMeal", errGetEmployee)
		return errGetEmployee
	}

	supervisor, errGetSupervisor := DBUser.GetUserByEmployeeNumber(respMeal.SupervisorID)
	if errGetSupervisor != nil {
		helpers.CheckErr("Failed get supervisor @SendEmailToSupervisor - logicMeal", errGetSupervisor)
		return errGetSupervisor
	}

	otherRequestorName, errGetOtherRequestor := DBUser.GetOtherRequestorNameByEmployeeNumbers(respMeal.OtherRequestorID)
	if errGetOtherRequestor != nil {
		helpers.CheckErr("Failed get other requestor name @SendEmailToSupervisor - logicMeal", errGetOtherRequestor)
		return errGetOtherRequestor
	}

	reqDateTime, errParseDate := time.Parse(constant.FormatDateTime, respMeal.RequestDate)
	if errParseDate != nil {
		helpers.CheckErr("Failed to parse request date @SendEmailToSupervisor - logicMeal", errParseDate)
		return errParseDate
	}

	token, errCreateToken := CreateToken(respMeal.ID)
	if errCreateToken != nil {
		helpers.CheckErr("Failed to create token @SendEmailToSupervisor - logicMeal", errCreateToken)
		return errCreateToken
	}
	baseURL := constant.GetClientURL() + "/approval?id=" + strconv.FormatInt(respMeal.ID, 10) + "&token=" + token
	approveLink := baseURL + "&result=true"
	rejectLink := baseURL + "&result=false"

	beego.Debug("COBAAIM => approveLink", approveLink)
	beego.Debug("COBAAIM => rejectLink", rejectLink)

	go helpers.GoMailSupervisorApproval(supervisor.Email, structLogic.InfoMailSupervisorApproval{
		SupervisorName: supervisor.Name,
		RequestorName:  requestor.Name,

		RequestDate:        reqDateTime.Format("2 January 2006"),
		RequestTime:        reqDateTime.Format("15:04"),
		OtherRequestorName: otherRequestorName,
		Amount:             fmt.Sprintf("IDR %s", humanize.Commaf(float64(respMeal.Amount))),
		BriefDescription:   respMeal.BriefDescription,
		Notes:              respMeal.Notes,

		ApproveLink: approveLink,
		RejectLink:  rejectLink,
	})

	return nil
}

// DownloadFormPDF ...
func DownloadFormPDF(ID int64, folderPath string, filenameFormat string) (filePath string, filename string, err error) {
	mealRequest, errGetMeal := DBMeal.GetMealRequestByID(ID)
	if errGetMeal != nil {
		helpers.CheckErr("Error get meal request @DownloadFormPDF - logicMeal", errGetMeal)
		err = errGetMeal
		return
	}

	if mealRequest.Status != constant.StatusApproved {
		err = errors.New("Overtime meals request for id : " + strconv.FormatInt(ID, 10) + " not approved yet")
		return
	}

	user, errGetUser := DBUser.GetUserByEmployeeNumber(mealRequest.RequestorID)
	if errGetUser != nil {
		helpers.CheckErr("Error get user @DownloadFormPDF - logicMeal", errGetUser)
		err = errGetUser
		return
	}

	supervisor, errGetSpv := DBUser.GetUserByEmployeeNumber(mealRequest.SupervisorID)
	if errGetSpv != nil {
		helpers.CheckErr("Error get supervisor @DownloadFormPDF - logicMeal", errGetSpv)
		err = errGetSpv
		return
	}

	name := strings.Replace(user.Name, " ", "", -1)
	date := mealRequest.RequestDate.Format("20060102")
	filename = strings.Replace(filenameFormat, "{NAME}", name, 1)
	filename = strings.Replace(filename, "{DATE}", date, 1)

	filePath = folderPath + filename

	err = writeFormPDF(filePath, mealRequest, user, supervisor)
	helpers.CheckErr("Error write PDF @DownloadFormPDF - logicMeal", err)

	return
}

// writeFormPDF ...
func writeFormPDF(
	filePath string,
	mealRequest structLogic.GetMealRequest,
	user structAPI.Employee,
	supervisor structAPI.Employee,
) error {
	otherRequestorName, errGet := DBUser.GetOtherRequestorNameByEmployeeNumbers(mealRequest.OtherRequestorID)
	helpers.CheckErr("Error get other requestor name @writeFormPDF - logicMeal", errGet)

	//init page
	//A4 measures 210 × 297 millimeters
	m := float64(20) //margin 20mm
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetMargins(m, m, m)

	//Header
	pdf.SetFont("Arial", "B", 20)
	pdf.WriteAligned(0, 10, "PT TNIS SERVICE INDONESIA", gofpdf.AlignCenter)
	pdf.Ln(16)
	pdf.WriteAligned(0, 10, "Overtime Meals Form", gofpdf.AlignCenter)
	pdf.Ln(10)
	pdf.Line(m, pdf.GetY(), m+170, pdf.GetY())

	//Body
	pdf.SetFont("Arial", "", 12)
	pdf.Write(10, "Request Date: "+mealRequest.RequestDate.Format("2 January 2006"))
	pdf.Ln(8)
	pdf.Write(10, "Request Time: "+mealRequest.RequestDate.Format("15:04"))
	pdf.Ln(8)
	pdf.Write(10, "Requestor Name: "+user.Name)
	pdf.Ln(8)
	pdf.Write(10, "Other Requestor (if any):")
	pdf.Ln(8)

	if len(otherRequestorName) > 0 {
		for _, otherRequestor := range otherRequestorName {
			pdf.Write(10, "                - "+otherRequestor.Name)
			pdf.Ln(8)
		}
	}

	pdf.Ln(4)
	pdf.Write(10, "Total Amount Requested: "+fmt.Sprintf("IDR %s", humanize.Commaf(float64(mealRequest.Amount))))
	pdf.Ln(4)
	pdf.SetFont("Arial", "", 11)
	pdf.Write(10, "( max. IDR 50,000/person )")

	pdf.Ln(12)
	pdf.SetFont("Arial", "", 12)
	pdf.Write(10, "Please give a brief description about your overtime work: ")
	pdf.Ln(8)
	pdf.Write(4, mealRequest.BriefDescription)

	pdf.Ln(8)
	pdf.Write(10, "Notes (optional): ")
	pdf.Ln(8)
	pdf.Write(4, mealRequest.Notes)

	//Signature
	pdf.Ln(12)
	pdf.Write(10, "Requestor:")
	pdf.Ln(4)
	pdf.Write(10, user.Name)
	pdf.Ln(12)
	pdf.Write(10, "Approved by:")
	pdf.Ln(4)
	pdf.Write(10, supervisor.Name)
	pdf.Ln(6)

	//New Page for Receipt
	pdf.AddPage()
	pdf.SetMargins(m, m, m)
	receiptFolderPath := constant.GOPATH + "/src/" + constant.GOAPP
	pdf.Image(receiptFolderPath+mealRequest.ReceiptUploadPath, 55, 10, 100, 0, true, "", 0, "")

	err := pdf.OutputFileAndClose(filePath)
	if err != nil {
		helpers.CheckErr("Error create pdf, with message = ", err)
		return err
	}

	return nil
}

// SaveReceiptFile ...
func SaveReceiptFile(
	employeeNumber int64,
	requestDate string,
	file multipart.File,
	fileHeader *multipart.FileHeader,
) (pathFile string, err error) {
	defer file.Close()

	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		return "", err
	}
	bytesToCheck := fileBytes[:512]
	//Validation only for Image type
	if !helpers.CheckValidImage(bytesToCheck) {
		return "", errors.New("Invalid File Extension, only .jpeg, .jpg, .png allowed")
	}

	user, errGet := DBUser.GetUserByEmployeeNumber(employeeNumber)
	if errGet != nil {
		return "", errGet
	}

	//Get current path
	dir, err := os.Getwd()
	if err != nil {
		return "", err
	}

	pathFile = "/storages/receipts/"
	reqDate, err := time.Parse(constant.FormatDateTime, requestDate)
	if err != nil {
		return "", err
	}
	filename := "OvertimeMealsReceipt_" + strings.Replace(user.Name, " ", "", -1) + "_" + reqDate.Format("20060102") + filepath.Ext(fileHeader.Filename)
	targetFile := pathFile + filename

	//Create New File
	f, err := os.OpenFile(dir+targetFile, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		return "", err
	}
	defer f.Close()

	//Overwrite with bytes from 'receipt' image
	f.Write(fileBytes)

	return targetFile, nil
}

// TokenValidation ...
func TokenValidation(mealRequestID int64, token string) (valid bool, err error) {
	tokenToCompare, errCreateToken := CreateToken(mealRequestID)
	if errCreateToken != nil {
		helpers.CheckErr("Failed to Create Token @TokenValidation - logicMeal", errCreateToken)
		return false, errCreateToken
	}

	if strings.Compare(token, tokenToCompare) != 0 {
		return false, nil
	}

	return true, nil
}

// CreateToken ...
// Format MD5Hash(MealRequestID_RequestorID_SupervisorID)
func CreateToken(mealRequestID int64) (token string, err error) {
	meal, errGetMeal := DBMeal.GetMealRequestByID(mealRequestID)
	if errGetMeal != nil {
		helpers.CheckErr("Failed Get Meal Request @CreateToken - logicMeal", errGetMeal)
		return "", errGetMeal
	}

	beforeHash := fmt.Sprintf("%d_%d_%d", meal.ID, meal.RequestorID, meal.SupervisorID)
	md5Hash := md5.Sum([]byte(beforeHash))
	afterHash := hex.EncodeToString(md5Hash[:])

	return afterHash, nil
}

// ApprovalMealRequest ...
func ApprovalMealRequest(req structAPI.PostApprovalMealRequest) error {
	meal, errGetMeal := DBMeal.GetMealRequestByID(req.ID)
	if errGetMeal != nil {
		helpers.CheckErr("Failed Get Meal Request @ApprovalMealRequest - logicMeal", errGetMeal)
		return errGetMeal
	}

	if meal.Status == constant.StatusApproved || meal.Status == constant.StatusRejected {
		return errors.New("Overtime meals request for ID : " + strconv.FormatInt(meal.ID, 10) + " already " + meal.Status)
	}

	errApproval := DBMeal.ApprovalMealRequest(req)
	if errApproval != nil {
		return errApproval
	}

	requestor, errGetEmployee := DBUser.GetUserByEmployeeNumber(meal.RequestorID)
	if errGetEmployee != nil {
		helpers.CheckErr("Failed get employee @ApprovalMealRequest - logicMeal", errGetEmployee)
		return errGetEmployee
	}

	supervisor, errGetSupervisor := DBUser.GetUserByEmployeeNumber(meal.SupervisorID)
	if errGetSupervisor != nil {
		helpers.CheckErr("Failed get supervisor @ApprovalMealRequest - logicMeal", errGetSupervisor)
		return errGetSupervisor
	}

	otherRequestorName, errGetOtherRequestor := DBUser.GetOtherRequestorNameByEmployeeNumbers(meal.OtherRequestorID)
	if errGetOtherRequestor != nil {
		helpers.CheckErr("Failed get other requestor name @ApprovalMealRequest - logicMeal", errGetOtherRequestor)
		return errGetOtherRequestor
	}

	meal, errGetMeal = DBMeal.GetMealRequestByID(req.ID)
	if errGetMeal != nil {
		helpers.CheckErr("Failed Get Meal Request 2nd @ApprovalMealRequest - logicMeal", errGetMeal)
		return errGetMeal
	}

	if req.Result {
		userHRD, errGetUserHRD := DBUser.GetUserHRD()
		if errGetUserHRD != nil {
			return errGetUserHRD
		}

		//Mail for Approved Request for HRD
		go helpers.GoMailApprovedRequest(userHRD.Email, structLogic.InfoMailApprovedRequest{
			HRName:         userHRD.Name,
			RequestorName:  requestor.Name,
			SupervisorName: supervisor.Name,

			RequestDate:        meal.RequestDate.Format("2 January 2006"),
			RequestTime:        meal.RequestDate.Format("15:04"),
			OtherRequestorName: otherRequestorName,
			Amount:             fmt.Sprintf("IDR %s", humanize.Commaf(float64(meal.Amount))),
			BriefDescription:   meal.BriefDescription,
			Notes:              meal.Notes,
		})
	}

	//Mail for Requestor that request has been submitted (Approved/Rejeceted)
	go helpers.GoMailSubmittedRequest(requestor.Email, structLogic.InfoMailSubmittedRequest{
		RequestorName:  requestor.Name,
		SupervisorName: supervisor.Name,

		RequestDate:        meal.RequestDate.Format("2 January 2006"),
		RequestTime:        meal.RequestDate.Format("15:04"),
		OtherRequestorName: otherRequestorName,
		Amount:             fmt.Sprintf("IDR %s", humanize.Commaf(float64(meal.Amount))),
		BriefDescription:   meal.BriefDescription,
		Notes:              meal.Notes,
		Reason:             meal.RejectReason,
		ActionDate:         meal.ActionDate.Format("2 January 2006 15:04"),
		Status:             strings.ToUpper(string(meal.Status[0])) + strings.ToLower(string(meal.Status[1:])),
	})

	return nil
}
