package adapter

import (
	"fmt"
	"io"

	"github.com/RamisL/server/broadcast"
	"github.com/RamisL/server/payment"
	"github.com/gin-gonic/gin"
)

type GinAdapter interface {
	Stream(c *gin.Context)
	PostRoom(c *gin.Context)
	DeleteRoom(c *gin.Context)
}

type MessageInput struct {
	Text  string
	Name  string
	Price string
}

type response struct {
	Success bool   `json:"success,omitempty"`
	Message string `json:"message,omitempty"`
	Data    string `json:"data,omitempty"`
}

type ginAdapter struct {
	broadcaster broadcast.Broadcaster
	paymentService payment.Service
}

func NewGinAdapter(broadcaster broadcast.Broadcaster, paymentService payment.Service) *ginAdapter {
	return &ginAdapter{
		broadcaster: broadcaster,
		paymentService: paymentService,
	}
}

// Stream godoc
// @Summary      Stream messages
// @Description  Stream messages from a room
// @Tags         chat
// @Produce      text/event-stream
// @Param        id   path      string  true  "Room ID"
// @Failure      400  {object}  HTTPError
// @Failure      500  {object}  HTTPError
// @Router       /stream/{id} [get]
func (ga *ginAdapter) Stream(c *gin.Context) {
	//create a new channel to handle the stream
	listener := make(chan interface{})

	// get the broadcaster

	ga.broadcaster.Register(listener)

	//close the channel when error message or client is gone
	defer ga.broadcaster.Unregister(listener)

	clientGone := c.Request.Context().Done()

	c.Stream(func(w io.Writer) bool {
		select {
		case <-clientGone:
			return false
		case message := <-listener:
			serviceMsg, ok := message.(MessageInput)
			if !ok {
				fmt.Println("not a message")
				c.SSEvent("message", message)
				return false
			}
			c.SSEvent("message", " "+serviceMsg.Text+" -> Nom: "+serviceMsg.Name+" -> Prix: "+serviceMsg.Price)
			return true
		}
	})
}