package user

import (
	structAPI "server/structs/api"
	structDB "server/structs/db"
	structLogic "server/structs/logic"
)

// IBaseUser ...
type IBaseUser interface {
	// UserLogin
	UserLogin(loginData *structAPI.ReqLogin) (
		result structDB.User,
		err error,
	)
	// ForgotPassword
	ForgotPassword(e *structLogic.PasswordReset) error
	// GetUser
	GetUser(email string) (employee structLogic.GetEmployee, err error)
	// CountUserEmail
	CountUserEmail(email string) (int, error)
	// CountUserEmployeeNumber
	CountUserEmployeeNumber(employeeNumber int64) (int, error)

	// UpdatePassword
	UpdatePassword(
		p *structLogic.NewPassword,
		employeeNumber int64,
	) (err error)

	// GetDirector
	GetDirector() (
		director structLogic.GetDirector,
		err error,
	)
	// GetSupervisors
	GetSupervisors() (
		supervisor []structLogic.GetSupervisors,
		err error,
	)
	// GetSupervisor
	GetSupervisor(employeeNumber int64) (
		supervisor structLogic.GetSupervisor,
		err error,
	)
	// GetEmployee
	GetEmployee(employeeNumber int64) (
		employee structLogic.GetEmployee,
		err error,
	)
	// GetEmployee
	GetEmployeeByEmployeeNumber(employeeID int64) (
		employee structLogic.GetEmployeeByNumber,
		err error,
	)

	// GetTypeLeave
	GetTypeLeave() (
		typeLeave []structDB.TypeLeave,
		err error,
	)
	// CreateUserTypeLeave
	CreateUserTypeLeave(
		employeeNumber int64,
		typeLeaveID int64,
		leaveRemaining float64,
	) error
	// GetUserTypeLeave
	GetUserTypeLeave(employeeNumber int64) (
		userTypeLeave []structLogic.UserTypeLeave,
		err error,
	)
	// GetSumarry
	GetSumarry(employeeNumber int64) (
		sumarry []structLogic.UserSumarry,
		err error,
	)
	// GetUserLeaveRemaining
	GetUserLeaveRemaining(
		typeID int64,
		employeeNumber int64,
	) (
		userTypeLeave structLogic.UserTypeLeave,
		err error,
	)

	// =================== Overtime Meals ================
	GetUserByEmployeeNumber(employeeNumber int64) (
		user structAPI.Employee,
		err error,
	)

	GetListEmployee() (
		listEmployee []structAPI.ListEmployee,
		err error,
	)

	GetOtherRequestorNameByEmployeeNumbers(employeeNumbers string) (
		OtherRequestorName []structLogic.GetOtherRequestorNameByEmployeeNumbers,
		err error,
	)

	GetSupervisorEmailByMealRequestID(ID int64) (email string, err error)

	GetUserHRD() (user structDB.User, err error)
}
