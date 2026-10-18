package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/massaome/go-intermediate/myapi/controllers"
	"github.com/massaome/go-intermediate/myapi/routers"
	"github.com/massaome/go-intermediate/myapi/services"

	_ "github.com/go-sql-driver/mysql"
)

var (
	dbUser     = os.Getenv("MYSQL_USER")
	dbPassword = os.Getenv("MYSQL_PASSWORD")
	dbHost     = os.Getenv("MYSQL_HOST")
	dbPort     = os.Getenv("MYSQL_PORT")
	dbDatabase = os.Getenv("MYSQL_DATABASE")
	dbConn     = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		dbUser, dbPassword, dbHost, dbPort, dbDatabase)
)

func main() {
	db, err := sql.Open("mysql", dbConn)
	if err != nil {
		log.Println("fail to connect DB")
		return
	}

	ser := services.NewMyAppService(db)
	con := controllers.NewMyAppController(ser)

	r := routers.NewRouter(con)

	log.Println("server start at port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
