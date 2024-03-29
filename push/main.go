package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

func main() {
	dr := &DataReceiver{
		conns: []*websocket.Conn{},
	}

	http.HandleFunc("/ws", dr.handleWS)
	fmt.Println("starting server on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

type DataReceiver struct {
	conns []*websocket.Conn
}

func (dr *DataReceiver) handleWS(w http.ResponseWriter, r *http.Request) {
	u := websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}

	conn, err := u.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal(err)
	}

	go dr.onConnect(conn)
	dr.conns = append(dr.conns, conn)
	go dr.receiveLoop(conn)
}

func (dr *DataReceiver) receiveLoop(conn *websocket.Conn) {
	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			log.Fatal(err)
		}

		for _, c := range dr.conns {
			err := c.WriteMessage(websocket.TextMessage, msg)
			if err != nil {
				log.Fatal(err)
			}
		}
	}
}

func (dr *DataReceiver) onConnect(conn *websocket.Conn) {
	for _, c := range dr.conns {
		err := c.WriteMessage(
			websocket.TextMessage,
			[]byte(fmt.Sprintf("%s connected", conn.RemoteAddr().String())))
		if err != nil {
			log.Fatal(err)
		}
	}
}
