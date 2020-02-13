package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"time"

	"server/helpers"
	"server/helpers/constant"
	logicMeal "server/models/logic/meal"
	structAPI "server/structs/api"

	"github.com/astaxie/beego"
)

//MealController ...
type MealController struct {
	beego.Controller
}

// GetMealRequestByID ...
func (c *MealController) GetMealRequestByID() {
	var resp structAPI.RespData

	idStr := c.Ctx.Input.Param(":id")
	ID, errCon := strconv.ParseInt(idStr, 0, 64)
	if errCon != nil {
		helpers.CheckErr("Convert id failed @GetMealRequestByID - controller", errCon)
		resp.Error = errors.New("Convert id failed").Error()
		return
	}

	resGet, errGetReject := logicMeal.GetMealRequestByID(ID)
	if errGetReject != nil {
		resp.Error = errGetReject.Error()
		c.Ctx.Output.SetStatus(400)
	} else {
		respMeal := structAPI.RespMeal{
			RequestorID:       resGet.RequestorID,
			OtherRequestorID:  resGet.OtherRequestorID,
			SupervisorID:      resGet.SupervisorID,
			Amount:            resGet.Amount,
			BriefDescription:  resGet.BriefDescription,
			Notes:             resGet.Notes,
			ReceiptUploadPath: resGet.ReceiptUploadPath,
			RequestDate:       resGet.RequestDate.Format(constant.FormatDateTime),
			Status:            resGet.Status,
			CreatedAt:         resGet.CreatedAt.Format(constant.FormatDateTime),
			UpdatedAt:         resGet.UpdatedAt.Format(constant.FormatDateTime),
		}
		resp.Body = respMeal
	}

	err := c.Ctx.Output.JSON(resp, false, false)
	if err != nil {
		helpers.CheckErr("Failed giving output @GetMealRequestByID - controller", err)
	}
}

// PostMealRequest ...
func (c *MealController) PostMealRequest() {
	var (
		req  structAPI.ReqMeal
		resp structAPI.RespData
	)
	idStr := c.Ctx.Input.Param(":id")
	ID, errCon := strconv.ParseInt(idStr, 0, 64)
	if errCon != nil {
		helpers.CheckErr("Convert id failed @PostMealRequest - controller", errCon)
		resp.Error = errors.New("Convert id failed").Error()
		c.Ctx.Output.SetStatus(500)
		c.Ctx.Output.JSON(resp, false, false)
		return
	}
	//Parse Form data
	var errFloat error
	req.RequestDate = c.GetString("request_date")
	req.OtherRequestorID = c.GetString("other_requestor_id")
	req.Amount, errFloat = c.GetFloat("amount")
	if errFloat != nil {
		helpers.CheckErr("Failed to convert float @PostMealRequest - controller", errFloat)
		resp.Error = errors.New("Failed to convert float").Error()
		c.Ctx.Output.SetStatus(500)
		c.Ctx.Output.JSON(resp, false, false)
		return
	}
	req.BriefDescription = c.GetString("brief_description")
	req.Notes = c.GetString("notes")
	if req.Notes == "null" {
		req.Notes = ""
	}

	beego.Debug("Post Meal Request : ", req)

	// Validation
	beego.Debug("Validation amount : ", req.Amount)
	if req.Amount < 0 {
		beego.Debug("Masuk validation amount")
		resp.Error = errors.New("Amount cannot be less than 0 (zero)").Error()
		c.Ctx.Output.SetStatus(400)
		c.Ctx.Output.JSON(resp, false, false)
		return
	}

	today, errParseDate := time.Parse(constant.FormatDateTime, time.Now().Format(constant.FormatDateTime))
	if errParseDate != nil {
		helpers.CheckErr("Failed to convert date @PostMealRequest - controller", errParseDate)
		resp.Error = errors.New("Failed to convert date").Error()
		c.Ctx.Output.SetStatus(500)
		c.Ctx.Output.JSON(resp, false, false)
		return
	}

	reqDate, errParseDate := time.Parse(constant.FormatDateTime, req.RequestDate)
	if errParseDate != nil {
		helpers.CheckErr("Failed to convert date @PostMealRequest - controller", errParseDate)
		resp.Error = errors.New("Failed to convert date").Error()
		c.Ctx.Output.SetStatus(500)
		c.Ctx.Output.JSON(resp, false, false)
		return
	}

	beego.Debug("Validation request date : ", reqDate, today, reqDate.After(today))
	if reqDate.After(today) {
		beego.Debug("Masuk validation date")
		resp.Error = errors.New("Request Date cannot be greater than Today's Date").Error()
		c.Ctx.Output.SetStatus(400)
		c.Ctx.Output.JSON(resp, false, false)
		return
	}

	//=======START, handle upload image 'receipt'
	file, fileHeader, errGetFile := c.GetFile("receipt")
	if errGetFile != nil {
		helpers.CheckErr("Failed to get file @PostMealRequest - controller", errGetFile)
		resp.Error = errors.New("Failed to get file").Error()
		c.Ctx.Output.SetStatus(500)
		c.Ctx.Output.JSON(resp, false, false)
		return
	}

	var errUpload error
	req.ReceiptUploadPath, errUpload = logicMeal.SaveReceiptFile(ID, req.RequestDate, file, fileHeader)
	if errUpload != nil {
		helpers.CheckErr("Failed to save file @PostMealRequest - controller", errUpload)
		resp.Error = errors.New("Failed to save file").Error()
		c.Ctx.Output.SetStatus(500)
		c.Ctx.Output.JSON(resp, false, false)
		return
	}
	//=======END, handle upload image 'receipt'

	//Create Meal Request to DB
	respMeal, errCreate := logicMeal.PostMealRequest(ID, &req)
	if errCreate != nil {
		helpers.CheckErr("Failed to create meal request @PostMealRequest - controller", errCreate)
		resp.Error = errors.New("Failed to create meal request").Error()
		c.Ctx.Output.SetStatus(500)
		c.Ctx.Output.JSON(resp, false, false)
		return
	}

	// Send Email Approval to Supervisor
	err := logicMeal.SendEmailToSupervisor(respMeal)
	if err != nil {
		helpers.CheckErr("Failed send email to supervisor @PostMealRequest - Controller", err)
	}

	resp.Body = respMeal

	errOutput := c.Ctx.Output.JSON(resp, false, false)
	if errOutput != nil {
		helpers.CheckErr("Failed giving output @PostMealRequest - controller", errOutput)
	}
}

