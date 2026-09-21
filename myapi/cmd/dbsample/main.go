package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/massaome/go-intermediate/myapi/models"
)

func main() {
	dbUser := os.Getenv("MYSQL_USER")
	dbPassword := os.Getenv("MYSQL_PASSWORD")
	dbHost := os.Getenv("MYSQL_HOST")
	dbPort := os.Getenv("MYSQL_PORT")
	dbDatabase := os.Getenv("MYSQL_DATABASE")

	dbConn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		dbUser, dbPassword, dbHost, dbPort, dbDatabase)

	db, err := sql.Open("mysql", dbConn)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	const sqlStr = `
		select article_id, title, contents, username, nice
		from articles;
	`
	rows, err := db.Query(sqlStr)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	articleArray := make([]models.Article, 0)
	for rows.Next() {
		var article models.Article
		err := rows.Scan(&article.ID, &article.Title, &article.Contents, &article.UserName, &article.NiceNum)
		if err != nil {
			fmt.Println(err)
		} else {
			articleArray = append(articleArray, article)
		}
	}

	fmt.Printf("%+v\n", articleArray)
}
