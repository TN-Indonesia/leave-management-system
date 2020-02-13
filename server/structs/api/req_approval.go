package api

// PostApprovalMealRequest ...
type PostApprovalMealRequest struct {
	ID           int64  `json:"id"`
	Result       bool   `json:"result"`
	RejectReason string `json:"reject_reason"`
	Token        string `json:"token"`
}
