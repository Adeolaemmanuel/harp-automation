package main

import (
	"github.com/gin-gonic/gin"
	"harp-automation.com/com/modules/automation"
)

var autoRouter = automation.Automation{}

func Main () {
	server := gin.Default()

	autoRouter.AutomationRouter(server.Group("/api/v1/automation"))

	server.Run(":3000")
}