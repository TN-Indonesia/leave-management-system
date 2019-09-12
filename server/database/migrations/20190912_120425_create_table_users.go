package main

import (
	"server/helpers/constant/tablename"

	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type CreateTableUsers_20190912_120425 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &CreateTableUsers_20190912_120425{}
	m.Created = "20190912_120425"

	migration.Register("CreateTableUsers_20190912_120425", m)
}

// Run the migrations
func (m *CreateTableUsers_20190912_120425) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	dt := "20190912"
	m.SQL(GetQuery(tablename.Users, dt, "create_table"))
}

// Reverse the migrations
func (m *CreateTableUsers_20190912_120425) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	dt := "20190912"
	m.SQL(GetQuery(tablename.Users, dt, "drop_table"))
}
