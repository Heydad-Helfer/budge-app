package main

import (
	"log"

	"github.com/gin-gonic/gin"
)

const (
	userkey = "user"
)

//AuthRequired blocks all un-authenticated requests
func authRequired(c *gin.Context) {
	// Continue down the chain to handler etc
	log.Printf("authCHeck")
	c.Next()
}
