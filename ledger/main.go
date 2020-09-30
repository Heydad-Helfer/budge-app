package main

import (
	"net/http"

	"github.com/Heydad-Helfer/budge-app/ledger/report"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Budge Ledger")
	})

	rAPI := r.Group("/api/ledger", authRequired)
	report.Route(rAPI)

	r.Run(":8080")
}
