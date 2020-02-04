package main

import (
	"server/helpers/constant/tablename"
	
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type IndexLeaverequest_20200204_110729 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &IndexLeaverequest_20200204_110729{}
	m.Created = "20200204_110729"

	migration.Register("IndexLeaverequest_20200204_110729", m)
}

//Up Run the migrations
func (m *IndexLeaverequest_20200204_110729) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
		// use m.SQL("CREATE TABLE ...") to make schema update
		dt := "20200204"
		m.SQL(GetQuery(tablename.LeaveRequest, dt, "create_index"))

}

//Down Reverse the migrations
func (m *IndexLeaverequest_20200204_110729) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update

}
