package main

import (
	"math/rand"
	"time"

	"github.com/alfred-zhong/prome"
	"github.com/gin-gonic/gin"
)

func main() {
	client := prome.NewClient("test", "")
	client.EnableRuntime = false

	e := gin.New()
	e.Use(client.MiddlewareRequestCount(""))
	e.Use(client.MiddlewareRequestDuration("", nil))

	e.GET("/foo", gin.WrapH(client.Handler()))
	e.GET("/hello/:name", func(c *gin.Context) {
		c.String(200, "indeed, %s", c.Param("name"))
	})
	e.GET("/sleep", func(c *gin.Context) {
		time.Sleep(time.Duration(rand.Int63n(1000)) * time.Millisecond)
	})
	if err := e.Run(":9527"); err != nil {
		panic(err)
	}
}
