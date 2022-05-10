package main

import (
	"github.com/meguriri/NewsCollecter/config"
	"github.com/meguriri/NewsCollecter/dao"
	"github.com/meguriri/NewsCollecter/log"
	"github.com/meguriri/NewsCollecter/logic"
	"github.com/meguriri/NewsCollecter/router"
)

func main() {
	config.InitConfig()
	log.InitLog()
	dao.InitDB()
	logic.AutoNewsHandler()
	r := router.InitRouter()
	r.Run(":" + config.Port)
}
