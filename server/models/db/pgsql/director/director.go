package director

import (
	"errors"
	"server/helpers"
	"server/helpers/constant"
	logicLeave "server/models/db/pgsql/leave_request"
	logicUser "server/models/db/pgsql/user"
	structDB "server/structs/db"
	structLogic "server/structs/logic"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

// Director ...
type Director struct{}

// AcceptByDirector ...
func (u *Director) AcceptByDirector(id int64, employeeNumber int64) error {
	var (
		dbLeave structDB.LeaveRequest
		user    logicUser.User
		leave   logicLeave.LeaveRequest
	)

	o := orm.NewOrm()

	getDirector, _ := user.GetDirector()
	getEmployee, _ := user.GetEmployee(employeeNumber)
	getLeave, _ := leave.GetLeave(id)

	statAcceptDirector := constant.StatusSuccessInDirector
	actionBy := getDirector.Name

	_, errRAW := o.Raw(`UPDATE `+dbLeave.TableName()+` SET status = ?, action_by = ? WHERE id = ? AND employee_number = ?`, statAcceptDirector, actionBy, id, employeeNumber).Exec()
	if errRAW != nil {
		helpers.CheckErr("error update status @AcceptByDirector", errRAW)
	}
	helpers.GoMailDirectorAccept(getEmployee.Email, getLeave.ID, getEmployee.Name, getDirector.Name)

	return errRAW
}

// RejectByDirector ...
func (u *Director) RejectByDirector(l *structDB.LeaveRequest, id int64, employeeNumber int64) error {

	var (
		dbLeave structDB.LeaveRequest
		user    logicUser.User
		leave   logicLeave.LeaveRequest
	)

	o := orm.NewOrm()

	getDirector, _ := user.GetDirector()
	getEmployee, _ := user.GetEmployee(employeeNumber)
	getLeave, _ := leave.GetLeave(id)

	statRejectDirector := constant.StatusRejectInDirector
	actionBy := getDirector.Name
	rejectReason := l.RejectReason

	_, errRAW := o.Raw(`UPDATE `+dbLeave.TableName()+` SET status = ?, reject_reason = ?, action_by = ? WHERE id = ? AND employee_number = ?`, statRejectDirector, rejectReason, actionBy, id, employeeNumber).Exec()
	if errRAW != nil {
		helpers.CheckErr("error update status @RejectByDirector", errRAW)
	}
	helpers.GoMailDirectorReject(getEmployee.Email, getLeave.ID, getEmployee.Name, getDirector.Name, rejectReason)

	return errRAW
}

// GetDirectorPendingRequest ...
func (u *Director) GetDirectorPendingRequest() ([]structLogic.RequestPending, error) {
	var (
		reqPending    []structLogic.RequestPending
		user          structDB.User
		leave         structDB.LeaveRequest
		typeLeave     structDB.TypeLeave
		userTypeLeave structDB.UserTypeLeave
	)

	statPendingDirector := constant.StatusPendingInDirector
	o := orm.NewOrm()
	qb, errQB := orm.NewQueryBuilder("mysql")
	if errQB != nil {
		helpers.CheckErr("Query builder failed @GetDirectorPendingRequest", errQB)
		return reqPending, errQB
	}

	qb.Select(
		leave.TableName()+".id",
		user.TableName()+".employee_number",
		user.TableName()+".name",
		user.TableName()+".gender",
		user.TableName()+".position",
		user.TableName()+".start_working_date",
		user.TableName()+".mobile_phone",
		user.TableName()+".email",
		user.TableName()+".role",
		typeLeave.TableName()+".type_name",
		userTypeLeave.TableName()+".leave_remaining",
		leave.TableName()+".reason",
		leave.TableName()+".date_from",
		leave.TableName()+".date_to",
		leave.TableName()+".total",
		leave.TableName()+".back_on",
		leave.TableName()+".contact_address",
		leave.TableName()+".contact_number",
		leave.TableName()+".status",
		leave.TableName()+".action_by").
		From(user.TableName()).
		InnerJoin(leave.TableName()).
		On(user.TableName() + ".employee_number" + "=" + leave.TableName() + ".employee_number").
		InnerJoin(typeLeave.TableName()).
		On(typeLeave.TableName() + ".id" + "=" + leave.TableName() + ".type_leave_id").
		InnerJoin(userTypeLeave.TableName()).
		On(userTypeLeave.TableName() + ".type_leave_id" + "=" + leave.TableName() + ".type_leave_id").
		And(userTypeLeave.TableName() + ".employee_number" + "=" + leave.TableName() + ".employee_number").
		Where(`status = ? `)
	sql := qb.String()

	count, errRaw := o.Raw(sql, statPendingDirector).QueryRows(&reqPending)
	if errRaw != nil {
		helpers.CheckErr("Failed Query Select @GetDirectorPendingRequest", errRaw)
		return reqPending, errors.New("error get leave request pending")
	}
	beego.Debug("Total pending request =", count)

	return reqPending, errRaw
}

// GetDirectorAcceptRequest ...
func (u *Director) GetDirectorAcceptRequest() ([]structLogic.RequestAccept, error) {
	var (
		reqAccept     []structLogic.RequestAccept
		user          structDB.User
		leave         structDB.LeaveRequest
		typeLeave     structDB.TypeLeave
		userTypeLeave structDB.UserTypeLeave
	)
	statAcceptDirector := constant.StatusSuccessInDirector

	o := orm.NewOrm()
	qb, errQB := orm.NewQueryBuilder("mysql")
	if errQB != nil {
		helpers.CheckErr("Query builder failed @GetDirectorAcceptRequest", errQB)
		return reqAccept, errQB
	}

	qb.Select(
		leave.TableName()+".id",
		user.TableName()+".employee_number",
		user.TableName()+".name",
		user.TableName()+".gender",
		user.TableName()+".position",
		user.TableName()+".start_working_date",
		user.TableName()+".mobile_phone",
		user.TableName()+".email",
		user.TableName()+".role",
		typeLeave.TableName()+".type_name",
		userTypeLeave.TableName()+".leave_remaining",
		leave.TableName()+".reason",
		leave.TableName()+".date_from",
		leave.TableName()+".date_to",
		leave.TableName()+".total",
		leave.TableName()+".back_on",
		leave.TableName()+".contact_address",
		leave.TableName()+".contact_number",
		leave.TableName()+".status",
		leave.TableName()+".action_by").
		From(user.TableName()).
		InnerJoin(leave.TableName()).
		On(user.TableName() + ".employee_number" + "=" + leave.TableName() + ".employee_number").
		InnerJoin(typeLeave.TableName()).
		On(typeLeave.TableName() + ".id" + "=" + leave.TableName() + ".type_leave_id").
		InnerJoin(userTypeLeave.TableName()).
		On(userTypeLeave.TableName() + ".type_leave_id" + "=" + leave.TableName() + ".type_leave_id").
		And(userTypeLeave.TableName() + ".employee_number" + "=" + leave.TableName() + ".employee_number").
		Where(`status = ? `)
	sql := qb.String()

	count, errRaw := o.Raw(sql, statAcceptDirector).QueryRows(&reqAccept)
	if errRaw != nil {
		helpers.CheckErr("Failed Query Select @GetDirectorAcceptRequest", errRaw)
		return reqAccept, errors.New("error get leave")
	}
	beego.Debug("Total accept request =", count)

	return reqAccept, errRaw
}

// GetDirectorRejectRequest ...
func (u *Director) GetDirectorRejectRequest() ([]structLogic.RequestReject, error) {
	var (
		reqReject     []structLogic.RequestReject
		user          structDB.User
		leave         structDB.LeaveRequest
		typeLeave     structDB.TypeLeave
		userTypeLeave structDB.UserTypeLeave
	)
	StatRejectInDirector := constant.StatusRejectInDirector

	o := orm.NewOrm()
	qb, errQB := orm.NewQueryBuilder("mysql")
	if errQB != nil {
		helpers.CheckErr("Query builder failed @GetDirectorRejectRequest", errQB)
		return reqReject, errQB
	}

	qb.Select(
		leave.TableName()+".id",
		user.TableName()+".employee_number",
		user.TableName()+".name",
		user.TableName()+".gender",
		user.TableName()+".position",
		user.TableName()+".start_working_date",
		user.TableName()+".mobile_phone",
		user.TableName()+".email",
		user.TableName()+".role",
		typeLeave.TableName()+".type_name",
		userTypeLeave.TableName()+".leave_remaining",
		leave.TableName()+".reason",
		leave.TableName()+".date_from",
		leave.TableName()+".date_to",
		leave.TableName()+".total",
		leave.TableName()+".back_on",
		leave.TableName()+".contact_address",
		leave.TableName()+".contact_number",
		leave.TableName()+".status",
		leave.TableName()+".action_by").
		From(user.TableName()).
		InnerJoin(leave.TableName()).
		On(user.TableName() + ".employee_number" + "=" + leave.TableName() + ".employee_number").
		InnerJoin(typeLeave.TableName()).
		On(typeLeave.TableName() + ".id" + "=" + leave.TableName() + ".type_leave_id").
		InnerJoin(userTypeLeave.TableName()).
		On(userTypeLeave.TableName() + ".type_leave_id" + "=" + leave.TableName() + ".type_leave_id").
		And(userTypeLeave.TableName() + ".employee_number" + "=" + leave.TableName() + ".employee_number").
		Where(`status = ? `)
	sql := qb.String()

	count, errRaw := o.Raw(sql, StatRejectInDirector).QueryRows(&reqReject)
	if errRaw != nil {
		helpers.CheckErr("Failed Query Select @GetDirectorRejectRequest", errRaw)
		return reqReject, errors.New("error get leave")
	}
	beego.Debug("Total reject request =", count)

	return reqReject, errRaw
}
