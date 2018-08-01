package db

import (
	"encoding/json"
	"io/ioutil"
	"strings"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"

	// _ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"

	"server/adapter"
	"server/helpers"
	"server/helpers/constant"
	dbStruct "server/structs/db"
)

// RegisterPGSQL ...
func RegisterPGSQL() {
	maxIdle := 30
	maxConn := 30

	// errRegisterDriver := orm.RegisterDriver("postgres", orm.DRPostgres)
	// if errRegisterDriver != nil {
	// 	helpers.CheckErr("error while register driver @RegisterPGSQL", errRegisterDriver)
	// }

	// errRegisterDataBase := orm.RegisterDataBase("default", "postgres",
	// 	adapter.CallPGSQL(),
	// 	maxIdle, maxConn)
	// if errRegisterDataBase != nil {
	// 	helpers.CheckErr("error while register DB @RegisterPGSQL", errRegisterDataBase)
	// }

	errRegisterDriver := orm.RegisterDriver("sqlite3", orm.DRSqlite)
	if errRegisterDriver != nil {
		helpers.CheckErr("error while register driver @RegisterPGSQL", errRegisterDriver)
	}

	errRegisterDataBase := orm.RegisterDataBase("default", "sqlite3", adapter.CallSQLITE(), maxIdle, maxConn)
	if errRegisterDataBase != nil {
		helpers.CheckErr("error while register DB @RegisterPGSQL", errRegisterDataBase)
	}
	beego.Debug(adapter.CallSQLITE())

	RegisterModel()
	CreateTableUser()
	CreateTableLeaveRequest()
	CreateTableTypeLeave()
	CreateTableUserTypeLeave()

	MigrateData("users")
	MigrateData("type_leave")
}

// RegisterModel to register database
func RegisterModel() {
	orm.RegisterModel(new(dbStruct.LeaveRequest))
	orm.RegisterModel(new(dbStruct.User))
	orm.RegisterModel(new(dbStruct.TypeLeave))
	orm.RegisterModel(new(dbStruct.UserTypeLeave))
}

// CreateTableUser ...
func CreateTableUser() {
	var user dbStruct.User

	o := orm.NewOrm()
	o.Using("default")

	qb := []string{
		"CREATE TABLE IF NOT EXISTS",
		user.TableName(),
		"(",
		"employee_number integer NOT NULL PRIMARY KEY,",
		"name text NOT NULL,",
		"gender text NOT NULL,",
		"position text NOT NULL,",
		"start_working_date text NOT NULL,",
		"mobile_phone text NOT NULL,",
		"email text NOT NULL,",
		"password varchar(100) NOT NULL,",
		"role text NOT NULL,",
		"supervisor_id integer,",
		"created_at timestamp with time zone NOT NULL default CURRENT_TIMESTAMP,",
		"updated_at timestamp with time zone);",
	}

	sql := strings.Join(qb, " ")
	beego.Debug(sql)
	res, err := o.Raw(sql).Exec()

	if err != nil {
		beego.Warning("error creating table users", err)
	}

	beego.Debug(res)
}

// CreateTableLeaveRequest ...
func CreateTableLeaveRequest() {
	var leave dbStruct.LeaveRequest

	o := orm.NewOrm()
	o.Using("default")

	qb := []string{
		"CREATE TABLE IF NOT EXISTS",
		leave.TableName(),
		"(",
		"id integer NOT NULL PRIMARY KEY AUTOINCREMENT,",
		"employee_number integer NOT NULL,",
		"type_leave_id integer NOT NULL,",
		"reason text NOT NULL,",
		"date_from text NOT NULL,",
		"date_to text NOT NULL,",
		"half_dates text [],",
		"back_on text NOT NULL,",
		"total float NOT NULL,",
		"contact_address text NOT NULL,",
		"contact_number text NOT NULL,",
		"status text NOT NULL,",
		"action_by text,",
		"reject_reason text,",
		"created_at timestamp NOT NULL default CURRENT_TIMESTAMP,",
		"updated_at timestamp with time zone,",
		"FOREIGN KEY (employee_number) REFERENCES users (employee_number) ON DELETE CASCADE ON UPDATE CASCADE);",
	}

	sql := strings.Join(qb, " ")
	beego.Debug(sql)
	res, err := o.Raw(sql).Exec()

	if err != nil {
		beego.Warning("error creating table leave_request", err)
	}

	beego.Debug(res)
}

