package models

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"     //引入beego的orm
	_ "github.com/go-sql-driver/mysql" //引入beego的mysql驱动
	"hash/crc32"
	"io"
)

func init() {
	//注册模型
	orm.RegisterModel(new(User))
	orm.RegisterModel(new(Articles))
	orm.RegisterModel(new(ArticlesContent))
	orm.RegisterModel(new(ArticlesCate))
	//数据库连接
	var mysql_host string = beego.AppConfig.String("mysql_host")
	var mysql_port string = beego.AppConfig.String("mysql_port")
	var mysql_user string = beego.AppConfig.String("mysql_user")
	var mysql_pass string = beego.AppConfig.String("mysql_pass")
	var mysql_db string = beego.AppConfig.String("mysql_db")
	var link string = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8", mysql_user, mysql_pass, mysql_host, mysql_port, mysql_db)
	fmt.Println(link)
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", link)
	orm.RunSyncdb("default", false, true)
}

func MD5(text string) string {
	ctx := md5.New()
	ctx.Write([]byte(text))
	return hex.EncodeToString(ctx.Sum(nil))
}

func Crc32(title string) uint32 {
	ieee := crc32.NewIEEE()
	io.WriteString(ieee, title)
	s := ieee.Sum32()
	return s
}
