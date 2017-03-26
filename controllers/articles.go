package controllers

import (
	"beego_cms/models"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"os"
	"strconv"
	"time"
)

type ArticlesController struct {
	beego.Controller
}

type Articles struct {
	Title       string `form:"title"`
	Description string `form:"description"`
	CateId      int32  `form:"cate_id"`
	Pic         string `form:"pic"`
	Author      string `form:"author"`
	Source      string `form:"source"`
	Content     string `form:"content"`
}

func (c *ArticlesController) Index() {
	c.TplName = "articles_list.tpl"
}

func (c *ArticlesController) List() {
	var page int = 1
	var perPage int = 200
	var where map[string]string
	articles := models.ArticleList(where, "modtime desc", perPage*(page-1), perPage)
	var datas []map[string]string
	//fmt.Println(articles)
	//fmt.Println(len(articles))

	//获取所有分类ID
	var cate_ids []string
	for _, v := range articles {
		cate_id, _ := v["cate_id"].(string)
		cate_ids = append(cate_ids, cate_id)
	}
	cates := models.GetCateNames(cate_ids)
	fmt.Println(cates, cate_ids)

	//格式化数据
	for _, v := range articles {
		//fmt.Println(v["pic"], v["modtime"])
		imgUrl, _ := v["pic"].(string)
		times, _ := v["modtime"].(string)
		id, _ := v["id"].(string)
		title, _ := v["title"].(string)
		cate_id, _ := v["cate_id"].(string)

		timestamp, _ := strconv.ParseInt(times, 10, 64)
		tm := time.Unix(timestamp, 0)
		article := make(map[string]string)
		//fmt.Println(v["modtime"])

		article["id"] = id
		article["title"] = title
		article["cate_name"] = cates[cate_id]
		article["imgUrl"] = "<img src=" + imgUrl + " width='50px' />"
		article["time"] = tm.Format("2006-01-02 15:04:05")
		article["op"] = fmt.Sprintf("<a href='/articles/add/%s'>修改</a>&nbsp;&nbsp;<a href='/articles/del/%s'>删除</a>", id, id)

		datas = append(datas, article)
	}
	//fmt.Println(datas)
	data, err := json.Marshal(datas)
	if err != nil {
		beego.Info(err)
	}
	c.Ctx.WriteString(string(data))
}

func (c *ArticlesController) Add() {
	id := c.Ctx.Input.Param(":id")
	aid, _ := strconv.ParseInt(id, 10, 64)
	if aid > 0 {
		article := models.FindArticle(aid)
		c.Data["Article"] = article

		content := models.FindContent(aid)
		c.Data["Content"] = content
	}
	var where map[string]string
	cates := models.CateList(where, "id asc", 0, 0)
	c.Data["Cates"] = cates
	c.TplName = "articles_add.tpl"
}

func (c *ArticlesController) DoAdd() {
	articles := Articles{}
	f, h, err := c.GetFile("pic")
	defer f.Close()
	if err != nil {
		beego.Info(err)
		c.Redirect("/articles/add", 302)
	} else {
		path, absPath := getUplodDir()
		var pic string = path + h.Filename
		var picPath string = absPath + h.Filename

		fmt.Println(h.Filename)
		fmt.Println(picPath)

		c.SaveToFile("pic", picPath)
		articles.Pic = pic
	}

	if err := c.ParseForm(&articles); err != nil {
		beego.Info(err)
	} else {
		var articleId int64
		articleId, err = models.AddArticles(articles.Title, articles.Description, articles.CateId, articles.Pic)
		if err != nil {
			beego.Info(err)
			c.Redirect("/articles/add", 302)
		}
		_, err = models.AddContent(articleId, articles.Content, articles.Author, articles.Source)
		if err != nil {
			beego.Info(err)
			c.Redirect("/articles/add", 302)
		}
		c.Redirect("/articles", 302)
	}
}

func (c *ArticlesController) Del() {
	id := c.Ctx.Input.Param(":id")
	c.Ctx.WriteString(id)
}

/*
	判断文件目录是否存在
	@param path string
	@return bool
*/
func pathIsExists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return false
}

/*
	创建目录
	@param path string 需创建的目录路径，前面加/
	@return bool
*/
func createDir(path string) bool {
	curDir, _ := os.Getwd()
	err := os.Mkdir(curDir+path, os.ModePerm)
	if err != nil {
		return false
	}
	return true
}

/*
	获取上传目录，没有则新建
	@param
*/
func getUplodDir() (string, string) {
	curDir, _ := os.Getwd()
	var nowDir string = time.Now().Format("2006-01-02")
	var nowPath string = "/static/upload/" + nowDir
	var path string = nowPath + "/"
	//绝对路径
	var absPath string = curDir + path
	if pathIsExists(nowPath) {
		return path, absPath
	}
	createDir(nowPath)
	return path, absPath
}
