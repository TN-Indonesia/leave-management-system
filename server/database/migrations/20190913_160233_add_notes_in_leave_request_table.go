package main

import (
	"server/helpers/constant/tablename"

	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type AddNotesInLeaveRequestTable_20190913_160233 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &AddNotesInLeaveRequestTable_20190913_160233{}
	m.Created = "20190913_160233"

	migration.Register("AddNotesInLeaveRequestTable_20190913_160233", m)
}

// Run the migrations
func (m *AddNotesInLeaveRequestTable_20190913_160233) Up() {
	dt := "20190913"
	m.SQL(GetQuery(tablename.LeaveRequest, dt, "alter_table"))

}

// Reverse the migrations
func (m *AddNotesInLeaveRequestTable_20190913_160233) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update

}
