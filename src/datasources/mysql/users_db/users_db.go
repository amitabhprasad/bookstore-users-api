package users_db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

// const (
// 	mysql_username = ""
// 	mysql_password = ""
// 	mysql_host     = ""
// 	mysql_schema   = ""
// )

var (
	Client   *sql.DB
	username = os.Getenv("mysql_username")
	password = os.Getenv("mysql_password")
	host     = os.Getenv("mysql_host")
	schema   = os.Getenv("mysql_schema")
)

func init() {
	fmt.Printf("Connection to DB %s on host %s using user %s  \n", schema, host, username)
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8",
		username, password, host, schema,
	)
	var err error
	Client, err = sql.Open("mysql", dataSourceName)
	if err != nil {
		panic(err)
	}
	if err := Client.Ping(); err != nil {
		panic(err)
	}
	log.Println("connected to DB successfully")
}