// GetDownloadFormPDF ...
func (c *MealController) GetDownloadFormPDF() {
	var resp structAPI.RespData

	idStr := c.Ctx.Input.Param(":id")
	ID, errCon := strconv.ParseInt(idStr, 0, 64)
	if errCon != nil {
		helpers.CheckErr("Convert id failed @GetDownloadFormPDF - controller", errCon)
		resp.Error = errors.New("Convert id failed").Error()
		c.Ctx.Output.SetStatus(500)
		c.Ctx.Output.JSON(resp, false, false)
		return
	}

	filenameFormat := "OvertimeMealsForm_{NAME}_{DATE}.pdf"
	folderPath := constant.GOPATH + "/src/" + constant.GOAPP + "/storages/forms/"

	filePath, filename, errDownload := logicMeal.DownloadFormPDF(ID, folderPath, filenameFormat)
	if errDownload != nil {
		helpers.CheckErr("Error download pdf @GetDownloadFormPDF", errDownload)
		resp.Error = errDownload.Error()
		c.Ctx.Output.SetStatus(400)
		c.Ctx.Output.JSON(resp, false, false)
		return
	}

	c.Ctx.Output.Download(filePath, filename)
}

// PostApprovalMealRequest ...
func (c *MealController) PostApprovalMealRequest() {
	var (
		resp        structAPI.RespData
		reqApproval structAPI.PostApprovalMealRequest
	)

	body := c.Ctx.Input.RequestBody
	fmt.Println("CREATE-LEAVE-REQUEST=======>", string(body))

	errMarshal := json.Unmarshal(body, &reqApproval)
	if errMarshal != nil {
		helpers.CheckErr("Failed unmarshall req body @PostApprovalMealRequest - controller", errMarshal)
		resp.Error = errors.New("Type request malform").Error()
		c.Ctx.Output.SetStatus(400)
		c.Ctx.Output.JSON(resp, false, false)
		return
	}

	// Token Validation
	validToken, errTokenValidation := logicMeal.TokenValidation(reqApproval.ID, reqApproval.Token)
	if errTokenValidation != nil {
		helpers.CheckErr("Failed to validate token @PostApprovalMealRequest - controller", errTokenValidation)
		resp.Error = errors.New("Bad Token").Error()
		c.Ctx.Output.SetStatus(400)
		c.Ctx.Output.JSON(resp, false, false)
		return
	}

	if !validToken {
		resp.Error = errors.New("Invalid Token").Error()
		c.Ctx.Output.SetStatus(400)
		c.Ctx.Output.JSON(resp, false, false)
		return
	}

	//Approval for Meal Request
	errApproval := logicMeal.ApprovalMealRequest(reqApproval)
	if errApproval != nil {
		helpers.CheckErr("Approval failed @PostApprovalMealRequest - controller", errApproval)
		resp.Error = errApproval.Error()
		c.Ctx.Output.SetStatus(400)
		c.Ctx.Output.JSON(resp, false, false)
		return
	}

	resp.Body = nil

	errOutput := c.Ctx.Output.JSON(resp, false, false)
	if errOutput != nil {
		helpers.CheckErr("Failed giving output @PostApprovalMealRequest - controller", errOutput)
	}
}
