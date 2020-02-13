package meal

import (
	structAPI "server/structs/api"
	structDB "server/structs/db"
	structLogic "server/structs/logic"
	"time"
)

// IBaseMeal ...
type IBaseMeal interface {
	// GetMealRequestByID
	GetMealRequestByID(ID int64) (
		mealRequest structLogic.GetMealRequest,
		err error,
	)
	// PostMealRequest ...
	PostMealRequest(
		RequestorID int64,
		OtherRequestorID string,
		SupervisorID int64,
		Amount float64,
		BriefDescription string,
		Notes string,
		ReceiptUploadPath string,
		RequestDate time.Time,
		Status string,
		CreatedAt time.Time,
		UpdatedAt time.Time,
	) (ID int64, err error)

	// ApprovalMealRequest ...
	ApprovalMealRequest(req structAPI.PostApprovalMealRequest) error

	// GetMealRequestForEmployeeInquiry ...
	GetMealRequestForEmployeeInquiry(ID int64, status string) (resDB []structDB.MealRequest, err error)
	// GetMealRequestForSupervisorInquiry ...
	GetMealRequestForSupervisorInquiry(ID int64, status string) (resDB []structDB.MealRequest, err error)
	// GetMealRequestForAdminInquiry ...
	GetMealRequestForAdminInquiry() (resDB []structDB.MealRequest, err error)
}
