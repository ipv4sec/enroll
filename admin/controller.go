package admin

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"strconv"
)

func Login(c *gin.Context) {
	var param struct {
		Name string `json:"name" binding:"required"`
		Pass string `json:"pass" binding:"required"`
	}
	if err := c.BindJSON(&param); err != nil {
		c.JSON(400, gin.H {
			"error": err.Error(),
			"message": "参数错误",
		})
		return
	} else {
		token, user, err := GenTokenByNameAndPass(param.Name, param.Pass)
		if err != nil {
			c.JSON(400, gin.H {
				"error": err.Error(),
				"message": "服务器内部错误",
			})
			return
		} else {
			c.JSON(200, gin.H{
				"user": user,
				"token": token,
				"message": "登录成功",
			})
		}
	}
}

func GetAll(c *gin.Context) {
	result, err := GetAllAdmins()
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
			"message": "服务器内部错误",
		})
		return
	}
	c.JSON(200, gin.H{
		"data": result,
		"message": "获取全部老师成功",
	})
}

func ChangePassword(c *gin.Context) {
	userIdStr := c.Query("userId")
	userId, err := strconv.ParseInt(userIdStr, 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
			"message": "参数错误",
		})
		return
	}
	var param struct {
		OldPass string `json:"oldPass" binding:"required"`
		NewPass string `json:"newPass" binding:"required"`
	}
	if err := c.BindJSON(&param); err != nil {
		c.JSON(400, gin.H {
			"error": err.Error(),
			"message": "参数错误",
		})
		return
	} else {
		err := ChangePasswordByUserId(userId, param.OldPass, param.NewPass)
		if err != nil {
			c.JSON(400, gin.H {
				"error": err.Error(),
				"message": "密码更改失败",
			})
			return
		} else {
			c.JSON(200, gin.H{
				"message": "密码更改成功",
			})
		}
	}
}

func DeleteByAdminId(c *gin.Context) {
	adminIdStr := c.Query("adminId")
	adminId, err := strconv.ParseInt(adminIdStr, 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
			"message": "参数错误",
		})
		return
	}
	err = DeleteAdminByAdminId(adminId)
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
func ChangePasswordByAdminId(c *gin.Context) {
	var param struct{
		Id json.Number `json:"id" binding:"required"`
		Pass string `json:"pass" binding:"required"`
	}
	if err := c.BindJSON(&param); err != nil {
		c.JSON(400, gin.H {
			"error": err.Error(),
			"message": "参数错误",
		})
		return
	} else {
		id, err := param.Id.Int64()
		if err != nil {
			c.JSON(400, gin.H {
				"error": err.Error(),
				"message": "参数错误",
			})
			return
		}
		err = ChangePasswd(id, param.Pass)
		if err != nil {
			c.JSON(400, gin.H{
				"error": err.Error(),
				"message": "修改密码失败",
			})
			return
		}
		c.JSON(200, gin.H{
			"message": "修改密码成功",
		})
	}
}
func Add(c *gin.Context) {
	var param struct {
		Name string `json:"name" binding:"required"`
		Pass string `json:"pass" binding:"required"`
		SiteId json.Number `json:"siteId" binding:"required"`
	}
	if err := c.BindJSON(&param); err != nil {
		c.JSON(400, gin.H {
			"error": err.Error(),
			"message": "参数错误",
		})
		return
	} else {
		siteId, err := param.SiteId.Int64()
		if err != nil {
			c.JSON(400, gin.H {
				"error": err.Error(),
				"message": "参数错误",
			})
			return
		}
		err = AddAdminByNameAndPassAndSiteId(param.Name, param.Pass, siteId)
		if err != nil {
			c.JSON(400, gin.H{
				"error": err.Error(),
				"message": "增加老师失败",
			})
			return
		}
		c.JSON(200, gin.H{
			"message": "增加老师成功",
		})
	}
}