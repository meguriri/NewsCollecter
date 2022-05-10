package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/meguriri/NewsCollecter/dao"
	"github.com/meguriri/NewsCollecter/log"
	"github.com/meguriri/NewsCollecter/logic"
)

func GetIndex() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	}
}

func GetNews() gin.HandlerFunc {
	return func(c *gin.Context) {
		NewsList, _ := dao.GetAllNews()
		news, err := json.Marshal(NewsList)
		if err != nil {
			fmt.Println("json err: ", err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{
				"msg":  "获取新闻失败",
				"code": 500,
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"msg":     "获取新闻成功",
			"code":    200,
			"content": string(news),
		})
	}
}

func SearchNews() gin.HandlerFunc {
	return func(c *gin.Context) {
		title := c.Query("title")
		fmt.Println(title)
		newslist, err := dao.GetSearchNews(title)
		if err != nil {
			fmt.Println("get query list err: ", err.Error())
		}
		fmt.Println("query list", newslist)
		content, err := json.Marshal(newslist)
		if err != nil {
			fmt.Println("json err: ", err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{
				"msg":  "获取新闻失败",
				"code": 500,
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"msg":     "获取查询成功",
			"code":    200,
			"content": string(content),
		})
	}
}

func GetCategoryNews() gin.HandlerFunc {
	return func(c *gin.Context) {
		category := c.Query("category")
		newslist, err := dao.GetCategoryNews(category)
		if err != nil {
			fmt.Println("get query list err: ", err.Error())
		}
		content, err := json.Marshal(newslist)
		if err != nil {
			fmt.Println("json err: ", err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{
				"msg":  "获取类别失败",
				"code": 500,
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"msg":     "获取类别成功",
			"code":    200,
			"content": string(content),
		})

	}
}

func GetAllNews() gin.HandlerFunc {
	return func(c *gin.Context) {
		newslist, err := dao.GetAllNews()
		if err != nil {
			fmt.Println("get query list err: ", err.Error())
		}
		fmt.Println("query list", newslist)
		content, err := json.Marshal(newslist)
		if err != nil {
			fmt.Println("json err: ", err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{
				"msg":  "获取新闻失败",
				"code": 500,
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"msg":     "获取查询成功",
			"code":    200,
			"content": string(content),
		})
	}
}

func GetConsoleIndex() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.HTML(http.StatusOK, "console.html", nil)
	}
}

func DeleteNews() gin.HandlerFunc {
	return func(c *gin.Context) {
		title := c.Param("title")
		ok, err := dao.DeleteNews(title)
		if ok {
			c.JSON(http.StatusOK, gin.H{
				"msg":  "删除新闻成功",
				"code": 200,
			})
		} else {
			fmt.Println("delete err: ", err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{
				"msg":  "删除新闻失败",
				"code": 500,
			})
		}
	}
}

func AlterNews() gin.HandlerFunc {
	return func(c *gin.Context) {
		oldtitle := c.Param("title")
		newtitle := c.PostForm("newtitle")
		ok, err := dao.AlterNews(oldtitle, newtitle)
		if ok {
			c.JSON(http.StatusOK, gin.H{
				"msg":  "更新新闻成功",
				"code": 200,
			})
		} else {
			fmt.Println("Alter err: ", err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{
				"msg":  "更新新闻失败",
				"code": 500,
			})
		}

	}
}

func UpdateNews() gin.HandlerFunc {
	return func(c *gin.Context) {
		logic.GetNewsList()
		logic.DeleteOverdueNews()
		c.JSON(http.StatusOK, gin.H{
			"msg":  "更新新闻成功",
			"code": 200,
		})
	}
}

func GetLogInfo() gin.HandlerFunc {
	return func(c *gin.Context) {
		loginfo := log.GetLog()
		fmt.Println(loginfo)
		//content, err := json.Marshal(loginfo)
		// if err != nil {
		// 	fmt.Println("json err: ", err.Error())
		// 	c.JSON(http.StatusInternalServerError, gin.H{
		// 		"msg":  "获取日志失败",
		// 		"code": 500,
		// 	})
		// 	return
		// }
		c.JSON(http.StatusOK, gin.H{
			"msg":  "获取日志成功",
			"log":  loginfo,
			"code": 200,
		})
	}
}
