package adapter

import (
	"os"
	"server/helpers/constant"
)

// CallPGSQL ...
func CallPGSQL() string {
	dbUser := os.Getenv("db_user")
	dbPwd := os.Getenv("db_password")
	dbName := os.Getenv("db_name")
	dbHost := os.Getenv("db_host")
	dbPort := os.Getenv("db_port")
	dbString := "user=" + dbUser + " password=" + dbPwd + " host=" + dbHost + " port=" + dbPort + " dbname=" + dbName + " sslmode=disable"
	return dbString
}

// CallPGSQLMeal ...
func CallPGSQLMeal() string {
	dbUser := os.Getenv("dbMeal_user")
	dbPwd := os.Getenv("dbMeal_password")
	dbName := os.Getenv("dbMeal_name")
	dbHost := os.Getenv("dbMeal_host")
	dbPort := os.Getenv("dbMeal_port")
	dbString := "user=" + dbUser + " password=" + dbPwd + " host=" + dbHost + " port=" + dbPort + " dbname=" + dbName + " sslmode=disable"
	return dbString
}

// CallSQLITE ...
func CallSQLITE() string {
	return constant.GOPATH + "/src/" + constant.GOAPP + "/database/sqlite/db_leave_request.db"
}
