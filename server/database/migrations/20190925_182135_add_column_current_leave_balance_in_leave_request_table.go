package main

import (
	"server/helpers/constant/tablename"

	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type AddColumnCurrentLeaveBalanceInLeaveRequestTable_20190925_182135 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &AddColumnCurrentLeaveBalanceInLeaveRequestTable_20190925_182135{}
	m.Created = "20190925_182135"

	migration.Register("AddColumnCurrentLeaveBalanceInLeaveRequestTable_20190925_182135", m)
}

// Run the migrations
func (m *AddColumnCurrentLeaveBalanceInLeaveRequestTable_20190925_182135) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	dt := "20190925"
	m.SQL(GetQuery(tablename.LeaveRequest, dt, "alter_table"))
}

// Reverse the migrations
func (m *AddColumnCurrentLeaveBalanceInLeaveRequestTable_20190925_182135) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update

}
