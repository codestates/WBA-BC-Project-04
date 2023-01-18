package router

import (
	"fmt"
	ctl "lecture/WBA-BC-Project-04/contorller"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

type Router struct {
	ct *ctl.Controller
}

func NewRouter(ct *ctl.Controller) (*Router, error) {
	r := &Router{
		ct: ct,
	}

	return r, nil
}

func CORS() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, X-Forwarded-For, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	}
}
func liteAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c == nil {
			c.Abort()
			return
		}
		auth := c.GetHeader("Authorization")
		fmt.Println("Authorization-word ", auth)

		c.Next()
	}
}

func (p *Router) Idx() *gin.Engine {

	e := gin.Default()
	e.Use(gin.Logger())
	e.Use(gin.Recovery())
	e.Use(CORS())

	e.GET("/health")

	papi := e.Group("api/v1", liteAuth())
	{
		game := papi.Group("game", liteAuth())
		{
			game.POST("/match", p.ct.GamePostMatchController) // match open을 위한 API
			// game.POST("/match/join", p.ct.GetTest)            // match join을 위한 API
			game.POST("/match/end", p.ct.GameEndMatchController) // match end를 위한 API
		}

		waitingRoom := papi.Group("waiting-room", liteAuth())
		{
			waitingRoom.GET("/", p.ct.GetWaitingRoomTest)
		}

	}

	websocket := e.Group("ws", CORS())
	{
		websocket.GET("", func(c *gin.Context) {
			wshandler(c.Writer, c.Request)
		})
	}

	return e
}

var wsupgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func wshandler(w http.ResponseWriter, r *http.Request) {
	conn, err := wsupgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("Failed to set websocket upgrade: %+v", err)
		return
	}

	for {
		t, msg, err := conn.ReadMessage()
		if err != nil {
			break
		}
		conn.WriteMessage(t, msg)
	}
}
