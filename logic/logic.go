package logic

import (
	"fmt"
	"log"
	"news/dao"
	"time"
)

func InitCategory() {
	dao.Category["国内"] = "http://news.baidu.com/widget?id=civilnews&t=1651559936161"
	dao.Category["国际"] = "http://news.baidu.com/widget?id=InternationalNews&t=1651559936167"
	dao.Category["娱乐"] = "http://news.baidu.com/widget?id=EnterNews&t=1651559936176"
	dao.Category["体育"] = "http://news.baidu.com/widget?id=SportNews&t=1651559936186"
	dao.Category["财经"] = "http://news.baidu.com/widget?id=FinanceNews&t=1651559936190"
	dao.Category["科技"] = "http://news.baidu.com/widget?id=TechNews&t=1651559936201"
	dao.Category["军事"] = "http://news.baidu.com/widget?id=MilitaryNews&t=1651559936209"
	dao.Category["互联网"] = "http://news.baidu.com/widget?id=InternetNews&t=1651559936216"
	dao.Category["探索"] = "http://news.baidu.com/widget?id=DiscoveryNews&t=1651559936222"
	dao.Category["女人"] = "http://news.baidu.com/widget?id=LadyNews&t=1651559936230"
	dao.Category["健康"] = "http://news.baidu.com/widget?id=HealthNews&t=1651559936236"
}

func GetNewsList() {
	cnt := 0
	hthtml := GetHotHTML("http://news.baidu.com")
	cnt += FindNews("热点新闻", hthtml)
	for i, v := range dao.Category {
		html := GetHTML(v)
		cnt += FindNews(i, html)
	}
	fmt.Println("更新结束,共更新", cnt, "个新闻")
	log.Println("更新结束,共更新", cnt, "个新闻")
}

func AutoNewsHandler() {
	InitCategory()
	//fmt.Println("开始更新。。", time.Now().Format("2006-01-02 15:04:05"))
	log.Println("开始更新。。", time.Now().Format("2006-01-02 15:04:05"))
	GetNewsList()
	UpTimer := time.NewTimer(time.Hour)
	DelTimer := time.NewTimer(time.Hour * 24)
	go func() {
		for {
			select {
			case <-UpTimer.C:
				//fmt.Println("开始更新。。", time.Now().Format("2006-01-02 15:04:05"))
				log.Println("开始更新。。", time.Now().Format("2006-01-02 15:04:05"))
				GetNewsList()
				UpTimer.Reset(time.Hour)
			case <-DelTimer.C:
				//fmt.Println("开始删除过期新闻。。", time.Now().Format("2006-01-02 15:04:05"))
				log.Println("开始删除过期新闻。。", time.Now().Format("2006-01-02 15:04:05"))
				DeleteOverdueNews()
				DelTimer.Reset(time.Hour * 24)
			}
		}
	}()
}

func DeleteOverdueNews() {
	now := time.Now()
	d, _ := time.ParseDuration("-24h")
	d1 := now.Add(2 * d)
	rows, err := dao.DeleteOverdueNews(d1)
	if err != nil {
		//fmt.Println("删除失败",err)
		log.Println("删除失败", err)
	} else {
		//fmt.Println("删除成功,删除了",rows,"条新闻")
		log.Println("删除成功,删除了", rows, "条新闻")
	}
}
