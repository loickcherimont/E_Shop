package main

import (
	"database/sql"
	"log"
	// "os"

	"github.com/go-sql-driver/mysql"
)

var db *sql.DB

func setupDatabase() {
	DBName := "e_shop"

	// cfg := mysql.Config{
	// 	User: os.Getenv("DBUSER"),
	// 	Passwd: os.Getenv("DBPASS"),
	// 	Net: "tcp",
	// 	Addr: "127.0.0.1:3306",
	// 	DBName: "e_shop",
	// }
	

	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/e_shop")
	if err != nil {
		log.Fatal(err.Error())
	}

	if pingErr := db.Ping(); pingErr != nil {
		log.Fatal(pingErr)
	}
	log.Printf("Connected to the database %v", DBName)

	defer db.Close()

}
