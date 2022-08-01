package main

import (
	"database/sql"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

type User struct {
	Name         string
	Email        string
	Age          uint8
	Birthday     *time.Time
	MemberNumber sql.NullString
	gorm.Model
}

func main() {
	dsn := "root:root@tcp(127.0.0.1:3306)/data_shop?charset=utf8mb4&parseTime=True&loc=Local"

	dbDriver, _ := sql.Open("mysql", dsn)
	db, _ := gorm.Open(mysql.New(mysql.Config{
		Conn: dbDriver,
	}), &gorm.Config{})
	now := time.Now()
	user := User{Name: "Jinzhu", Age: 18, Birthday: &now}

	//创建
	result := db.Create(&user)
	fmt.Println(result.RowsAffected)
	db.Select("Name", "Age", "CreatedAt").Create(&user)
	//sql查询
	sql := "select * from users where id = ?"
	var user01 User
	db.Raw(sql, 38).Scan(&user01)
	fmt.Println(user01)
	db.Model(&User{}).Where("", "").Select("").Row()
}
