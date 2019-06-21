package user

import (
	"enroll/config"
	"enroll/csv"
	"enroll/logger"
	"github.com/gin-gonic/gin"
	"io"
	"os"
	"strconv"
)

func GetBySiteId(c *gin.Context) {
	siteIdStr := c.Query("siteId")
	siteId, err := strconv.ParseInt(siteIdStr, 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
			"message": "参数错误",
		})
		return
	}
	result, err := GetUserBySiteId(siteId)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
			"message": "服务器内部错误",
		})
		return
	}
	c.JSON(200, gin.H{
		"data": result,
		"message": "获取已导入数据成功",
	})
}

func GetByCardNum(c *gin.Context) {
	cardNum := c.Query("cardNum")
	result, err := SearchByCardNum(cardNum)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
			"message": "服务器内部错误",
		})
		return
	}
	c.JSON(200, gin.H{
		"data": result,
		"message": "获取搜索数据成功",
	})
}

func Import(c *gin.Context) {
	csvFilename := c.Query("filename")
	siteIdStr := c.Query("siteId")
	siteId, err := strconv.ParseInt(siteIdStr, 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
			"message": "参数错误",
		})
		return
	}

	result, err := csv.Read(config.Conf.Csv.Uploaded + csvFilename)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
			"message": "解析失败",
		})
		return
	}
	if len(result) < 2 {
		c.JSON(400, gin.H{
			"error": "未找到数据",
			"message": "解析失败",
		})
		return
	}
	err = SaveImportedCsvDatas(result[1:], siteId)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
			"message": "导入失败",
		})
		return
	}
	c.JSON(200, gin.H{
		"message": "导入成功",
	})
}

func DeleteNotConfirmedUser(c *gin.Context) {
	userIdStr := c.Query("userId")
	userId, err := strconv.ParseInt(userIdStr, 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
			"message": "参数错误",
		})
		return
	}
	err = DeleteNotConfirmedUserByUserId(userId)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
			"message": "删除失败",
		})
		return
	}
	c.JSON(200, gin.H{
		"message": "删除成功",
	})

}

func ConfirmAll(c *gin.Context) {
	siteIdStr := c.Query("siteId")
	siteId, err := strconv.ParseInt(siteIdStr, 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
			"message": "参数错误",
		})
		return
	}
	rowsAffected := ConfirmUserBySiteId(siteId)
	c.JSON(200, gin.H{
		"message": "正式提交成功",
		"affected": rowsAffected,
	})
}

func GetCensusResult(c *gin.Context) {
	adminId, _ := c.Get("adminId")
	result, err := CensusByAdmin(adminId.(int64))
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
			"message": "服务器内部错误",
		})
		return
	}
	c.JSON(200, gin.H{
		"message": "获取统计成功",
		"data": result,
	})
}

func GetPermission(c *gin.Context) {
	adminId, _ := c.Get("adminId")
	c.JSON(200, gin.H{
		"message": "获取统计成功",
		"data": adminId.(int64) == 1,
	})
}

func DownloadCsvFile(c *gin.Context) {
	siteIdStr := c.Query("siteId")
	siteId, err := strconv.ParseInt(siteIdStr, 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
			"message": "参数错误",
		})
		return
	}
	csvFilename := GetDownloadCsvFilePathBySiteId(siteId)

	header := c.Writer.Header()
	header["Content-type"] = []string{"application/octet-stream"}
	header["Content-Disposition"] = []string{"attachment; filename= example.csv"}

	file, err := os.Open(config.Conf.Csv.Generated + csvFilename)
	if err != nil {
		logger.Error("读取CSV样例文件失败:", err.Error())
		c.Status(404)
		return
	}
	defer file.Close()
	io.Copy(c.Writer, file)
}