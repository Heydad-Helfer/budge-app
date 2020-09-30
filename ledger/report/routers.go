package report

import "github.com/gin-gonic/gin"

//Route adds route to the main router
func Route(r *gin.RouterGroup) {
	r.POST("/report", endpoint)
}
