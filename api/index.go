package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"harp-automation.com/com/modules/automation"
)

var autoRouter = automation.Automation{}



func Handler(w http.ResponseWriter, r *http.Request) {
	server := gin.Default()

	autoRouter.AutomationRouter(server.Group("/api/v1/automation"))

	server.ServeHTTP(w, r)
}