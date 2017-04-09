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
	o.Raw(sql).Values(&articles_cate)
	return articles_cate
}

/*
	通过分类ID获取分类名称
*/
func GetCateNames(cateIds []string) map[string]string {
	var ids string = ""
	for k, v := range cateIds {
		if k > 0 {
			ids += ","
		}
		ids += v
	}
	var sql string = fmt.Sprintf("select * from articles_cate where id in (%s) order by id desc", ids)
	fmt.Println(sql)
	o := orm.NewOrm()
	var articles_cate []orm.Params
	o.Raw(sql).Values(&articles_cate)
	cates := make(map[string]string)
	for _, v := range articles_cate {
		id, _ := v["id"].(string)
		name, _ := v["name"].(string)
		cates[id] = name
	}
	return cates
}

/*
	获取所有分类
*/
/*
func GetCates(parentId int) []map[string]string {
	var parent_id string = strconv.Itoa(parentId)
	where := make(map[string]string)
	where["parent_id"] = parent_id
	order := "id asc"
	cates = CateList(where, order, 0, 100)
	for v, _ := range cates {

	}
}
*/
