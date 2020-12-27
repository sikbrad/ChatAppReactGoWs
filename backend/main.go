package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/sikbrad/ChatAppReactGoWs/pkg/websocket"
)


//define websocket endpoint
func serveWs(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Host : %v\n", r.Host)

	ws, err := websocket.Upgrade(w, r)
	if err != nil {
		log.Println(err)
	}

	//listen indefinitely
	//reader(ws)

	// reader thing is relocated
	go websocket.Writer(ws)
	//read indefinitely?
	websocket.Reader(ws)

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
