package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	dns := "root:root@tcp(127.0.0.1:3306)/data_shop"
	db, err := sql.Open("mysql", dns)
	if err != nil {
		fmt.Println("db open error:", err)
		panic(err)
	}
	defer db.Close()
}
