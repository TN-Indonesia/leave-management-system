package main

import (
	"server/helpers/constant/tablename"

	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type CreateTablePublicHoliday_20190912_142419 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &CreateTablePublicHoliday_20190912_142419{}
	m.Created = "20190912_142419"

	migration.Register("CreateTablePublicHoliday_20190912_142419", m)
}

// Run the migrations
func (m *CreateTablePublicHoliday_20190912_142419) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	dt := "20190912"
	m.SQL(GetQuery(tablename.PublicHoliday, dt, "create_table"))
}

// Reverse the migrations
func (m *CreateTablePublicHoliday_20190912_142419) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	dt := "20190912"
	m.SQL(GetQuery(tablename.PublicHoliday, dt, "drop_table"))
}
