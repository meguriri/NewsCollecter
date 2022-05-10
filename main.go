package main

import (
	"news/config"
	"news/dao"
	"news/log"
	"news/logic"
	"news/router"
)

func main() {
	config.InitConfig()
	log.InitLog()
	dao.InitDB()
	logic.AutoNewsHandler()
	r := router.InitRouter()
	r.Run(":" + config.Port)
}
