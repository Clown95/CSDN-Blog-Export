package main

import (
	"CSDN/models"
	"CSDN/transaction"
	"CSDN/utils"
	"fmt"
	"time"
)

func main() {

	fmt.Println("设置成功，开始导出blog,时间较长请等待！！")
	for i := 1; i <= models.TotalPage; i++ {
		time.Sleep(800) //设置延时
		url := fmt.Sprintf("%s/article/list/%d", models.BlogUrl, i)
		resq := transaction.GetHtml(url)
		transaction.GetdetailID(resq)
	}
	runpath := utils.GetRunPath()

	for i := 0; i < len(models.ArrDetailID); i++ {
		jsonurl := fmt.Sprintf("https://mp.csdn.net/mdeditor/getArticle?id=%s", models.ArrDetailID[i])
		name, content := transaction.ParseArticleJson(jsonurl)
		utils.WriteWithIoutil(runpath+"/"+name, content)
		time.Sleep(1000) //设置延时
	}
}
