package logic

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"news/dao"
	"regexp"
	"strings"
	"time"
)

func GetHTML(url string) string {
	res, err := http.Get(url)
	if err != nil {
		log.Println(err)
	}
	html, _ := ioutil.ReadAll(res.Body)
	res.Body.Close()
	return string(html)
}

func GetHotHTML(url string) string {
	res, err := http.Get(url)
	if err != nil {
		log.Println(err)
	}
	html, _ := ioutil.ReadAll(res.Body)
	reg := regexp.MustCompile(`id="body"(([\s\S])*?)<div id="foot`)
	body := reg.FindString(string(html))
	res.Body.Close()
	return body
}

func FindNews(category, html string) int {
	reg := regexp.MustCompile(`<li(([\s\S])*?)</li>`)
	list := reg.FindAllString(html, -1)
	cnt := 0
	for _, v := range list {
		if strings.Contains(v, "#") {
			continue
		}
		if strings.Contains(v, "javascript:void(0);") {
			continue
		}
		if strings.Contains(v, "辟谣") {
			continue
		}
		if strings.Contains(v, "举报") {
			continue
		}
		reg := regexp.MustCompile(`<a(([\s\S])*?)</a>`)
		list := reg.FindAllString(v, -1)
		for _, v := range list {

			reg1 := regexp.MustCompile(`>(([\s\S])+?)<`)
			s1 := reg1.FindString(v)
			reg2 := regexp.MustCompile(`href="(([\s\S])*?)"`)
			s2 := reg2.FindString(v)
			var title string
			b1 := []byte(s1)
			if strings.Contains(s1, "<b>") {
				title = string(b1[4 : len(b1)-1])
			} else {
				title = string(b1[1 : len(b1)-1])
			}
			b2 := []byte(s2)
			url := string(b2[6 : len(b2)-1])

			news := dao.News{
				Category: category,
				Title:    title,
				Time:     time.Now(),
				Url:      url,
			}
			ok, _ := dao.InsertNews(news)
			if ok {
				fmt.Println("插入成功", "类别: "+news.Category+" 标题: "+news.Title+" url: "+news.Url)
				cnt++
			} else {
				fmt.Println("插入失败: " + news.Title + " 新闻已存在")
			}
		}
	}
	return cnt
}
