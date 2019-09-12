package main

import (
	"server/helpers/constant/tablename"

	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type AddFkToTypeLeaveId_20190912_122726 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &AddFkToTypeLeaveId_20190912_122726{}
	m.Created = "20190912_122726"

	migration.Register("AddFkToTypeLeaveId_20190912_122726", m)
}

// Run the migrations
func (m *AddFkToTypeLeaveId_20190912_122726) Up() {
	dt := "20190912"
	m.SQL(GetQuery(tablename.TypeLeave, dt, "add_fk"))

}

// Reverse the migrations
func (m *AddFkToTypeLeaveId_20190912_122726) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update

}
