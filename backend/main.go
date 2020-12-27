package main

import (
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

//define websocket endpoint
func serveWs(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Host : %v\n", r.Host)

	ws, err := upgrader.Upgrade(w, r, nil)

	if err != nil {
		log.Println(err)
	}

	//listen indefinitely
	reader(ws)
}

func reader(conn *websocket.Conn) {
	for {
		//read
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}

		//print
		fmt.Println(string(p))

		if err := conn.WriteMessage(messageType, p); err != nil {
			log.Println(err)
			return
		}
	}
}

func main() {
	fmt.Println("Chat app v0.0.1 started")

	setupRoutes()
	http.ListenAndServe(":8080", nil)

	fmt.Println("Chat app v0.0.1 finished")
}

func setupRoutes() {
	http.HandleFunc("/",
		func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintf(w, "Simple server is open")
		})

	http.HandleFunc("/ws", serveWs)
}
