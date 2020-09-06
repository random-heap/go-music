package main

import (
	"github.com/gin-gonic/gin"
	"go-music/controller"
)

func main() {

	router := gin.Default()

	router.POST("/admin/login/status", controller.VerityPasswd)
	router.POST("/collection/add", controller.AddCollection)

	router.Run()
}
