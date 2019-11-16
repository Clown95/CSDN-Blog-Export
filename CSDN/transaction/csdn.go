package transaction

import (
	"CSDN/models"
	"encoding/json"
	"github.com/PuerkitoBio/goquery"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"
)

func GetHtml(url string) *http.Response {
	client := &http.Client{ //要管理HTTP客户端的头域、重定向策略和其他设置，创建一个Client
		Timeout: time.Second * 2,
	}
	req, err := http.NewRequest("GET", url, nil) //NewRequest使用指定的方法、网址和可选的主题创建并返回一个新的*Request。

	if err != nil {
		log.Println(err)
	}
	req.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/78.0.3904.97 Safari/537.36") //模拟浏览器User-Agent
	req.Header.Add("Cookie", models.Cookie)
	resp, err := client.Do(req) //Do方法发送请求，返回HTTP回复
	if err != nil {
		log.Println(err)
	}
	return resp //返回网页响应
}

func GetdetailID(resp *http.Response) {
	defer resp.Body.Close()

	dom, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	dom.Find("h4").Each(func(i int, selection *goquery.Selection) {
		time.Sleep(1 * time.Second) //防止访问次数过于频繁
		detailurl, _ := selection.Find("a").Attr("href")
		index := strings.LastIndex(detailurl, "/")
		models.ArrDetailID = append(models.ArrDetailID, detailurl[index+1:])
	})
}

func ParseArticleJson(jsonurl string) (string, string) {
	resp := GetHtml(jsonurl)
	defer resp.Body.Close()
	resp_byte, _ := ioutil.ReadAll(resp.Body)
	respHtml := string(resp_byte)

	var article models.Article
	json.Unmarshal([]byte(respHtml), &article)
	title := article.Data.Title
	content := article.Data.Content
	markdown := article.Data.Markdowncontent
	if markdown == "" {
		return title + ".html", content
	} else {

		return title + ".md", markdown
	}

}