// CreateTableTypeLeave ...
func CreateTableTypeLeave() {
	var typeLeave dbStruct.TypeLeave

	o := orm.NewOrm()
	o.Using("default")

	qb := []string{
		"CREATE TABLE IF NOT EXISTS",
		typeLeave.TableName(),
		"(",
		"id integer PRIMARY KEY NOT NULL,",
		"type_name text NOT NULL);",
	}

	sql := strings.Join(qb, " ")
	beego.Debug(sql)
	res, err := o.Raw(sql).Exec()

	if err != nil {
		beego.Warning("error creating table type_leave", err)
	}

	beego.Debug(res)
}

// CreateTableUserTypeLeave ...
func CreateTableUserTypeLeave() {
	var userTypeLeave dbStruct.UserTypeLeave

	o := orm.NewOrm()
	o.Using("default")

	qb := []string{
		"CREATE TABLE IF NOT EXISTS",
		userTypeLeave.TableName(),
		"(",
		"id integer NOT NULL PRIMARY KEY AUTOINCREMENT,",
		"employee_number integer NOT NULL,",
		"type_leave_id integer NOT NULL,",
		"leave_remaining float NOT NULL,",
		"FOREIGN KEY (employee_number) REFERENCES users (employee_number) ON DELETE CASCADE ON UPDATE CASCADE,",
		"FOREIGN KEY (type_leave_id) REFERENCES leave_request (type_leave_id) ON DELETE CASCADE ON UPDATE CASCADE);",
	}

	sql := strings.Join(qb, " ")
	beego.Debug(sql)
	res, err := o.Raw(sql).Exec()

	if err != nil {
		beego.Warning("error creating table user_type_leave", err)
	}

	beego.Debug(res)
}

// MigrateData ...
func MigrateData(param string) {
	o := orm.NewOrm()

	if param == "type_leave" {
		var typeLeave []dbStruct.TypeLeave
		fl := constant.GOPATH + "src/" + constant.GOAPP + "/database/sqlite/seeders/data_type_leave.json"

		raw, err := ioutil.ReadFile(fl)
		if err != nil {
			beego.Warning("failed read file seeder", err)
		}

		err = json.Unmarshal(raw, &typeLeave)
		if err != nil {
			beego.Warning("failed unmarshall seeders", err)
		}

		cnt, errMulti := o.InsertMulti(len(typeLeave), typeLeave)
		beego.Debug(cnt, errMulti)
	} else if param == "users" {
		var users []dbStruct.User
		fl := constant.GOPATH + "src/" + constant.GOAPP + "/database/sqlite/seeders/data_admin.json"

		raw, err := ioutil.ReadFile(fl)
		if err != nil {
			beego.Warning("failed read file seeder", err)
		}

		err = json.Unmarshal(raw, &users)
		if err != nil {
			beego.Warning("failed unmarshall seeders", err)
		}

		cnt, errMulti := o.InsertMulti(len(users), users)
		beego.Debug(cnt, errMulti)
	}
}

// ResetDB ...
func ResetDB() {
	var (
		user          dbStruct.User
		leave         dbStruct.LeaveRequest
		typeLeave     dbStruct.TypeLeave
		userTypeLeave dbStruct.UserTypeLeave
	)

	o := orm.NewOrm()
	err := o.Begin()
	if err != nil {
		beego.Warning(err)
	}

	res1, errRaw1 := o.Raw(`DELETE FROM ` + user.TableName()).Exec()
	if errRaw1 != nil {
		beego.Warning("error reset users", errRaw1)
	}

	res2, errRaw2 := o.Raw(`DELETE FROM ` + leave.TableName()).Exec()
	if errRaw2 != nil {
		beego.Warning("error reset leave_request", errRaw2)
	}

	res3, errRaw3 := o.Raw(`DELETE FROM ` + typeLeave.TableName()).Exec()
	if errRaw3 != nil {
		beego.Warning("error reset type_leave", errRaw3)
	}

	res4, errRaw3 := o.Raw(`DELETE FROM ` + userTypeLeave.TableName()).Exec()
	if errRaw3 != nil {
		beego.Warning("error reset user_type_leave", errRaw3)
	}

	err = o.Commit()
	if err != nil {
		beego.Warning(err)
	}

	beego.Debug(res1, res2, res3, res4)
}
