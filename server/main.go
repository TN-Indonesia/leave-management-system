package main

import (
	"log"
	"os"
	"server/models/db"

	_ "server/routers"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/plugins/cors"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
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
	db.RegisterPGSQL()

	// path := os.Getenv("GOPATH") + "/src/" + os.Getenv("GOAPP") + "/conf/env"
	// err := godotenv.Load(path)
	// if err != nil {
	// 	log.Fatal("Error loading .env file")
	// } else {
	// 	db.RegisterPGSQL()
	// }
}
