package api

// ReqRegister ...
type ReqRegister struct {
	EmployeeNumber   int64  `json:"employee_number" `
	Name             string `json:"name"`
	Gender           string `json:"gender"`
	Position         string `json:"position"`
	StartWorkingDate string `json:"start_working_date"`
	MobilePhone      string `json:"mobile_phone"`
	Email            string `json:"email"`
	Password         string `json:"password"`
	Role             string `json:"role"`
	SupervisorID     int64  `json:"supervisor_id"`
}
