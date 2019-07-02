package main

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

//Students为结构体定义名称
type Students struct {
	Id          int
	Name, Email string
}

//连接数据库并在数据内创建表
func main() {
	orm.RegisterDataBase("default", "mysql", "root:root@tcp(127.0.0.1:3306)/class2")
	orm.RegisterModel(new(Students))
	orm.RunSyncdb("default", true, true)
}
