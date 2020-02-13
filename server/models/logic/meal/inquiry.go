package meal

import (
	"server/helpers"
	"server/helpers/constant"
	structAPI "server/structs/api"
	structDB "server/structs/db"
	"strconv"
)

// GetMealRequestForEmployeeInquiry ...
func GetMealRequestForEmployeeInquiry(ID int64, status string) (respInquiry []structAPI.RespInquiry, err error) {
	resDB, err := DBMeal.GetMealRequestForEmployeeInquiry(ID, status)
	if err != nil {
		helpers.CheckErr("Error get overtime meal request @GetMealRequestForEmployeeInquiry - logicMeal", err)
		return
	}

	respInquiry, err = ParseStructMealRequestToRespInquiry(resDB)
	if err != nil {
		helpers.CheckErr("Error get overtime meal request @GetMealRequestForEmployeeInquiry - logicMeal", err)
		return
	}

	return
}

// GetMealRequestForSupervisorInquiry ...
func GetMealRequestForSupervisorInquiry(ID int64, status string) (respInquiry []structAPI.RespInquiry, err error) {
	resDB, err := DBMeal.GetMealRequestForSupervisorInquiry(ID, status)
	if err != nil {
		helpers.CheckErr("Error get overtime meal request @GetMealRequestForSupervisorInquiry - logicMeal", err)
		return
	}

	respInquiry, err = ParseStructMealRequestToRespInquiry(resDB)
	if err != nil {
		helpers.CheckErr("Error get overtime meal request @GetMealRequestForSupervisorInquiry - logicMeal", err)
		return
	}

	return
}

// GetMealRequestForAdminInquiry ...
func GetMealRequestForAdminInquiry() (respInquiry []structAPI.RespInquiry, err error) {
	resDB, err := DBMeal.GetMealRequestForAdminInquiry()
	if err != nil {
		helpers.CheckErr("Error get overtime meal request @GetMealRequestForAdminInquiry - logicMeal", err)
		return
	}

	respInquiry, err = ParseStructMealRequestToRespInquiry(resDB)
	if err != nil {
		helpers.CheckErr("Error get overtime meal request @GetMealRequestForAdminInquiry - logicMeal", err)
		return
	}

	return
}

// ParseStructMealRequestToRespInquiry ...
func ParseStructMealRequestToRespInquiry(resDB []structDB.MealRequest) (respInquiry []structAPI.RespInquiry, err error) {
	for _, meal := range resDB {
		otherRequestor, errGet := DBUser.GetOtherRequestorNameByEmployeeNumbers(meal.OtherRequestorID)
		if errGet != nil {
			helpers.CheckErr("Failed to get toher requestor name @ParseStructMealRequestToRespInquiry - logicMeal ", errGet)
			err = errGet
			return
		}

		requestor, errGet := DBUser.GetUserByEmployeeNumber(meal.RequestorID)
		if errGet != nil {
			helpers.CheckErr("Failed to get user with employee number : "+strconv.FormatInt(meal.ID, 10)+" @ParseStructMealRequestToRespInquiry - logicMeal", errGet)
			err = errGet
			return
		}

		supervisor, errGet := DBUser.GetUserByEmployeeNumber(meal.SupervisorID)
		if errGet != nil {
			helpers.CheckErr("Failed to get user with employee number : "+strconv.FormatInt(meal.SupervisorID, 10)+" @ParseStructMealRequestToRespInquiry - logicMeal", errGet)
			err = errGet
			return
		}

		respInquiry = append(respInquiry, structAPI.RespInquiry{
			ID:                meal.ID,
			RequestorID:       requestor.EmployeeNumber,
			RequestorName:     requestor.Name,
			OtherRequestor:    otherRequestor,
			SupervisorID:      supervisor.EmployeeNumber,
			SupervisorName:    supervisor.Name,
			Amount:            meal.Amount,
			BriefDescription:  meal.BriefDescription,
			Notes:             meal.Notes,
			ReceiptUploadPath: meal.ReceiptUploadPath,
			RequestDate:       meal.RequestDate.Format(constant.FormatDateTime),
			Status:            meal.Status,
			RejectReason:      meal.RejectReason,
			ActionDate:        meal.ActionDate.Format(constant.FormatDateTime),
			CreatedAt:         meal.CreatedAt.Format(constant.FormatDateTime),
			UpdatedAt:         meal.UpdatedAt.Format(constant.FormatDateTime),
		})
	}

	return
}
