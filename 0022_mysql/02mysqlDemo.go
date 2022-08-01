package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

/*
*
其中sql.DB是表示连接的数据库对象（结构体实例），
它保存了连接数据库相关的所有信息。它内部维护着一个具有零到多个底层连接的连接池，
它可以安全地被多个goroutine同时使用。
*/
var db *sql.DB

/*
*
Open函数可能只是验证其参数格式是否正确，实际上并不创建与数据库的连接。如果要检查数据源的名称是否真实有效，应该调用Ping方法。
*/
func initDb() (err error) {
	dsn := "root:root@tcp(127.0.0.1:3306)/data_shop"
	//不会校验账号密码是否正确
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		return err
	}
	//	尝试与数据库建立连接（校验dsn是否正确）
	err = db.Ping()
	if err != nil {
		return err
	}
	db.SetMaxOpenConns(100)
	db.SetMaxIdleConns(20)
	return nil
}
func main() {
	err := initDb()
	if err != nil {
		fmt.Println("init db error:", err)
		return
	}
}
