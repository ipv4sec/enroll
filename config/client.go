package config

import (
	"enroll/logger"
	"fmt"
	"gopkg.in/yaml.v2"
	"os"
)

var Conf *Config

func Init(bytes []byte) {
	err := yaml.Unmarshal(bytes, &Conf)
	if err != nil {
		logger.Error("Parse Config.Yaml Fail:", err.Error())
		os.Exit(0)
	}
	logger.Info("Loaded Conf:", *Conf)
}

func Datasource() string {
	return fmt.Sprintf(
		"%v:%v@tcp(%v:%v)/%v?charset=utf8&parseTime=True&loc=Local",
		Conf.Mysql.User,
		Conf.Mysql.Pass,
		Conf.Mysql.Host,
		Conf.Mysql.Port,
		Conf.Mysql.Db)
}
