package main

import (
	"server/helpers/constant/tablename"
	
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type CreateTableMealRequest_20191008_164907 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &CreateTableMealRequest_20191008_164907{}
	m.Created = "20191008_164907"

	migration.Register("CreateTableMealRequest_20191008_164907", m)
}

// Run the migrations
func (m *CreateTableMealRequest_20191008_164907) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	dt := "20191008"
	m.SQL(GetQuery(tablename.MealRequest, dt, "create_table"))
}

// Reverse the migrations
func (m *CreateTableMealRequest_20191008_164907) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	dt := "20191008"
	m.SQL(GetQuery(tablename.MealRequest, dt, "drop_table"))
}
