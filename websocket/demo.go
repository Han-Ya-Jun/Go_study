package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gopkg.in/olahol/melody.v1"
	"net/http"
)

func main() {
	r := gin.Default()
	m := melody.New()

	r.GET("/", func(c *gin.Context) {
		http.ServeFile(c.Writer, c.Request, "index.html")
	})

	r.GET("/ws", func(c *gin.Context) {
		m.HandleRequest(c.Writer, c.Request)
	})

	m.HandleMessage(func(s *melody.Session, msg []byte) {
		fmt.Printf("received message:%s\n", string(msg))
		m.Broadcast(msg)
		println(msg)
	})
	m.HandleConnect(func(session *melody.Session) {
		println("Connected")
	})
	m.HandleClose(func(session *melody.Session, i int, s string) error {
		println("close")
		return nil
	})
	r.Run(":5000")
}
