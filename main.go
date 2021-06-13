package main

import (
	"github.com/gin-gonic/gin"
	"unique_hackday/conf"
	"unique_hackday/handler"
)

func main() {
	r := gin.Default()
	gin.SetMode(conf.RunMode)

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"msg": "pong",
		})
	})

	scenes := r.Group("/scenes")
	{
		scenes.POST("",handler.AddSceneHandler)
		scenes.GET("",handler.GetSceneListByTagNameHandler)
	}

	tags := r.Group("/tags")
	{
		tags.POST("",handler.AddTagHandler)
		tags.GET("",handler.GetTagListBySceneNameHandler)
	}
	r.Run()
}
