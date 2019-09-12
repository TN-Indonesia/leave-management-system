package main

import (
	"log"
	"server/models/db"

	_ "server/routers"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/plugins/cors"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowAllOrigins: true,
		AllowOrigins:    []string{"*"},
		AllowMethods:    []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:    []string{"Origin", "Authorization", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers"},
		ExposeHeaders:   []string{"Origin", "Content-Type", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers"},
	}))

	beego.Run()
}

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	} else {
		//db.RegisterSQLite()
		db.RegisterPGSQL()
	}
}
