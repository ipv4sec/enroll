package user

import (
	"enroll/config"
	"enroll/csv"
	"github.com/gin-gonic/gin"
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