package main

import (
	"io/ioutil"
	"log"
	"server/helpers/constant"
)

// GetQuery ...
func GetQuery(table string, dt string, dbfile string) string {
	fl := constant.GOPATH + "/src/" + constant.GOAPP +
		"/database/db/" + table + "/" + dt + "/" + dbfile + ".sql"
	raw, err := ioutil.ReadFile(fl)
	if err != nil {
		log.Println("failed Open File SQL")
		panic(err)
	}
	query := string(raw)
	return query
}

// GetQuerySeeder ...
func GetQuerySeeder(table string, dt string, dbfile string) string {
	fl := constant.GOPATH + "/src/" + constant.GOAPP +
		"/database/seeders/" + table + "/" + dt + "/" + dbfile + ".sql"
	raw, err := ioutil.ReadFile(fl)
	if err != nil {
		log.Println("failed Open File SQL")
		panic(err)
	}
	query := string(raw)
	return query
}
