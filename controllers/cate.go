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

type CateController struct {
	beego.Controller
}

type Cate struct {
	ParentId int64  `form:"parent_id"`
	Name     string `form:"name"`
	Sort     int32  `form:"sort"`
}

func (c *CateController) Prepare() {
	c.Data["MenuName"] = "cate"
}

func (c *CateController) Index() {
	c.TplName = "articles_cate_list.tpl"
}

func (c *CateController) List() {
	var page int = 1
	var perPage int = 200
	var where map[string]string
	cates := models.CateList(where, "id asc", perPage*(page-1), perPage)
	var datas []map[string]string
	for _, v := range cates {

		id, _ := v["id"].(string)
		name, _ := v["name"].(string)
		parent_id, _ := v["parent_id"].(string)
		sort, _ := v["sort"].(string)
		addtime, _ := v["addtime"].(string)

		timestamp, _ := strconv.ParseInt(addtime, 10, 64)
		tm := time.Unix(timestamp, 0)
		cate := make(map[string]string)

		cate["id"] = id
		cate["parent_id"] = parent_id
		cate["name"] = name
		cate["sort"] = sort
		cate["time"] = tm.Format("2006-01-02 15:04:05")
		cate["op"] = fmt.Sprintf("<a href='/cate/add/%s'>修改</a>&nbsp;&nbsp;<a href='/cate/del/%s'>删除</a>", id, id)

		datas = append(datas, cate)
	}
	data, err := json.Marshal(datas)
	if err != nil {
		beego.Info(err)
	}
	c.Ctx.WriteString(string(data))
}

func (c *CateController) Add() {
	c.TplName = "articles_cate_add.tpl"
}

func (c *CateController) DoAdd() {
	cate := Cate{}
	if err := c.ParseForm(&cate); err != nil {
		beego.Info(err)
	} else {
		var cateId int64
		cateId, err = models.AddCate(cate.ParentId, cate.Name, cate.Sort)
		if err != nil {
			beego.Info(err)
			c.Redirect("/cate/add", 302)
		}
		fmt.Println(cateId)
		c.Redirect("/cate", 302)
	}
}
