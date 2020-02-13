package main

import (
	"server/helpers/constant/tablename"

	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type AlterTableMealRequest_20191009_111610 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &AlterTableMealRequest_20191009_111610{}
	m.Created = "20191009_111610"

	migration.Register("AlterTableMealRequest_20191009_111610", m)
}

// Run the migrations
func (m *AlterTableMealRequest_20191009_111610) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	dt := "20191009"
	m.SQL(GetQuery(tablename.MealRequest, dt, "alter_table"))
}

// Reverse the migrations
func (m *AlterTableMealRequest_20191009_111610) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update

}
