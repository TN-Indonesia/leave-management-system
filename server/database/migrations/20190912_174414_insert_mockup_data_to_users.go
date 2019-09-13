package main

import (
	"server/helpers/constant/tablename"

	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type InsertMockupDataToUsers_20190912_174414 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &InsertMockupDataToUsers_20190912_174414{}
	m.Created = "20190912_174414"

	migration.Register("InsertMockupDataToUsers_20190912_174414", m)
}

// Run the migrations
func (m *InsertMockupDataToUsers_20190912_174414) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	dt := "20190912"
	m.SQL(GetQuerySeeder(tablename.Users, dt, "insert_data"))

}

// Reverse the migrations
func (m *InsertMockupDataToUsers_20190912_174414) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update

}
