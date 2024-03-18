package config

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func ConnectDB() {
	var err error
	DB, err = sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/HealthTrack?parseTime=true")
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("Connected to Database")
}
