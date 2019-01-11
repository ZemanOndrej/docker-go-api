package main

import (
	"net/http"

	"docker-api/api/src/db"
	"log"

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
	ctx := db.GetDBContext()
	err := ctx.GetDB().Ping()
	if err != nil {
		log.Fatal(err)
		panic("db ping didnt work")

	}
	defer ctx.Close()
	router.Run()
}
