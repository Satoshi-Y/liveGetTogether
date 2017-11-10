package main

import (
	"github.com/gin-gonic/gin"
	"github.com/Satoshi-Y/liveGetTogether/lgt/controller"
)

func main() {
	router := gin.Default()

	// load all templates file
	router.LoadHTMLGlob("templates/*")

	// set routing.
	v1 := router.Group("/v1")
	{
		v1.GET("/room/:roomid", controller.RoomGET)
		v1.POST("/room/:roomid", controller.RoomPOST)
		v1.DELETE("/room/:roomid", controller.RoomDELETE)
		v1.GET("/roomStream/:roomid", controller.RoomStream)
	}


	// start server
	router.Run(":8080")
}

