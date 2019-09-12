package main

import (
	"server/helpers/constant/tablename"

	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type CreateTableSequenceLeaveRequest_20190912_122631 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &CreateTableSequenceLeaveRequest_20190912_122631{}
	m.Created = "20190912_122631"

	migration.Register("CreateTableSequenceLeaveRequest_20190912_122631", m)
}

// Run the migrations
func (m *CreateTableSequenceLeaveRequest_20190912_122631) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	dt := "20180629"
	m.SQL(GetQuery(tablename.LeaveRequest, dt, "create_sequence"))
}

// Reverse the migrations
func (m *CreateTableSequenceLeaveRequest_20190912_122631) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	dt := "20180629"
	m.SQL(GetQuery(tablename.LeaveRequest, dt, "drop_sequence"))
}
