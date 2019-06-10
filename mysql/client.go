package mysql

import (
	"enroll/config"
	"enroll/logger"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"os"
)

var Clinet *gorm.DB

func Init() {
	var err error
	Clinet, err = gorm.Open("mysql", config.Datasource())
	if err != nil {
		logger.Error("Connect MySQL Fail:", err.Error())
		os.Exit(0)
	}
	logger.Info("Connected To MySQL:", config.Conf.Mysql.Host, config.Conf.Mysql.Port)
}
