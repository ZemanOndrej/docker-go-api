package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	public := router.Group("/api")
	{
		public.GET("/", func(c *gin.Context) {
			c.String(http.StatusOK, "Welcome to sample dockerized golang api")
		})
	}
	router.Run(":5513")
}
