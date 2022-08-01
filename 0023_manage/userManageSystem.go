package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net/http"
	"strconv"
)

const CodeSuccess = "0000"
const CodeFail = "9999"
const DbUrl = "root:root@tcp(127.0.0.1:3306)/data_shop"

var db *sql.DB

type UserLoginDto struct {
	name, password string
}

func NewUserLoginDto(name string, password string) *UserLoginDto {
	return &UserLoginDto{name: name, password: password}
}

func (u *UserLoginDto) Login() (result *ApiResult) {
	sql := "select * from user where name = ? and password = ?"
	stmt, err := db.Prepare(sql)
	if err != nil {
		log.Fatalln("select sql prepare error:", err)
		result = NewApiResult(CodeFail, "登录失败！")
		return result
	}
	rows, err := stmt.Query(u.name, u.password)
	if err != nil {
		log.Fatalln("select sql query error:", err)
		result = NewApiResult(CodeFail, "登录失败！")
		return result
	}
	defer rows.Close()
	if rows.Next() {
		result = NewApiResult(CodeSuccess, "登录成功！")
	} else {
		result = NewApiResult(CodeFail, "用户名或密码不正确！")
	}
	return result
}

type UserRegisterDto struct {
	confirmPassword string
	age             int
	*UserLoginDto   //通过嵌套匿名结构体实现继承
}

func NewUserRegisterDto(confirmPassword string, age int, userLoginDto *UserLoginDto) *UserRegisterDto {
	return &UserRegisterDto{confirmPassword: confirmPassword, age: age, UserLoginDto: userLoginDto}
}

func (u UserRegisterDto) register() *ApiResult {
	var result *ApiResult
	//校验是否存在
	sql := "select * from user where name = ?"
	stmt, err := db.Prepare(sql)
	if err != nil {
		log.Fatalln("select sql prepare error:", err)
		result = NewApiResult(CodeFail, "登录失败！")
		return result
	}
	rows, err := stmt.Query(u.name)
	if err != nil {
		log.Fatalln("select sql query error:", err)
		result = NewApiResult(CodeFail, "登录失败！")
		return result
	}
	defer rows.Close()
	if rows.Next() {
		result = NewApiResult(CodeFail, "用户已存在！")
		return result
	}
	sql = "insert into user(name,age,password)values (?,?,?)"
	stmt2, err := db.Prepare(sql)
	if err != nil {
		log.Fatalln("insert sql prepare error:", err)
		result = NewApiResult(CodeFail, "注册失败！")
		return result
	}
	_, err = stmt2.Exec(u.name, u.age, u.password)
	if err != nil {
		log.Fatalln("exec sql prepare error:", err)
		result = NewApiResult(CodeFail, "注册失败！")
		return result
	}
	result = NewApiResultWithData(CodeSuccess, "注册成功！", u)
	return result
}

type ApiResult struct {
	Code    string      `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func NewApiResultWithData(code string, message string, data interface{}) *ApiResult {
	return &ApiResult{Code: code, Message: message, Data: data}
}

func NewApiResult(code string, message string) *ApiResult {
	return &ApiResult{Code: code, Message: message}
}

func (a *ApiResult) getBytes() []byte {
	data, err := json.Marshal(a)
	if err != nil {
		fmt.Println("json encode error:", err)
		return nil
	}
	return data
}

func loginHandlerFunc(writer http.ResponseWriter, request *http.Request) {
	err := request.ParseForm()
	if err != nil {
		//发送错误信息
		result := NewApiResult(CodeFail, "ParseForm error")
		writer.Write(result.getBytes())
		return
	}
	name := request.Form.Get("name")
	password := request.Form.Get("password")
	if len(name) < 1 {
		result := NewApiResult(CodeFail, "name must not be empty!")
		writer.Write(result.getBytes())
		return
	}
	if len(password) < 1 {
		result := NewApiResult(CodeFail, "password must not be empty!")
		writer.Write(result.getBytes())
		return
	}
	dto := NewUserLoginDto(name, password)
	result := dto.Login()
	writer.Write(result.getBytes())
}

func registerHandlerFunc(writer http.ResponseWriter, request *http.Request) {
	err := request.ParseForm()
	if err != nil {
		//发送错误信息
		result := NewApiResult(CodeFail, "ParseForm error")
		writer.Write(result.getBytes())
		return
	}
	name := request.Form.Get("name")
	password := request.Form.Get("password")
	age := request.Form.Get("age")
	rePassword := request.Form.Get("rePassword")
	if len(name) < 1 {
		result := NewApiResult(CodeFail, "name must not be empty!")
		writer.Write(result.getBytes())
		return
	}
	if len(password) < 1 {
		result := NewApiResult(CodeFail, "password must not be empty!")
		writer.Write(result.getBytes())
		return
	}
	if len(age) < 1 {
		result := NewApiResult(CodeFail, "age must not be empty!")
		writer.Write(result.getBytes())
		return
	}
	if len(rePassword) < 1 {
		result := NewApiResult(CodeFail, "rePassword must not be empty!")
		writer.Write(result.getBytes())
		return
	}
	if password != rePassword {
		result := NewApiResult(CodeFail, "The passwords entered twice are inconsistent \n!")
		writer.Write(result.getBytes())
		return
	}
	ageInt, err := strconv.Atoi(age)
	if err != nil {
		result := NewApiResult(CodeFail, "The age must be a number!")
		writer.Write(result.getBytes())
		return
	}
	dto := NewUserRegisterDto(rePassword, ageInt, NewUserLoginDto(name, password))
	result := dto.register()
	writer.Write(result.getBytes())
}
func init() {
	log.SetFlags(log.Llongfile | log.Lmicroseconds | log.Ldate)
	log.SetPrefix("[...log...]")
}
func initDB() (err error) {
	db, err = sql.Open("mysql", DbUrl)
	if err != nil {
		log.Fatalln("db open error:", err)
		return err
	}
	err = db.Ping()
	if err != nil {
		log.Fatalln("db ping error:", err)
		return err
	}
	return nil
}
func main() {
	initDB()
	defer db.Close()
	http.HandleFunc("/user/login", loginHandlerFunc)
	http.HandleFunc("/user/register", registerHandlerFunc)
	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		fmt.Println("server listen error:", err)
	}
}
