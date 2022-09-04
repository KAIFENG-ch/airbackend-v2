package api

import (
	"github.com/gin-gonic/gin"
	"test/dao"
	"test/utils"
)

func Register(c *gin.Context) {
	dbName := utils.Generate()
	resp := dao.CreateDb(c.PostForm("enterprise"), dbName, c.PostForm("username"), c.PostForm("password"))
	if resp.Status != 200 {
		c.JSON(400, "no response")
	}
	c.JSON(200, "SUCCESS")
}
