package admin

import (
	structDB "server/structs/db"
	structLogic "server/structs/logic"
)

// IBaseAdmin ...
type IBaseAdmin interface {
	// AddUser
	AddUser(user structDB.User) error
	// DeleteUser
	DeleteUser(employeeNumber int64) error
	// GetUsers
	GetUsers() (
		[]structDB.User,
		error,
	)
	// GetUser
	GetUser(employeeNumber int64) (
		result structDB.User,
		err error,
	)
	// UpdateUser
	UpdateUser(
		e *structDB.User,
		employeeNumber int64,
	) (err error)
	// GetLeaveRequest
	GetLeaveRequest() ([]structLogic.RequestAccept, error)
	// GetLeaveRequest
	GetLeaveRequestPending() ([]structLogic.RequestPending, error)
	// GetLeaveRequest
	GetLeaveRequestReject() ([]structLogic.RequestReject, error)
	// CreateUserTypeLeave
	CreateUserTypeLeave(
		employeeNumber int64,
		typeLeaveID int64,
		total float64,
	) error
	// UpdateLeaveRemaning
	UpdateLeaveRemaning(total float64, employeeNumber int64, typeID int64) (err error)
}
