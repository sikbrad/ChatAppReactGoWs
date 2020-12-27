package main

import (
	"fmt"
	"github.com/sikbrad/ChatAppReactGoWs/pkg/websocket"
	"net/http"
)


//define websocket endpoint
func serveWs(w http.ResponseWriter, r *http.Request, pool *websocket.Pool) {
	fmt.Printf("Host : %v\n", r.Host)

	//ws, err := websocket.Upgrade(w, r)
	//if err != nil {
	//	log.Println(err)
	//}

	//listen indefinitely
	//reader(ws)

	// reader thing is relocated to websocket
	//go websocket.Writer(ws)
	////read indefinitely?
	//websocket.Reader(ws)

	//and it again relocated to websocket package, where Start resides.
	//so above code is long gone now


	// after tut 4..
	conn, err := websocket.Upgrade(w,r)
	if err != nil {
		fmt.Fprint(w, "%+v\n", err)
	}

	client := &websocket.Client{
		Conn: conn,
		Pool: pool,
	}

	//send chan that new client is registerd
	pool.Register <- client
	client.Read()

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

	//ws pools
	pool := websocket.NewPool()
	go pool.Start()

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request){
		serveWs(w,r,pool)
	})
}
