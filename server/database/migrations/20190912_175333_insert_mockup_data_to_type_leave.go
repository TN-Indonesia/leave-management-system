package main

import (
	"server/helpers/constant/tablename"

	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type InsertMockupDataToTypeLeave_20190912_175333 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &InsertMockupDataToTypeLeave_20190912_175333{}
	m.Created = "20190912_175333"

	migration.Register("InsertMockupDataToTypeLeave_20190912_175333", m)
}

// Run the migrations
func (m *InsertMockupDataToTypeLeave_20190912_175333) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	dt := "20190912"
	m.SQL(GetQuerySeeder(tablename.TypeLeave, dt, "insert_data"))
}

// Reverse the migrations
func (m *InsertMockupDataToTypeLeave_20190912_175333) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update

}
