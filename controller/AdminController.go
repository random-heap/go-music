package controller

import (
	"github.com/gin-gonic/gin"
	"go-music/dao"
	"net/http"
)

func VerityPasswd(context *gin.Context) {

	name := context.Query("name")
	password := context.Query("password")

	adminDao := new(dao.AdminDao)
	count := adminDao.SelectByNameAndPassword(name, password)

	if count > 0 {
		context.JSON(http.StatusOK, gin.H{
			"code":1,
			"msg":"登录成功",

		})
		context.Set("name", name)
	} else {
		context.JSON(http.StatusOK, gin.H{
			"code":0,
			"msg":"用户名或密码错误",
		})
	}


}