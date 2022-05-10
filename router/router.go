package router

import (
	h "news/handler"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")
	r.Static("/static", "./static/")
	r.GET("/", h.GetIndex())
	listGroup := r.Group("/news")
	{
		listGroup.GET("/", h.GetNews())
		listGroup.GET("/category", h.GetCategoryNews())
		listGroup.GET("/search", h.SearchNews())
	}
	console := r.Group("/console")
	{
		console.GET("/", h.GetConsoleIndex())
		console.GET("/update", h.UpdateNews())
		console.GET("/all", h.GetAllNews())
		console.GET("/search", h.SearchNews())
		console.GET("/log", h.GetLogInfo())
		console.DELETE("/del/:title", h.DeleteNews())
		console.POST("/alt/:title", h.AlterNews())
	}
	return r
}
