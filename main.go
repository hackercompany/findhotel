package main

import (
	"fmt"

	"config"
	"logger"
	"middleware"
	"router"

	"github.com/gin-gonic/gin"
)

func main() {
	config.DoInit()
	logger.DoInit()
	middleware.DoInit()

	mainRouter := gin.Default()

	mainRouter.Use(middleware.MySQLConnector)

	router.AddRoutingUrls(mainRouter)

	mainRouter.Run(fmt.Sprintf(":%s", config.Config.GetString("server.port")))

}
