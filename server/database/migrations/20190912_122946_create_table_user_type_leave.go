package main

import (
	"server/helpers/constant/tablename"

	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type CreateTableUserTypeLeave_20190912_122946 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &CreateTableUserTypeLeave_20190912_122946{}
	m.Created = "20190912_122946"

	migration.Register("CreateTableUserTypeLeave_20190912_122946", m)
}

// Run the migrations
func (m *CreateTableUserTypeLeave_20190912_122946) Up() {
	dt := "20190912"
	m.SQL(GetQuery(tablename.UserTypeLeave, dt, "create_table"))
}

// Reverse the migrations
func (m *CreateTableUserTypeLeave_20190912_122946) Down() {
	dt := "20190912"
	m.SQL(GetQuery(tablename.UserTypeLeave, dt, "drop_table"))
}
