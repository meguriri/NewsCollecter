package config

import (
	"fmt"

	"github.com/meguriri/NewsCollecter/dao"
	l "github.com/meguriri/NewsCollecter/log"
	"github.com/spf13/viper"
)

var (
	Port string
)

func InitConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("viper err: ", err)
	}
	dao.Username = viper.GetString("mysql.username")
	dao.Password = viper.GetString("mysql.password")
	dao.Host = viper.GetString("mysql.host")
	dao.Port = viper.GetString("mysql.port")
	dao.Database = viper.GetString("mysql.database")
	Port = viper.GetString("router.port")
	l.File = viper.GetString("log.file")
}
