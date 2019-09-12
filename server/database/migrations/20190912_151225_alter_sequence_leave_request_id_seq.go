package main

import (
	"server/helpers/constant/tablename"

	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type AlterSequenceLeaveRequestIdSeq_20190912_151225 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &AlterSequenceLeaveRequestIdSeq_20190912_151225{}
	m.Created = "20190912_151225"

	migration.Register("AlterSequenceLeaveRequestIdSeq_20190912_151225", m)
}

// Run the migrations
func (m *AlterSequenceLeaveRequestIdSeq_20190912_151225) Up() {
	dt := "20190912"
	m.SQL(GetQuery(tablename.TypeLeave, dt, "alter_sequence"))

}

// Reverse the migrations
func (m *AlterSequenceLeaveRequestIdSeq_20190912_151225) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update

}
