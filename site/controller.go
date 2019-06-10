package site

import (
	"github.com/gin-gonic/gin"
	"strconv"
)

func GetAll(c *gin.Context) {
	// adminId, _ := c.Get("adminId")
	result, err := GetAllSites()
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
			"message": "服务器内部错误",
		})
		return
	}
	c.JSON(200, gin.H{
		"data": result,
		"message": "获取全部站点成功",
	})
}


func DeleteBySiteId(c *gin.Context) {
	siteIdStr := c.Query("siteId")
	siteId, err := strconv.ParseInt(siteIdStr, 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
			"message": "参数错误",
		})
		return
	}
	err = DeleteSiteBySiteId(siteId)
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

func Add(c *gin.Context) {
	var param struct {
		Name string `json:"name" binding:"required"`
	}
	if err := c.BindJSON(&param); err != nil {
		c.JSON(400, gin.H {
			"error": err.Error(),
			"message": "参数错误",
		})
		return
	} else {
		err = AddSiteByName(param.Name)
		if err != nil {
			c.JSON(400, gin.H{
				"error": err.Error(),
				"message": "增加站点失败",
			})
			return
		}
		c.JSON(200, gin.H{
			"message": "增加站点成功",
		})
	}
}
