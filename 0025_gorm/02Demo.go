package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type Student struct {
	gorm.Model
	Name string
	Age  int64
}

func main() {
	dsn := "root:root@tcp(127.0.0.1:3306)/data_shop?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN: dsn,
	}), &gorm.Config{
		SkipDefaultTransaction: false, //跳过默认事务
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // 复数形式 User的表名应该是students
			TablePrefix:   "t_", //表名前缀 User的表名应该是t_students
		},
		DisableForeignKeyConstraintWhenMigrating: true, //设置成为逻辑外键(在物理数据库上没有外键，仅体现在代码上)
	})
	if err != nil {
		fmt.Println("mysql open error:", err)
	}
	db.AutoMigrate(&Student{})
	////数据插入
	//student1 := Student{
	//	Name: "tom",
	//	Age:  10,
	//}
	//student2 := Student{
	//	Name: "jemmy",
	//	Age:  10,
	//}
	//db.Create(&student1)
	//db.Create(&student2)
	var student []Student
	db.Debug().First(&student)
	fmt.Println("根据主键查询第一条记录：", student)
	db.Debug().Take(&student)
	fmt.Println("随机获取一条记录：", student)
	db.Debug().Last(&student)
	fmt.Println("获取最后一条记录：", student)
	db.Find(&student)
	fmt.Println("查找所有记录：", student)
	db.Find(&student, 2)
	fmt.Println("查找某条记录：", student)
	//where使用
	db.Debug().Where("name=?", "tom").First(&student)
	fmt.Println("查询第一条匹配条件记录：", &student)
	db.Debug().Where("name = ?", "tom").Find(&student)
	fmt.Println("查询所有匹配条件记录：", &student)
	db.Debug().Where("name <> ?", "tom").Find(&student)
	fmt.Println("查询不等于条件的记录：", student)
	db.Debug().Where("name in ?", []string{"tom", "jimmy"}).Find(&student)
	fmt.Println("in的使用：", student)
	db.Debug().Where("name like ?", "%mm%").Find(&student)
	fmt.Println("模糊查询", student)
	db.Debug().Where("name = ? and id = ?", "tom", "1").Find(&student)
	fmt.Println("多条件查询", student)
}
