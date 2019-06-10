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
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
)

var conf = flag.String("config", "config.yaml", "Used Cofnig.Yaml")

func main() {
	flag.Parse()

	bytes, err := ioutil.ReadFile(filepath.Join(filepath.Dir(os.Args[0]), *conf))
	if err != nil {
		logger.Error("Read Config.Yaml Fail:", err.Error())
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

	router.GET("/example", func(c *gin.Context) {
		header := c.Writer.Header()
		header["Content-type"] = []string{"application/octet-stream"}
		header["Content-Disposition"] = []string{"attachment; filename= example.csv"}

		file, err := os.Open(config.Conf.Csv.Example)
		if err != nil {
			logger.Error("读取CSV样例文件失败:", err.Error())
			c.Status(404)
			return
		}
		defer file.Close()
		io.Copy(c.Writer, file)
	})
	router.Static("/dist", "./node_modules")
	router.Use(static.Serve("/", static.LocalFile("./static", true)))

	v1 := router.Group("/v1")
	{
		v1.Use(middleware.TokenChecker([]string{"/v1/token", "/v1/file"}))
		v1.Use(middleware.AdminChecker([]string{"/v1/site"}))

		v1.POST("/file", csv.Upload)
		v1.GET("/file", csv.Parse)

		v1.POST("/user", user.Import) // 导入CSV
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

	}

	go router.Run(config.Conf.Application.Address)
	logger.Info("Listened Addr:", config.Conf.Application.Address)
	select {}
}