package api

type (
	// ReqMeal ...
	ReqMeal struct {
		RequestDate       string  `json:"request_date"`
		OtherRequestorID  string  `json:"other_requestor_id"`
		Amount            float64 `json:"amount"`
		BriefDescription  string  `json:"brief_description"`
		Notes             string  `json:"notes"`
		ReceiptUploadPath string  `json:"receipt_upload_path"`
	}
)
