package models

import (
	"CSDN/utils"
	"github.com/Unknwon/goconfig"
)

var ArrDetailID []string //保存文章ID

var BlogUrl string // 配置博客地址
var Cookie string  // 配置Cookie
var TotalPage int  // 博客文章列表总页数

type Article struct { //用来解析json
	Data struct {

		Title string `json:"title"`
		Content         string `json:"content"`
		Markdowncontent string `json:"markdowncontent"`

	} `json:"data"`
}



func init() {

	runpath := utils.GetRunPath()
	cfg, err := goconfig.LoadConfigFile(runpath + "/conf/conf.ini")
	if err != nil {
		panic("没有加载到配置文件")
	}

	BlogUrl, err = cfg.GetValue("csdn", "blogurl")
	if err != nil {
		panic("blogurl错误")
	}

	Cookie, err = cfg.GetValue("csdn", "cookie")
	if err != nil {
		panic("cookie错误")
	}
	TotalPage, err = cfg.Int("csdn", "totalpage")
	if err != nil {
		panic("totalpage错误")
	}
}