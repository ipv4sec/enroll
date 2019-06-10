package csv

import (
	"enroll/config"
	"github.com/gin-gonic/gin"
	"github.com/satori/go.uuid"
	"io"
	"os"
)

func Upload(c *gin.Context) {
	uploaded, err := c.FormFile("filepond")
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
			"message": "上传失败",
		})
		return
	}
	newFilename := uuid.NewV4().String() + ".csv"
	newFile, err := os.OpenFile(config.Conf.Csv.Uploaded + newFilename, os.O_TRUNC|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
			"message": "上传失败",
		})
		return
	}
	file, err := uploaded.Open()
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
			"message": "上传失败",
		})
		return
	}
	_, err = io.Copy(newFile, file)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
			"message": "上传失败",
		})
		return
	}
	newFile.Close()
	file.Close()

	//adminId, ok := c.Get("adminId")
	//logger.Info(adminId, ok)
	//if !ok {
	//	adminId = int64(-1) // 错误的上传记录
	//}
	//err = SaveNameAndAdminId(newFilename, adminId.(int64))
	//if err != nil {
	//	c.JSON(400, gin.H{
	//		"error": err.Error(),
	//		"message": "上传失败",
	//	})
	//	return
	//} 因为上传问价没做校验, 暂时放弃
	c.String(200, newFilename)
	//c.JSON(400, gin.H{
	//	"error": "sss",
	//	"message": "上传失败",
	//})
}

func Parse(c *gin.Context) {
	csvFilename := c.Query("filename")

	result, err := Read(config.Conf.Csv.Uploaded + csvFilename)
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
	c.JSON(200, gin.H{
		"data": result[1:],
		"message": "解析成功",
	})
}