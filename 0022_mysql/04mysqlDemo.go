package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

/*
*
事务
*/
var dbSetting *sql.DB

func main() {
	err := initDbSet()
	if err != nil {
		fmt.Println(err)
		return
	}
	defer dbSetting.Close()
	tx, err := dbSetting.Begin()
	if err != nil {
		if tx != nil {
			tx.Rollback()
		}
	}
	sql := "insert into user(name,age) values (?,?)"
	stmt, err := dbSetting.Prepare(sql)
	if err != nil {
		fmt.Println("prepare error:", err)
		tx.Rollback()
		return
	}
	result, err := stmt.Exec("cook", 20)
	if err != nil {
		fmt.Println("exec error:", err)
		tx.Rollback()
		return
	}
	id, err := result.LastInsertId()
	if err != nil {
		fmt.Println("gtet error:", err)
		tx.Rollback()
		return
	}
	fmt.Printf("exec success!,id is: %d\n", id)
	tx.Commit()
}

func initDbSet() (err error) {
	url := "root:root@tcp(127.0.0.1:3306)/data_shop"
	dbSetting, err = sql.Open("mysql", url)
	if err != nil {
		fmt.Println("open error:", err)
		return err
	}
	err = dbSetting.Ping()
	if err != nil {
		fmt.Println("ping error:", err)
		return err
	}
	return nil
}
