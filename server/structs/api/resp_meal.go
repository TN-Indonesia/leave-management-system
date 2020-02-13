package api

type (
	// RespMeal ...
	RespMeal struct {
		ID                int64   `json:"id"`
		RequestorID       int64   `json:"requestor_id"`
		OtherRequestorID  string  `json:"other_requestor_id"`
		SupervisorID      int64   `json:"supervisor_id"`
		Amount            float64 `json:"amount"`
		BriefDescription  string  `json:"brief_description"`
		Notes             string  `json:"notes"`
		ReceiptUploadPath string  `json:"receipt_upload_path"`
		RequestDate       string  `json:"request_date"`
		Status            string  `json:"status"`
		CreatedAt         string  `json:"created_at"`
		UpdatedAt         string  `json:"updated_at"`
	}
)
