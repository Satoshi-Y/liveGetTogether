package main

import (
	"fmt"
	"io"
	"math/rand"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// load all templates file
	router.LoadHTMLGlob("templates/*")

	// set routing.
	v1 := router.Group("/v1")
	{
		v1.GET("/room/:roomid", roomGET,)
		v1.POST("/room/:roomid", roomPOST)
		v1.DELETE("/room/:roomid", roomDELETE)
		v1.GET("/stream/:roomid", stream)
	}


	// start server
	router.Run(":8080")
}

func stream(c *gin.Context) {
	roomid := c.Param("roomid")
	listener := openListener(roomid)
	defer closeListener(roomid, listener)

	c.Stream(func(w io.Writer) bool {
		c.SSEvent("message", <-listener)
		return true
	})
}

func roomGET(c *gin.Context) {
	roomid := c.Param("roomid")
	userid := fmt.Sprint(rand.Int31())
	c.HTML(200, "index.tmpl", gin.H{
		"roomid": roomid,
		"userid": userid,
	})
}

func roomPOST(c *gin.Context) {
	roomid := c.Param("roomid")
	userid := c.PostForm("user")
	message := c.PostForm("message")
	room(roomid).Submit(userid + ": " + message)

	c.JSON(200, gin.H{
		"status":  "success",
		"message": message,
	})
}

func roomDELETE(c *gin.Context) {
	roomid := c.Param("roomid")
	deleteBroadcast(roomid)
}