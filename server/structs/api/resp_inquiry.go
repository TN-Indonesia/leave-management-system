package api

import (
	structLogic "server/structs/logic"
)

// RespInquiry ...
type RespInquiry struct {
	ID                int64                                                `json:"id"`
	RequestorID       int64                                                `json:"requestor_id"`
	RequestorName     string                                               `json:"requestor_name"`
	OtherRequestor    []structLogic.GetOtherRequestorNameByEmployeeNumbers `json:"other_requestor_id"`
	SupervisorID      int64                                                `json:"supervisor_id"`
	SupervisorName    string                                               `json:"supervisor_name"`
	Amount            float64                                              `json:"amount"`
	BriefDescription  string                                               `json:"brief_description"`
	Notes             string                                               `json:"notes"`
	ReceiptUploadPath string                                               `json:"receipt_upload_path"`
	RequestDate       string                                               `json:"request_date"`
	Status            string                                               `json:"status"`
	RejectReason      string                                               `json:"reject_reason"`
	ActionDate        string                                               `json:"action_date"`
	CreatedAt         string                                               `json:"created_at"`
	UpdatedAt         string                                               `json:"updated_at"`
}
