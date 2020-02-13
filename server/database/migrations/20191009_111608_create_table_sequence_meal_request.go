package main

import (
	"server/helpers/constant/tablename"

	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type CreateTableSequenceMealRequest_20191009_111608 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &CreateTableSequenceMealRequest_20191009_111608{}
	m.Created = "20191009_111608"

	migration.Register("CreateTableSequenceMealRequest_20191009_111608", m)
}

// Run the migrations
func (m *CreateTableSequenceMealRequest_20191009_111608) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	dt := "20191009"
	m.SQL(GetQuery(tablename.MealRequest, dt, "create_sequence"))
}

// Reverse the migrations
func (m *CreateTableSequenceMealRequest_20191009_111608) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	dt := "20191009"
	m.SQL(GetQuery(tablename.MealRequest, dt, "create_sequence"))
}
