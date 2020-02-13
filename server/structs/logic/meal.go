package logic

import "time"

// GetMealRequest ...
type GetMealRequest struct {
	ID                int64     `json:"id" orm:"column(id);pk"`
	RequestorID       int64     `json:"requestor_id" orm:"column(requestor_id)"`
	OtherRequestorID  string    `json:"other_requestor_id" orm:"column(other_requestor_id)"`
	SupervisorID      int64     `json:"supervisor_id" orm:"column(supervisor_id)"`
	Amount            float64   `json:"amount" orm:"column(amount)"`
	BriefDescription  string    `json:"brief_description" orm:"column(brief_description)"`
	Notes             string    `json:"notes" orm:"column(notes)"`
	ReceiptUploadPath string    `json:"receipt_upload_path" orm:"column(receipt_upload_path)"`
	RequestDate       time.Time `json:"request_date" orm:"column(request_date)"`
	Status            string    `json:"status" orm:"column(status)"`
	RejectReason      string    `json:"reject_reason" orm:"column(reject_reason)"`
	ActionDate        time.Time `json:"action_date" orm:"column(action_date)"`
	CreatedAt         time.Time `json:"created_at" orm:"column(created_at)"`
	UpdatedAt         time.Time `json:"updated_at" orm:"column(updated_at)"`
}
