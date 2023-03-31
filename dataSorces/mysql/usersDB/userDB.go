package usersDB

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var (
	Client *sql.DB
)

func init() {
	username := "book_user"
	password := "book_pass"
	host := "127.0.0.1"
	port := "4306"
	db := "book_db"

	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8", username, password, host, port, db)
	fmt.Println(dataSourceName)
	var err error
	Client, err = sql.Open("mysql", dataSourceName)
	if err != nil {
		panic(err)
	}

	Client.SetMaxIdleConns(5)
	Client.SetMaxIdleConns(20)
	Client.SetConnMaxLifetime(60 * time.Minute)
	Client.SetConnMaxIdleTime(10 * time.Minute)

	if err = Client.Ping(); err != nil {
		panic(err)
	}

	log.Println("Database successfully configured!")
}
