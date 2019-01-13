package main

import (
	"go-docker-api/api/db"
	"net/http"
	"time"

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
	time.Sleep(20 * time.Second) //sleep until db image is running
	ctx := db.GetDBContext()
	err := ctx.GetDB().Ping()
	if err != nil {
		log.Fatal(err)
		panic("db ping didnt work")

	}
	defer ctx.Close()
	router.Run()
}
