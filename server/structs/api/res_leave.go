package api

// CreateLeaveRequest ...
type CreateLeaveRequest struct {
	EmployeeNumber int64    `json:"employee_number"`
	TypeLeaveID    int64    `json:"type_leave_id"`
	Reason         string   `json:"reason"`
	DateFrom       string   `json:"date_from"`
	DateTo         string   `json:"date_to"`
	HalfDates      []string `json:"half_dates"`
	Total          float64  `json:"total"`
	BackOn         string   `json:"back_on"`
	ContactAddress string   `json:"contact_address"`
	ContactNumber  string   `json:"contact_number"`
	Status         string   `json:"status"`
	Notes          string   `json:"notes"`
}

// TableName ...
func (u *CreateLeaveRequest) TableName() string {
	return "leave_request"
}
