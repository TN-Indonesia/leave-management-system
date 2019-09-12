package main

import (
	"server/helpers/constant/tablename"

	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type CreateTableTypeLeave_20190912_122124 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &CreateTableTypeLeave_20190912_122124{}
	m.Created = "20190912_122124"

	migration.Register("CreateTableTypeLeave_20190912_122124", m)
}

// Run the migrations
func (m *CreateTableTypeLeave_20190912_122124) Up() {
	dt := "20190912"
	m.SQL(GetQuery(tablename.TypeLeave, dt, "create_table"))

}

// Reverse the migrations
func (m *CreateTableTypeLeave_20190912_122124) Down() {
	dt := "20190912"
	m.SQL(GetQuery(tablename.TypeLeave, dt, "drop_table"))
}
