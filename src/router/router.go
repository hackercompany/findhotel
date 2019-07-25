package router

import (
	"api"

	"github.com/gin-gonic/gin"
)

// Wrapper to add list of URLs to a given router.
func AddRoutingUrls(router *gin.Engine) {
	router.GET("/ping", api.Ping)
	router.GET("/ipinformation", api.IpInfo)
}
