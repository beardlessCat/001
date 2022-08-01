package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"math/rand"
)

var dbInfo *sql.DB

type user struct {
	id, age int
	name    string
}

func init() {
	initMySql()
}

func initMySql() (err error) {
	dsn := "root:root@tcp(127.0.0.1:3306)/data_shop"
	dbInfo, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("dataSource open failed:", err)
		return err
	}
	err = dbInfo.Ping()
	if err != nil {
		log.Fatal("dataSource Ping failed:", err)
		return err
	}
	return nil
}

func main() {
	defer dbInfo.Close()
	insert()
	fmt.Println("=================insertEnd==========================")
	update()
	fmt.Println("=================updateEnd==========================")
	deleteUser()
	fmt.Println("=================deleteEnd==========================")
	queryAll()
	fmt.Println("=================queryAllEnd==========================")
	queryById()
	fmt.Println("=================queryByIdEnd==========================")

}

func deleteUser() {
	sql := "delete from user where id = ?"
	stmt, err := dbInfo.Prepare(sql)
	if err != nil {
		fmt.Printf("prepare failed, err:%v\n", err)
	}
	defer stmt.Close()
	exec, err := stmt.Exec(2)
	if err != nil {
		fmt.Println("delete error：", err)
		return
	}
	id, err := exec.LastInsertId()
	if err != nil {
		fmt.Printf("get lastinsert ID failed, err:%v\n", err)
		return
	}
	fmt.Printf("delete success, the id is %d.\n", id)
}

func update() {
	sql := "update user set age = ? where id = ?"
	stmt, err := dbInfo.Prepare(sql)
	if err != nil {
		fmt.Printf("prepare failed, err:%v\n", err)
	}
	defer stmt.Close()
	exec, err := stmt.Exec(rand.Int(), 1)
	if err != nil {
		fmt.Println("update error：", err)
		return
	}
	id, err := exec.LastInsertId()
	if err != nil {
		fmt.Printf("get lastinsert ID failed, err:%v\n", err)
		return
	}
	fmt.Printf("update success, the id is %d.\n", id)
}

func insert() {
	sql := "insert into user(name,age)values(?,?)"
	stmt, err := dbInfo.Prepare(sql)
	if err != nil {
		fmt.Printf("prepare failed, err:%v\n", err)
	}
	defer stmt.Close()
	exec, err := stmt.Exec("jim", 20)
	if err != nil {
		fmt.Println("insert err:", err)
		return
	}
	id, err := exec.LastInsertId()
	if err != nil {
		fmt.Printf("get lastinsert ID failed, err:%v\n", err)
		return
	}
	fmt.Printf("insert success, the id is %d.\n", id)
}

func queryById() {
	//查询
	var user user
	sql := "select * from user where id = ? "
	stmt, err := dbInfo.Prepare(sql)
	if err != nil {
		fmt.Printf("prepare failed, err:%v\n", err)
	}
	defer stmt.Close()
	err = stmt.QueryRow(1).Scan(&user.id, &user.name, &user.age)
	if err != nil {
		log.Fatalln("sql query error:", err)
		return
	}
	fmt.Printf("id:%d,name:%s,age:%d\n", user.id, user.name, user.age)
}

func queryAll() {
	sqlAll := "select * from user"
	stmt, err := dbInfo.Prepare(sqlAll)
	if err != nil {
		fmt.Printf("prepare failed, err:%v\n", err)
		return
	}
	defer stmt.Close()
	rows, err := stmt.Query()
	if err != nil {
		fmt.Printf("query failed, err:%v\n", err)
		return
	}
	defer rows.Close()
	for rows.Next() {
		var u user
		err2 := rows.Scan(&u.id, &u.name, &u.age)
		if err2 != nil {
			fmt.Println("query error:", err2)
		}
		fmt.Printf("id:%d,name:%s,age:%d\n", u.id, u.name, u.age)
	}
}
