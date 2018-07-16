package api

// CreateLeaveRequest ...
type CreateLeaveRequest struct {
	EmployeeNumber int64  `json:"employee_number" orm:"column(employee_number)"`
	TypeLeaveID    int64  `json:"type_leave_id" orm:"column(type_leave_id)"`
	Reason         string `json:"reason" orm:"column(reason)"`
	DateFrom       string `json:"date_from" orm:"column(date_from)"`
	DateTo         string `json:"date_to" orm:"column(date_to)"`
	// DateRanges     []DateRange `json:"date_ranges"`
	Total          int    `json:"total" orm:"column(total)"`
	BackOn         string `json:"back_on" orm:"column(back_on)"`
	ContactAddress string `json:"contact_address" orm:"column(contact_address)"`
	ContactNumber  string `json:"contact_number" orm:"column(contact_number)"`
	Status         string `json:"status" orm:"column(status)"`
}

// TableName ...
func (u *CreateLeaveRequest) TableName() string {
	return "leave_request"
}
