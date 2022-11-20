package main

import (
	"EventManagement/config"
	"github.com/gin-gonic/gin"
)

func main() {

	gin.SetMode(gin.ReleaseMode)
	config.InitialMigrationForStaging()
	router := gin.New()
	InitializeRoutes(router)
	err := router.Run(":8081")
	if err != nil {
		return
	}

}
