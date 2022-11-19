package main

import (
	"EventManagement/config"
	"github.com/gin-gonic/gin"
)

var (
	router *gin.Engine
	key    = "1234567891234567"
)

func main() {

	gin.SetMode(gin.ReleaseMode)

	config.InitialMigrationForStaging()
	router := gin.New()
	//router.Use(
	//	gin.LoggerWithWriter(gin.DefaultWriter, "/ugc-terms-and-conditions/1/1/1"),
	//	gin.Recovery(),
	//)
	//router.Use(common.Logger())

	InitializeRoutes(router)

	err := router.Run(":8081")
	if err != nil {
		return
	}

}
