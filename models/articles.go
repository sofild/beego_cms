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
func FindArticle(id int64) Articles {
	o := orm.NewOrm()
	o.Using("default")

	var articles Articles
	o.QueryTable("articles").Filter("id", id).One(&articles)
	return articles
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
	num, err := o.Raw(sql).Values(&articles)
	fmt.Printf("print:%d, %s", num, err)
	return articles
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
