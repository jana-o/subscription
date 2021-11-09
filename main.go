package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jana-o/subscription/app/handlers"
	"net/http"
)

func main() {
	// init router
	router := gin.Default()
	router.Use(gin.Recovery())

	router.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{"code":"PAGE_NOT_FOUND", "message":"Page not found"}) // 404
	})

	// route handlers / endpoints
	handlers.Routes(router)

}
