package websocket

import (
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

func Upgrade(w http.ResponseWriter, r *http.Request) (*websocket.Conn, error) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return ws, err
	}
	return ws, nil
}


//reader and writer is relocated to websocket package's start selector
//after tut 4
//func Reader(conn *websocket.Conn) {
//	for {
//		//read
//		messageType, p, err := conn.ReadMessage()
//		if err != nil {
//			log.Println(err)
//			return
//		}
//
//		//print
//		fmt.Println(string(p))
//
//		if err := conn.WriteMessage(messageType, p); err != nil {
//			log.Println(err)
//			return
//		}
//	}
//}
//
//
//func Writer(conn *websocket.Conn){
//	for{
//		fmt.Println("Sending")
//		messageType, r, err := conn.NextReader()
//
//		if err!= nil{
//			fmt.Println(err)
//			return
//		}
//
//		w, err := conn.NextWriter(messageType)
//		if err != nil{
//			fmt.Println(err)
//			return
//		}
//
//		if _, err := io.Copy(w,r); err != nil{
//			fmt.Println(err)
//			return
//		}
//
//		if err := w.Close(); err != nil{
//			fmt.Println(err)
//			return
//		}
//	}
//}