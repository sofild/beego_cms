package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"     //引入beego的orm
	_ "github.com/go-sql-driver/mysql" //引入beego的mysql驱动
	"strconv"
	"time"
)

type Articles struct {
	Id          int64
	Title       string
	TitleCrc32  uint32
	Description string
	CateId      int32
	Pic         string
	Addtime     int64
	Modtime     int64
}

type ArticlesContent struct {
	ArticleId int64 `orm:"pk"`
	Content   string
	Author    string
	Source    string
}

/*
	新增文章
	@param title string
	@param description string
	@param cateId int32
	@param pic string
*/
func AddArticles(title string, description string, cateId int32, pic string) (int64, error) {
	o := orm.NewOrm()
	//o.Using("default")
	articles := new(Articles)
	articles.Title = title
	articles.TitleCrc32 = Crc32(title)
	articles.Description = description
	articles.CateId = cateId
	articles.Pic = pic
	articles.Addtime = time.Now().Unix()
	articles.Modtime = time.Now().Unix()
	id, err := o.Insert(articles)
	return id, err
}

/*
	修改文章
*/
func EditArticle(article Articles) bool {
	o := orm.NewOrm()
	//o.Using("default")
	_, err := o.Update(&article)
	if err != nil {
		return false
	}
	return true
}

/*
	获取单篇文章
*/
func FindArticle(id int64) map[string]string {
	o := orm.NewOrm()
	o.Using("default")

	var articles []orm.Params
	o.QueryTable("articles").Filter("id", id).Values(&articles)
	article := articles[0]

	pic := GetInterfaceValue(article["Pic"])
	times := GetInterfaceValue(article["Modtime"])
	aid := GetInterfaceValue(article["Id"])
	title := GetInterfaceValue(article["Title"])
	cate_id := GetInterfaceValue(article["CateId"])
	description := GetInterfaceValue(article["Description"])

	timestamp, _ := strconv.ParseInt(times, 10, 64)
	tm := time.Unix(timestamp, 0)
	data := make(map[string]string)

	data["id"] = aid
	data["title"] = title
	data["description"] = description
	data["pic"] = pic
	data["cate_id"] = cate_id
	data["modtime"] = tm.Format("2006-01-02 15:04:05")

	fmt.Println(data)

	return data
}

/*
	获取文章列表
*/
func ArticleList(where map[string]string, order string, offset int, limit int) []orm.Params {
	o := orm.NewOrm()
	var articles []orm.Params
	var sql string = "select id,title,cate_id,pic,modtime from articles"
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
	sql += " limit " + strconv.Itoa(offset) + "," + strconv.Itoa(limit)
	fmt.Println(sql)
	//num, err := o.QueryTable("articles").ValuesList(&articles, "id", "title", "__modtime")
	o.Raw(sql).Values(&articles)
	return articles
}

/*
	删除文章
*/
func DelArticle(id int64) bool {
	o := orm.NewOrm()
	if num, err := o.Delete(&Articles{Id: id}); err == nil {
		if num > 0 {
			return true
		}
		return false
	}
	return false
}

/*
	新增文章内容
*/
func AddContent(articleId int64, content string, author string, source string) (int64, error) {
	o := orm.NewOrm()
	articles_content := new(ArticlesContent)
	articles_content.ArticleId = articleId
	articles_content.Content = content
	articles_content.Author = author
	articles_content.Source = source
	id, err := o.Insert(articles_content)
	return id, err
}

/*
	获取单篇文章内容
*/
func FindContent(id int64) map[string]string {
	o := orm.NewOrm()
	o.Using("default")

	var content []orm.Params
	o.QueryTable("articles_content").Filter("article_id", id).Values(&content)
	cont := content[0]

	conts := GetInterfaceValue(cont["Content"])
	author := GetInterfaceValue(cont["Author"])
	source := GetInterfaceValue(cont["Source"])

	data := make(map[string]string)

	data["content"] = conts
	data["author"] = author
	data["source"] = source
	return data
}

/*
	类型断言，获取interface的值
*/
func GetInterfaceValue(value interface{}) string {
	var ret string
	switch v := value.(type) {
	case int:
		var val int
		val = v
		ret = strconv.Itoa(val)
	case string:
		ret = v
	case int64:
		var val int64
		val = v
		ret = strconv.FormatInt(val, 10)
	case int32:
		var val int32
		val = v
		ret = strconv.FormatInt(int64(val), 10)
	default:
		ret = fmt.Sprintf("%s", value)
	}
	return ret
}

/*
	删除文章内容
*/
func DelContent(id int64) bool {
	o := orm.NewOrm()
	if num, err := o.Delete(&ArticlesContent{ArticleId: id}); err == nil {
		if num > 0 {
			return true
		}
		return false
	}
	return false
}
