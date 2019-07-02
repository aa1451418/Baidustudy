package main

/* 导入2个环境 */
import "github.com/astaxie/beego/orm"
import _ "github.com/go-sql-driver/mysql"

/* 结构体 */
type User struct {
	Id    int
	Name  string
	Age   int
	Email string
}

/* 连接数据库 */
func main() {
	orm.RegisterDataBase("default", "mysql", "root:root@tcp(127.0.0.1:3306)/class2")
	orm.RegisterModel(new(Students))
	orm.RunSyncdb("default", true, true)
}
