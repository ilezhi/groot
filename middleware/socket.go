package middleware

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

// 保存每个客户端连接信息
type Client struct {
	id uint

	hub *Hub

	send chan interface{}

	conn *websocket.Conn
}

type Params interface{}

type Message struct {
	to []uint
	data interface{}
}

func (c *Client) readPump() {
	defer func() {
		c.hub.unregister <- c
		c.conn.Close()
	}()

	for {
		var m Params
		err := c.conn.ReadJSON(&m)
		if err != nil {
			fmt.Println("readjson error", err)
			break
		}

		fmt.Println("readjson success", m)
		c.Others(m)
	}
}

func (c *Client) writePump() {
	defer func() {
		c.conn.Close()
	}()

	for {
		select {
		case message, ok := <- c.send:
			if !ok {
				return
			}

			c.conn.WriteJSON(message)
		}
	}
}

func (c *Client) To(id uint, message interface{}) {
	msg := &Message{
		data: message,
		to: []uint{id},
	}
	c.hub.broadcast <- msg
}

func (c *Client) All(message interface{}) {
	len := len(c.hub.clients)
	ids := make([]uint, 0, len - 1)

	for id := range c.hub.clients {
		ids = append(ids, id)
	}
	msg := &Message{
		data: message,
		to: ids,
	}

	c.hub.broadcast <- msg
}

func (c *Client) Others(message interface{}) {
	len := len(c.hub.clients)
	ids := make([]uint, 0, len - 1)

	for id := range c.hub.clients {
		if id == c.id {
			continue
		}
		ids = append(ids, id)
	}
	msg := &Message{
		data: message,
		to: ids,
	}

	c.hub.broadcast <- msg
}

// 1. 所有人
// 2. 除了自己的所有人
// 3. 指定要发送的人
// TODO:
// 4. 分组发送
type Hub struct {
	clients map[uint]*Client

	broadcast chan *Message

	register chan *Client

	unregister chan *Client
}

func (h *Hub) run() {
	for {
		select {
		case client := <- h.register:
			h.clients[client.id] = client

		case client := <- h.unregister:
			if _, ok := h.clients[client.id]; ok {
				delete(h.clients, client.id)
				close(client.send)
			}

		case message := <- h.broadcast:
			h.send(message)
		}
	}
}

func (h *Hub) send(message *Message) {
	ids := message.to

	for _, id := range ids {
		client := h.clients[id]
		select {
		case client.send <- message.data:
		default:
			close(client.send)
			delete(h.clients, id)
		}
	}
}

func newHub() *Hub {
	return &Hub{
		clients: 		make(map[uint]*Client),
		register: 	make(chan *Client),
		unregister: make(chan *Client),
		broadcast: 		make(chan *Message),
	}
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

var hub *Hub

func init() {
	hub = newHub()
	go hub.run()
}

func WSConn(ctx *Context) {
	id, _ := ctx.Params().GetInt("id")
	conn, err := upgrader.Upgrade(ctx.ResponseWriter(), ctx.Request(), nil)
	if err != nil {
		log.Println(err)
		return
	}

	client := &Client{id: uint(id) , hub: hub, conn: conn, send: make(chan interface{})}
	client.hub.register <- client

	go client.readPump()
	go client.writePump()
}
