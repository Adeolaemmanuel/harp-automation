package handler

import (
	"github.com/gin-gonic/gin"
	"harp-automation.com/com/modules/automation"
)

var autoRouter = automation.Automation{}



func Handler(c *gin.Context) {
	server := gin.Default()

	autoRouter.AutomationRouter(server.Group("/api/v1/automation"))

	server.ServeHTTP(c.Writer, c.Request)
}