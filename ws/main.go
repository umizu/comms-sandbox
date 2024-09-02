package main

import (
	"fmt"
	"log"
	"net"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true // testing from chrome
	},
}

func main() {
	wsService := &WebSocketService{
		conns: make(map[string]*websocket.Conn),
	}

	http.HandleFunc("/", wsService.Handler)

	port := 8080
	log.Printf("listening on port :%d", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}

func (s *WebSocketService) Handler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print(err)
	}
	uuid := uuid.NewString()
	s.conns[uuid] = conn

	defer func() {
		delete(s.conns, uuid)
	}()

	go s.NotifyAll(conn.NetConn())

	for {
		mt, message, err := conn.ReadMessage()
		if err != nil {
			log.Printf("read:%v", err)
			break
		}

		log.Printf("recv: %s", message)
		err = conn.WriteMessage(mt, message)
		if err != nil {
			log.Printf("write:%v", err)
			break
		}
	}
}

func (s *WebSocketService) NotifyAll(newConn net.Conn) {
	for _, conn := range s.conns {
		conn.WriteJSON(fmt.Sprintf("new client connected: %s", newConn.RemoteAddr()))
	}
}

type WebSocketService struct {
	conns map[string]*websocket.Conn
}
