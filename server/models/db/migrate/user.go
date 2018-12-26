package migrate

import (
	"server/helpers"
	"server/structs/db"
	"strings"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

// AlterUserAddID ...
func AlterUserAddID() {
	var user db.User

	o := orm.NewOrm()
	o.Using("default")

	// Create temp user table
	qb := []string{
		"CREATE TABLE IF NOT EXISTS",
		"temp_table",
		"(",
		"id integer NOT NULL PRIMARY KEY AUTOINCREMENT,",
		"employee_number integer NOT NULL UNIQUE,",
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
		helpers.CheckErr("error creating table temp table @AlterUserAddID", err)
	}
	beego.Debug(res)

	//Migrate data from user to temp table
	qb = []string{
		"INSERT INTO temp_table (",
		"id,",
		"employee_number,",
		"name, gender,",
		"position,",
		"start_working_date,",
		"mobile_phone,",
		"email,password,role,",
		"supervisor_id,",
		"created_at,",
		"updated_at)",
		"SELECT",
		"id,",
		"employee_number,",
		"name, gender,",
		"position,",
		"start_working_date,",
		"mobile_phone,",
		"email,password,role,",
		"supervisor_id,",
		"created_at,",
		"updated_at",
		"FROM " + user.TableName() + ";",
	}

	sql = strings.Join(qb, " ")
	beego.Debug(sql)

	res, err = o.Raw(sql).Exec()
	if err != nil {
		helpers.CheckErr("error insert from temp table @AlterUserAddID", err)
	}
	beego.Debug(res)

	// Drop table user
	qb = []string{
		"DROP TABLE " + user.TableName() + ";",
	}

	sql = strings.Join(qb, " ")
	beego.Debug(sql)

	res, err = o.Raw(sql).Exec()
	if err != nil {
		helpers.CheckErr("error drop table @AlterUserAddID", err)
	}
	beego.Debug(res)

	// rename temp table to user
	qb = []string{
		"ALTER TABLE temp_table RENAME TO " + user.TableName() + ";",
	}

	sql = strings.Join(qb, " ")
	beego.Debug(sql)

	res, err = o.Raw(sql).Exec()
	if err != nil {
		helpers.CheckErr("error rename table @AlterUserAddID", err)
	}
	beego.Debug(res)
}
