package automation

import (
	"github.com/gin-gonic/gin"
	"harp-automation.com/com/modules/controls"
)

var con = controls.Controls{}

func (a Automation) AutomationRouter(r *gin.RouterGroup) {
	r.POST("/", combinatorHandler)
}

func combinatorHandler(c *gin.Context) {
	payload := controls.Controls{}

	if err := c.ShouldBind(&payload); err == nil {
		ok, controlErr := con.Control(payload)

		if controlErr != nil {
			c.JSON(400, gin.H{
				"error": controlErr.Error(),
				"message": "An error occurred",
				"status": false,
			})		
			return	
		}

		c.JSON(200, gin.H{
			"status": true,
			"data": ok,
			"message": "Success",
		})

	} else {
		c.JSON(400, gin.H{
			"message": "Invalid body provided",
			"error": "An error occurred",
			"status": false,
		})
		return
	}
}