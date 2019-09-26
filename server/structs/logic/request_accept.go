package logic

// RequestAccept ...
type RequestAccept struct {
	ID                 int64   `json:"id" orm:"column(id);pk"`
	EmployeeNumber     int64   `json:"employee_number" orm:"column(employee_number);pk"`
	Name               string  `json:"name" orm:"column(name)"`
	Gender             string  `json:"gender" orm:"column(gender)"`
	Position           string  `json:"position" orm:"column(position)"`
	StartWorkingDate   string  `json:"start_working_date" orm:"column(start_working_date)"`
	Email              string  `json:"email" orm:"column(email)"`
	Role               string  `json:"role" orm:"column(role)"`
	TypeName           string  `json:"type_name" orm:"column(type_name)"`
	LeaveRemaining     float64 `json:"leave_remaining" orm:"column(leave_remaining)"`
	BeforeLeaveBalance float64 `json:"before_leave_balance" orm:"column(before_leave_balance)"`
	AfterLeaveBalance  float64 `json:"after_leave_balance" orm:"column(after_leave_balance)"`
	Reason             string  `json:"reason" orm:"column(reason)"`
	DateFrom           string  `json:"date_from" orm:"column(date_from)"`
	DateTo             string  `json:"date_to" orm:"column(date_to)"`
	HalfDates          string  `json:"half_dates" orm:"column(half_dates)"`
	BackOn             string  `json:"back_on" orm:"column(back_on)"`
	Total              float64 `json:"total" orm:"column(total)"`
	ContactAddress     string  `json:"contact_address" orm:"column(contact_address)"`
	ContactNumber      string  `json:"contact_number" orm:"column(contact_number)"`
	Status             string  `json:"status" orm:"column(status)"`
	ActionBy           string  `json:"action_by" orm:"column(action_by)"`
	Notes              string  `json:"notes" orm:"column(notes)"`
}

// ReportLeaveRequest ...
type ReportLeaveRequest struct {
	ID               int64   `json:"Request ID" orm:"column(id);pk"`
	EmployeeNumber   int64   `json:"Employee Number" orm:"column(employee_number);pk"`
	Name             string  `json:"Name" orm:"column(name)"`
	Gender           string  `json:"Gender" orm:"column(gender)"`
	Position         string  `json:"Position" orm:"column(position)"`
	StartWorkingDate string  `json:"Start Working Date" orm:"column(start_working_date)"`
	Email            string  `json:"Email" orm:"column(email)"`
	TypeName         string  `json:"Type of Leave" orm:"column(type_name)"`
	Reason           string  `json:"Reason" orm:"column(reason)"`
	DateFrom         string  `json:"From" orm:"column(date_from)"`
	DateTo           string  `json:"To" orm:"column(date_to)"`
	HalfDates        string  `json:"Half Day" orm:"column(half_dates)"`
	BackOn           string  `json:"Back To Work" orm:"column(back_on)"`
	Total            float64 `json:"Total Leave" orm:"column(total)"`
	LeaveRemaining   float64 `json:"Leave Balance" orm:"column(leave_remaining)"`
	ContactAddress   string  `json:"Contact Address" orm:"column(contact_address)"`
	ContactNumber    string  `json:"Contact Number" orm:"column(contact_number)"`
	Notes            string  `json:"notes" orm:"column(notes)"`
}
