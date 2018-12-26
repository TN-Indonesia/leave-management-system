package admin

import (
	"errors"
	"server/helpers"
	"server/helpers/constant"
	structDB "server/structs/db"
	structLogic "server/structs/logic"
	"strings"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

// Admin ...
type Admin struct{}

// AddUser ...
func (u *Admin) AddUser(user structDB.User) (int64, error) {
	o := orm.NewOrm()

	// _, err := o.Insert(&user)
	// if err != nil {
	// 	helpers.CheckErr("Error insert @AddUser", err)
	// 	return errors.New("Insert users failed")
	// }

	qb, err := orm.NewQueryBuilder("mysql")
	if err != nil {
		helpers.CheckErr("Query builder failed @AddUser", err)
		return 0, err
	}

	qb.InsertInto(user.TableName(),
		"employee_number",
		"name",
		"gender",
		"position",
		"start_working_date",
		"mobile_phone",
		"email",
		"password",
		"role",
		"supervisor_id",
		"updated_at",
	).Values("?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?")
	sql := qb.String()

	res, err := o.Raw(sql,
		user.EmployeeNumber,
		user.Name,
		user.Gender,
		user.Position,
		user.StartWorkingDate,
		user.MobilePhone,
		user.Email,
		user.Password,
		user.Role,
		user.SupervisorID,
		user.UpdatedAt,
	).Exec()

	if err != nil {
		helpers.CheckErr("Error insert data User @AddUser", err)
		return 0, errors.New("Insert users failed")
	}

	_, err = res.RowsAffected()
	if err != nil {
		helpers.CheckErr("Error get rows affected @AddUser", err)
		return 0, err
	}

	lastRowID, err := res.LastInsertId()
	if err != nil {
		helpers.CheckErr("Error get last Insert ID", err)
		return 0, err
	}

	return lastRowID, err
}

// GetUsers ...
func (u *Admin) GetUsers() (users []structDB.User, err error) {
	var (
		dbUser structDB.User
		roles  []string
	)
	roles = append(roles, "employee", "supervisor", "director")

	o := orm.NewOrm()
	count, err := o.Raw("SELECT * FROM "+dbUser.TableName()+" WHERE role IN (?,?,?)", roles).QueryRows(&users)
	if err != nil {
		helpers.CheckErr("Failed get users @GetUsers", err)
		return users, err
	}
	beego.Debug("Total user =", count)

	return users, err
}

// GetUser ...
func (u *Admin) GetUser(employeeNumber int64) (user structDB.User, err error) {
	o := orm.NewOrm()
	qb, errQB := orm.NewQueryBuilder("mysql")
	if errQB != nil {
		helpers.CheckErr("Query builder failed @GetUser", errQB)
		return user, errQB
	}

	qb.Select("*").From(user.TableName()).
		Where(`id = ? `)
	qb.Limit(1)
	sql := qb.String()

	errRaw := o.Raw(sql, employeeNumber).QueryRow(&user)
	if errRaw != nil {
		helpers.CheckErr("Failed query select item @GetUser", errRaw)
		return user, errors.New("Employee number not exist")
	}

	return user, err
}

// DeleteUser ...
func (u *Admin) DeleteUser(ID int64) (err error) {
	o := orm.NewOrm()
	v := structDB.User{ID: ID}

	err = o.Read(&v)
	if err == nil {
		var num int64
		if num, err = o.Delete(&structDB.User{ID: ID}); err == nil {
			beego.Debug("Number of records deleted in database:", num)
		} else if err != nil {
			helpers.CheckErr("Error delete user @DeleteUser", err)
			return errors.New("Error delete user")
		}
	}
	if err != nil {
		helpers.CheckErr("Error delete user @DeleteUser", err)
		return errors.New("Delete failed, id not exist")
	}

	return err
}

// UpdateUser ...
func (u *Admin) UpdateUser(e *structDB.User, ID int64) (err error) {
	var (
		user  structLogic.GetEmployee
		count int
	)

	o := orm.NewOrm()
	qb, errQB := orm.NewQueryBuilder("mysql")
	if errQB != nil {
		helpers.CheckErr("Query builder failed @UpdateUser", errQB)
		return errQB
	}

	o.Raw(`SELECT name, email FROM `+e.TableName()+` WHERE id = ?`, ID).QueryRow(&user)

	if e.Email != user.Email {
		o.Raw(`SELECT count(*) as Count FROM `+e.TableName()+` WHERE email = ?`, e.Email).QueryRow(&count)
		if count > 0 {
			return errors.New("Email already register")
		} else {
			qb.Update(e.TableName()).
				Set("employee_number = ?",
					"name = ?",
					"gender = ?",
					"position = ?",
					"start_working_date = ?",
					"mobile_phone = ?",
					"email= ?",
					"role = ?",
					"supervisor_id = ?",
					"updated_at = ?").Where("id = ? ")
			sql := qb.String()

			e.Email = strings.ToLower(e.Email)

			res, errRaw := o.Raw(sql,
				e.EmployeeNumber,
				e.Name,
				e.Gender,
				e.Position,
				e.StartWorkingDate,
				e.MobilePhone,
				e.Email,
				e.Role,
				e.SupervisorID,
				e.UpdatedAt,
				ID).Exec()

			if errRaw != nil {
				helpers.CheckErr("Error update user @UpdateUser", errRaw)
				return errors.New("Update user failed")
			}

			_, errRow := res.RowsAffected()
			if errRow != nil {
				helpers.CheckErr("Error get rows affected @UpdateUser", errRow)
				return errRow
			}
		}
	} else {
		qb.Update(e.TableName()).
			Set("employee_number = ?",
				"name = ?",
				"gender = ?",
				"position = ?",
				"start_working_date = ?",
				"mobile_phone = ?",
				"email= ?",
				"role = ?",
				"supervisor_id = ?",
				"updated_at = ?").Where("id = ? ")
		sql := qb.String()

		e.Email = strings.ToLower(e.Email)

		res, errRaw := o.Raw(sql,
			e.EmployeeNumber,
			e.Name,
			e.Gender,
			e.Position,
			e.StartWorkingDate,
			e.MobilePhone,
			e.Email,
			e.Role,
			e.SupervisorID,
			e.UpdatedAt,
			ID).Exec()

		if errRaw != nil {
			helpers.CheckErr("Error update user @UpdateUser", errRaw)
			return errors.New("Update user failed")
		}

		_, errRow := res.RowsAffected()
		if errRow != nil {
			helpers.CheckErr("Error get rows affected @UpdateUser", errRow)
			return errRow
		}
	}

	return err
}

// GetLeaveRequestPending ...
func (u *Admin) GetLeaveRequestPending() (reqPending []structLogic.RequestPending, err error) {
	var (
		user          structDB.User
		leave         structDB.LeaveRequest
		typeLeave     structDB.TypeLeave
		userTypeLeave structDB.UserTypeLeave
	)

	o := orm.NewOrm()
	qb, errQB := orm.NewQueryBuilder("mysql")
	if errQB != nil {
		helpers.CheckErr("Query builder failed @GetLeaveRequestPending", errQB)
		return reqPending, errQB
	}

	qb.Select(
		leave.TableName()+".id",
		user.TableName()+".id as employee_id",
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
		leave.TableName()+".half_dates",
		leave.TableName()+".total",
		leave.TableName()+".back_on",
		leave.TableName()+".contact_address",
		leave.TableName()+".contact_number",
		leave.TableName()+".status",
		leave.TableName()+".action_by",
		leave.TableName()+".notes").
		From(user.TableName()).
		InnerJoin(leave.TableName()).
		On(user.TableName() + ".id" + "=" + leave.TableName() + ".employee_number").
		InnerJoin(typeLeave.TableName()).
		On(typeLeave.TableName() + ".id" + "=" + leave.TableName() + ".type_leave_id").
		InnerJoin(userTypeLeave.TableName()).
		On(userTypeLeave.TableName() + ".type_leave_id" + "=" + leave.TableName() + ".type_leave_id").
		And(userTypeLeave.TableName() + ".employee_number" + "=" + leave.TableName() + ".employee_number").
		Where(`(status = ? OR status = ? )`).
		OrderBy(leave.TableName() + ".created_at DESC")
	sql := qb.String()

	statPendingInSupervisor := constant.StatusPendingInSupervisor
	statPendingInDirector := constant.StatusPendingInDirector

	count, errRaw := o.Raw(sql, statPendingInSupervisor, statPendingInDirector).QueryRows(&reqPending)
	if errRaw != nil {
		helpers.CheckErr("Failed query select @GetLeaveRequestPending", errRaw)
		return reqPending, errors.New("Error get leave request pending")
	}
	beego.Debug("Total pending request =", count)

	return reqPending, errRaw
}

// GetLeaveRequestApproved ...
func (u *Admin) GetLeaveRequestApproved() (reqApprove []structLogic.RequestAccept, err error) {
	var (
		user          structDB.User
		leave         structDB.LeaveRequest
		typeLeave     structDB.TypeLeave
		userTypeLeave structDB.UserTypeLeave
	)

	o := orm.NewOrm()
	qb, errQB := orm.NewQueryBuilder("mysql")
	if errQB != nil {
		helpers.CheckErr("Query builder failed @GetLeaveRequestApproved", errQB)
		return reqApprove, errQB
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
		leave.TableName()+".half_dates",
		leave.TableName()+".total",
		leave.TableName()+".back_on",
		leave.TableName()+".contact_address",
		leave.TableName()+".contact_number",
		leave.TableName()+".status",
		leave.TableName()+".action_by",
		leave.TableName()+".notes").
		From(user.TableName()).
		InnerJoin(leave.TableName()).
		On(user.TableName() + ".id" + "=" + leave.TableName() + ".employee_number").
		InnerJoin(typeLeave.TableName()).
		On(typeLeave.TableName() + ".id" + "=" + leave.TableName() + ".type_leave_id").
		InnerJoin(userTypeLeave.TableName()).
		On(userTypeLeave.TableName() + ".type_leave_id" + "=" + leave.TableName() + ".type_leave_id").
		And(userTypeLeave.TableName() + ".employee_number" + "=" + leave.TableName() + ".employee_number").
		Where(`status = ? `).
		OrderBy(leave.TableName() + ".created_at DESC")
	sql := qb.String()

	statApproveDirector := constant.StatusSuccessInDirector

	count, errRaw := o.Raw(sql, statApproveDirector).QueryRows(&reqApprove)
	if errRaw != nil {
		helpers.CheckErr("Failed query select @GetLeaveRequestApproved", errRaw)
		return reqApprove, errors.New("Error get leave request approved")
	}
	beego.Debug("Total approved request =", count)

	return reqApprove, errRaw
}

// GetLeaveRequestRejected ...
func (u *Admin) GetLeaveRequestRejected() (reqReject []structLogic.RequestReject, err error) {
	var (
		user          structDB.User
		leave         structDB.LeaveRequest
		typeLeave     structDB.TypeLeave
		userTypeLeave structDB.UserTypeLeave
	)

	o := orm.NewOrm()
	qb, errQB := orm.NewQueryBuilder("mysql")
	if errQB != nil {
		helpers.CheckErr("Query builder failed @GetLeaveRequestRejected", errQB)
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
		leave.TableName()+".half_dates",
		leave.TableName()+".total",
		leave.TableName()+".back_on",
		leave.TableName()+".contact_address",
		leave.TableName()+".contact_number",
		leave.TableName()+".status",
		leave.TableName()+".reject_reason",
		leave.TableName()+".action_by",
		leave.TableName()+".notes").
		From(user.TableName()).
		InnerJoin(leave.TableName()).
		On(user.TableName() + ".id" + "=" + leave.TableName() + ".employee_number").
		InnerJoin(typeLeave.TableName()).
		On(typeLeave.TableName() + ".id" + "=" + leave.TableName() + ".type_leave_id").
		InnerJoin(userTypeLeave.TableName()).
		On(userTypeLeave.TableName() + ".type_leave_id" + "=" + leave.TableName() + ".type_leave_id").
		And(userTypeLeave.TableName() + ".employee_number" + "=" + leave.TableName() + ".employee_number").
		Where(`(status = ? OR status = ? )`).
		OrderBy(leave.TableName() + ".created_at DESC")
	sql := qb.String()

	statRejectInSuperVisor := constant.StatusRejectInSuperVisor
	statRejectInDirector := constant.StatusRejectInDirector

	count, errRaw := o.Raw(sql, statRejectInSuperVisor, statRejectInDirector).QueryRows(&reqReject)
	if errRaw != nil {
		helpers.CheckErr("Failed query select @GetLeaveRequestRejected", errRaw)
		return reqReject, errors.New("Error get leave request reject")
	}
	beego.Debug("Total reject request =", count)

	return reqReject, errRaw
}

// ResetUserTypeLeave ...
func (u *Admin) ResetUserTypeLeave(leaveRemaining float64, typeLeaveID int64) error {
	var typeLeave structDB.UserTypeLeave

	o := orm.NewOrm()

	qb, errQB := orm.NewQueryBuilder("mysql")
	if errQB != nil {
		helpers.CheckErr("Query builder failed @ResetUserTypeLeave", errQB)
		return errQB
	}

	qb.Update(typeLeave.TableName()).
		Set("leave_remaining = ?").
		Where("type_leave_id = ?")
	sql := qb.String()

	res, errRaw := o.Raw(sql, leaveRemaining, typeLeaveID).Exec()
	if errRaw != nil {
		helpers.CheckErr("Error update leave balance @ResetUserTypeLeave", errRaw)
		return errors.New("Reset leave balance failed")
	}

	_, errRow := res.RowsAffected()
	helpers.CheckErr("Error get rows affected @ResetUserTypeLeave", errRow)

	return errRow
}

// UpdateUserTypeLeave ...
func (u *Admin) UpdateUserTypeLeave(
	leaveRemaining float64,
	typeLeaveID int64,
	employeeNumber int64,
) error {
	var typeLeave structDB.UserTypeLeave

	o := orm.NewOrm()

	qb, errQB := orm.NewQueryBuilder("mysql")
	if errQB != nil {
		helpers.CheckErr("Query builder failed @UpdateUserTypeLeave", errQB)
		return errQB
	}

	qb.Update(typeLeave.TableName()).
		Set(
			"leave_remaining = ?",
			"type_leave_id = ?",
		).
		Where("employee_number = ?")
	sql := qb.String()

	res, errRaw := o.Raw(sql, leaveRemaining, typeLeaveID, employeeNumber).Exec()
	if errRaw != nil {
		helpers.CheckErr("Error update @UpdateUserTypeLeave", errRaw)
		return errors.New("Update request failed")
	}

	_, errRow := res.RowsAffected()
	helpers.CheckErr("Error get rows affected @UpdateUserTypeLeave", errRow)

	return errRow
}
