package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"     //引入beego的orm
	_ "github.com/go-sql-driver/mysql" //引入beego的mysql驱动
	"strconv"
	"time"
)

type ArticlesCate struct {
	Id       int64
	ParentId int64
	Name     string
	Sort     int32
	Addtime  int64
}

/*
	新增分类
	@param title string
	@param description string
	@param cateId int32
	@param pic string
*/
func AddCate(parentId int64, name string, sort int32) (int64, error) {
	o := orm.NewOrm()
	articles_cate := new(ArticlesCate)
	articles_cate.ParentId = parentId
	articles_cate.Name = name
	articles_cate.Sort = sort
	articles_cate.Addtime = time.Now().Unix()
	id, err := o.Insert(articles_cate)
	return id, err
}

/*
	修改分类
*/
func EditCate(article_cate ArticlesCate) bool {
	o := orm.NewOrm()
	_, err := o.Update(&article_cate)
	if err != nil {
		return false
	}
	return true
}

/*
	获取分类信息
*/
func FindCate(id int64) ArticlesCate {
	o := orm.NewOrm()

	var articles_cate ArticlesCate
	o.QueryTable("articles_cate").Filter("id", id).One(&articles_cate)
	return articles_cate
}

/*
	获取分类列表
*/
func CateList(where map[string]string, order string, offset int, limit int) []orm.Params {
	o := orm.NewOrm()
	var articles_cate []orm.Params
	var sql string = "select * from articles_cate"
	var cond string = ""
	var i int = 0
	for k, v := range where {
		cond += (k + "=" + v)
		if i > 0 {
			cond += " and "
		}
		i++
	}
	sql += " " + cond
	sql += " order by " + order
	if limit > 0 {
		sql += " limit " + strconv.Itoa(offset) + "," + strconv.Itoa(limit)
	}
	fmt.Println(sql)
	o.Raw(sql).Values(&articles_cate)
	return articles_cate
}
