package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

//表的设计
type User struct {
	Id int
	Name string `orm:"unique"`
	Pwd string
	//Article []*Article `orm:"rel(m2m)"`
}
//文章结构体
type Article struct {
	Id int `orm:"pk;auto"`
	ArtiName string `orm:"size(20)"`
	Atime time.Time `orm:"auto_now"`
	Acount int `orm:"default(0);null"`
	Acontent string
	Aimg string
	Atype string

	//ArticleType*ArticleType `orm:"rel(fk)"`
	//User []*User `orm:"reverse(many)"`
}

//类型表
//type ArticleType struct {
//	Id int
//	Tname string
//	Article []*Article `orm:"reverse(many)"`
//}


func init() {
	/*
		1,连接数据库
		2,注册表
		3,生成表
	*/
	orm.RegisterDataBase("default", "mysql", "root:root@tcp(127.0.0.1:3306)/class2?charset=utf8")

	orm.RegisterModel(new(User),new(Article))//,new(ArticleType)

	orm.RunSyncdb("default", false, true)
}

