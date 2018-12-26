package logic

// GetEmployee ...
type GetEmployee struct {
	Name             string `json:"name" orm:"column(name)"`
	Email            string `json:"email" orm:"column(email)"`
	StartWorkingDate string `json:"start_working_date" orm:"column(start_working_date)"`
}

// GetEmployeeByNumber ...
type GetEmployeeByNumber struct {
	ID               int64  `json:"id" orm:"column(id)"`
	EmployeeNumber   int64  `json:"employee_number" orm:"column(employee_number)"`
	Name             string `json:"name" orm:"column(name)"`
	Email            string `json:"email" orm:"column(email)"`
	StartWorkingDate string `json:"start_working_date" orm:"column(start_working_date)"`
}

// NewPassword ...
type NewPassword struct {
	OldPassword     string `json:"old_password" orm:"column(old_password)"`
	NewPassword     string `json:"new_password" orm:"column(new_password)"`
	ConfirmPassword string `json:"confirm_password" orm:"column(confirm_password)"`
}

// PasswordReset ...
type PasswordReset struct {
	Email string `json:"email" orm:"column(email)"`
}
