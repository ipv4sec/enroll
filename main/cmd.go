package main

import (
	"enroll/admin"
	"enroll/config"
	"enroll/csv"
	"enroll/logger"
	"enroll/middleware"
	"enroll/mysql"
	"enroll/redis"
	"enroll/site"
	"enroll/user"
	"flag"
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"os"
	"path/filepath"
)

var (
	conf = flag.String("config", "config.yaml", "Used Cofnig.Yaml")
	build = ""
)

func main() {
	flag.Parse()

	logger.Info("版权所有 齐凤龙")
	logger.Info("此软件未经许可禁止拷贝、修改。如有问题请联系技术支持")
	logger.Info("联系方式: <qifenglong@foxmail.com>.")
	logger.Info("联系方式: <ipv4sec@gmail.com>.")
	logger.Info("联系方式: 18510088520.")

	logger.Info("Copyright (c) 2019 Qi Fenglong Author.")
	logger.Info("All rights reserved.")
	logger.Info("FAQ: <qifenglong@foxmail.com>.")
	logger.Info("FAQ: <ipv4sec@gmail.com>.")
	logger.Info("FAQ: 18510088520.")

	logger.Info("Build: " + build)


	bytes, err := ioutil.ReadFile(filepath.Join(filepath.Dir(os.Args[0]), *conf))
	if err != nil {
		logger.Error("Read Err:", err.Error())
		os.Exit(0)
	}
	config.Init(bytes)
	mysql.Init()
	redis.Init()

	// mysql.Clinet.AutoMigrate(&csv.Csv{})

	//mysql.Clinet.AutoMigrate(&user.User{})
	//mysql.Clinet.AutoMigrate(&site.Site{})
	//mysql.Clinet.AutoMigrate(&admin.Admin{})
	//admin.AddAdminByNameAndPassAndSiteId("admin", "Admin123...", 0)

	gin.SetMode(gin.ReleaseMode)

	router := gin.New()
	router.Use(gin.Recovery())
	router.Static("/dist", "./node_modules")
	router.Static("/uploads", "./uploads")
	router.Static("/generates", "./generates")

	router.Use(static.Serve("/", static.LocalFile("./static", true)))

	v1 := router.Group("/v1")
	{
		v1.Use(middleware.TokenChecker([]string{"/v1/token", "/v1/file", "/v1/download"}))
		// v1.Use(middleware.AdminChecker([]string{"/v1/site"})) // TODO 暂时修改

		v1.POST("/file", csv.Upload)
		v1.GET("/file", csv.Parse)

		// v1.POST("/user", user.Import) // 导入CSV
		v1.POST("/user", user.ImportJson) // 导入JSON
		v1.GET("/user", user.GetBySiteId)
		v1.DELETE("/user", user.DeleteNotConfirmedUser) // 删除单个记录
		v1.PUT("/user", user.ConfirmAll) // 正式提交

		// 登录接口
		v1.POST("/token", admin.Login)

		v1.GET("/admin", admin.GetAll)
		v1.DELETE("/admin", admin.DeleteByAdminId)
		v1.POST("/admin", admin.Add)
		v1.PUT("/admin", admin.ChangePasswordByAdminId)// 修改密码

		// 修改密码
		v1.PUT("/password", admin.ChangePassword)

		// Sites
		v1.GET("/site", site.GetAll)
		v1.DELETE("/site", site.DeleteBySiteId)
		v1.POST("/site", site.Add)

		v1.GET("/search", user.GetByCardNum)
		v1.GET("/census", user.GetCensusResult)

		v1.GET("/permission", user.GetPermission)

		v1.GET("/download", user.DownloadCsvFile)

	}

	go router.Run(config.Conf.Application.Address)
	logger.Info("Listened Addr:", config.Conf.Application.Address)
	select {}
}