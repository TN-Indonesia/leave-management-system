package main

import (
	"server/helpers/constant/tablename"

	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type CreateTableLeaveRequest_20190912_120625 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &CreateTableLeaveRequest_20190912_120625{}
	m.Created = "20190912_120625"

	migration.Register("CreateTableLeaveRequest_20190912_120625", m)
}

// Run the migrations
func (m *CreateTableLeaveRequest_20190912_120625) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	dt := "20180607"
	m.SQL(GetQuery(tablename.LeaveRequest, dt, "create_table"))
}

// Reverse the migrations
func (m *CreateTableLeaveRequest_20190912_120625) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	dt := "20180607"
	m.SQL(GetQuery(tablename.LeaveRequest, dt, "drop_table"))
}
