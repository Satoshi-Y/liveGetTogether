package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/Satoshi-Y/liveGetTogether/lgt/modules"
	"io"
	"fmt"
	"math/rand"
)

func RoomStream(c *gin.Context) {
	roomid := c.Param("roomid")
	listener := modules.OpenListener(roomid)
	defer modules.CloseListener(roomid, listener)

	c.Stream(func(w io.Writer) bool {
		c.SSEvent("message", <-listener)
		return true
	})
}

func RoomGET(c *gin.Context) {
	roomid := c.Param("roomid")
	userid := fmt.Sprint(rand.Int31())

	c.HTML(200, "index.tmpl", gin.H{
		"roomid": roomid,
		"userid": userid,
	})
}

func RoomPOST(c *gin.Context) {
	roomid := c.Param("roomid")
	userid := c.PostForm("user")
	message := c.PostForm("message")
	modules.Room(roomid).Submit(userid + ": " + message)

	c.JSON(200, gin.H{
		"status":  "success",
		"message": message,
	})
}

func RoomDELETE(c *gin.Context) {
	roomid := c.Param("roomid")
	modules.DeleteBroadcast(roomid)
}