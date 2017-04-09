package models

import (
	//"fmt"
	"github.com/astaxie/beego/orm"     //引入beego的orm
	_ "github.com/go-sql-driver/mysql" //引入beego的mysql驱动
	"strconv"
	"time"
)

type Spider struct {
	Id      int64
	Name    string
	Url     string
	List    string
	Link    string
	Title   string
	Author  string
	Content string
	Addtime int64
	CateId  int32
	Pic     string
}

/*
	新增分类
	@param title string
	@param description string
	@param cateId int32
	@param pic string
*/
func Add(name string, url string, list string, link string, title string, author string, content string, cate_id int32, pic string) (int64, error) {
	o := orm.NewOrm()
	spider := new(Spider)
	spider.Name = name
	spider.Url = url
	spider.List = list
	spider.Title = title
	spider.Author = author
	spider.Content = content
	spider.CateId = cate_id
	spider.Pic = pic
	spider.Addtime = time.Now().Unix()
	id, err := o.Insert(spider)
	return id, err
}

/*
	修改分类
*/
func Edit(spider Spider) bool {
	o := orm.NewOrm()
	_, err := o.Update(&spider)
	if err != nil {
		return false
	}
	return true
}

/*
	获取信息
*/
func Find(id int64) Spider {
	o := orm.NewOrm()

	var spider Spider
	o.QueryTable("spider").Filter("id", id).One(&spider)
	return spider
}

/*
	获取列表
*/
func List(where map[string]string, order string, offset int, limit int) []orm.Params {
	o := orm.NewOrm()
	var spider []orm.Params
	var sql string = "select * from spider"
	var cond string = ""
	var i int = 0
	for k, v := range where {
		if i > 0 {
			cond += " and "
		}
		cond += (k + "=" + v)
		i++
	}
	sql += " " + cond
	sql += " order by " + order
	if limit > 0 {
		sql += " limit " + strconv.Itoa(offset) + "," + strconv.Itoa(limit)
	}
	//fmt.Println(sql)
	o.Raw(sql).Values(&spider)
	return spider
}
