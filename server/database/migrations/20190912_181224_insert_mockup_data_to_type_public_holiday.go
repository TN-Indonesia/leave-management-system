package main

import (
	"server/helpers/constant/tablename"

	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type InsertMockupDataToTypePublicHoliday_20190912_181224 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &InsertMockupDataToTypePublicHoliday_20190912_181224{}
	m.Created = "20190912_181224"

	migration.Register("InsertMockupDataToTypePublicHoliday_20190912_181224", m)
}

// Run the migrations
func (m *InsertMockupDataToTypePublicHoliday_20190912_181224) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	dt := "20190912"
	m.SQL(GetQuerySeeder(tablename.PublicHoliday, dt, "insert_data"))
}

// Reverse the migrations
func (m *InsertMockupDataToTypePublicHoliday_20190912_181224) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update

}
