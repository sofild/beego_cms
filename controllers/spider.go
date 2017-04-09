package controllers

import (
	"beego_cms/models"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	//"os"
	"strconv"
	"time"
)

type SpiderController struct {
	beego.Controller
}

type Spider struct {
	Name    string `form:"name"`
	Url     string `form:"url"`
	List    string `form:"list"`
	Link    string `form:"link"`
	Title   string `form:"title"`
	Author  string `form:"author"`
	Content string `form:"content"`
	CateId  int32  `form:"cate_id"`
	Pic     string `form:"pic"`
}

func (c *SpiderController) Index() {
	c.TplName = "spider_list.tpl"
}

func (c *SpiderController) List() {
	var page int = 1
	var perPage int = 200
	var where map[string]string
	spiders := models.List(where, "id desc", perPage*(page-1), perPage)

	//获取所有分类ID
	var cate_ids []string
	for _, v := range spiders {
		cate_id, _ := v["cate_id"].(string)
		cate_ids = append(cate_ids, cate_id)
	}
	cates := models.GetCateNames(cate_ids)
	fmt.Println(cates, cate_ids)

	var datas []map[string]string
	for _, v := range spiders {

		id, _ := v["id"].(string)
		name, _ := v["name"].(string)
		url, _ := v["url"].(string)
		list, _ := v["list"].(string)
		link, _ := v["link"].(string)
		title, _ := v["title"].(string)
		author, _ := v["author"].(string)
		content, _ := v["content"].(string)
		addtime, _ := v["addtime"].(string)
		cate_id, _ := v["cate_id"].(string)
		pic, _ := v["pic"].(string)

		timestamp, _ := strconv.ParseInt(addtime, 10, 64)
		tm := time.Unix(timestamp, 0)
		spider := make(map[string]string)

		spider["id"] = id
		spider["name"] = name
		spider["url"] = url
		spider["list"] = list
		spider["link"] = link
		spider["title"] = title
		spider["author"] = author
		spider["content"] = content
		spider["cate_id"] = cate_id
		spider["cate_name"] = cates[cate_id]
		spider["pic"] = pic
		spider["time"] = tm.Format("2006-01-02 15:04:05")
		spider["op"] = fmt.Sprintf("<a href='/cate/add/%s'>修改</a>&nbsp;&nbsp;<a href='/cate/del/%s'>删除</a>", id, id)

		datas = append(datas, spider)
	}
	data, err := json.Marshal(datas)
	if err != nil {
		beego.Info(err)
	}
	c.Ctx.WriteString(string(data))
}

func (c *SpiderController) Add() {
	var where map[string]string
	cates := models.CateList(where, "id asc", 0, 0)
	c.Data["Cates"] = cates
	c.TplName = "spider_add.tpl"
}

func (c *SpiderController) DoAdd() {
	spider := Spider{}
	if err := c.ParseForm(&spider); err != nil {
		beego.Info(err)
	} else {
		var id int64
		id, err = models.Add(spider.Name, spider.Url, spider.List, spider.Link, spider.Title, spider.Author, spider.Content, spider.CateId, spider.Pic)
		if err != nil {
			beego.Info(err)
			c.Redirect("/spider/add", 302)
		}
		fmt.Println(id)
		c.Redirect("/spider", 302)
	}
}
