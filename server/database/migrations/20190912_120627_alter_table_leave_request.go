package main

import (
	"server/helpers/constant/tablename"

	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type AlterTableLeaveRequest_20190917_171805 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &AlterTableLeaveRequest_20190917_171805{}
	m.Created = "20190917_171805"

	migration.Register("AlterTableLeaveRequest_20190917_171805", m)
}

// Run the migrations
func (m *AlterTableLeaveRequest_20190917_171805) Up() {
	dt := "20190912"
	m.SQL(GetQuery(tablename.LeaveRequest, dt, "alter_table"))

}

// Reverse the migrations
func (m *AlterTableLeaveRequest_20190917_171805) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update

}
