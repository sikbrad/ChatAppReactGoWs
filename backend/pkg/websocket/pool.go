package websocket

import "fmt"

type Pool struct{
	Register		chan *Client
	Unregister		chan *Client
	Clients			map[*Client]bool
	Broadcast		chan Message


}

func (pool* Pool) Start() {
	for{
		select{ //iterate forever
		case client := <- pool.Register: //if new register msg
			pool.Clients[client] = true //set hashtable as true
			fmt.Println("Registered 1 client -> Size of connection pool is : ", len(pool.Clients))
			for key, elem := range pool.Clients{ //iter for all clients,
				cli := key
				_ = elem
				fmt.Printf("%+v",cli)
				cli.Conn.WriteJSON(Message{
					Type:1,
					Body: "New user joined",
				})
			}
			break

		case client:= <- pool.Unregister:
			delete(pool.Clients, client) //delete from hash
			fmt.Println("UnRegistered 1 client -> Size of a connection pool after del : ", len(pool.Clients))
			for client, _ := range pool.Clients{
				client.Conn.WriteJSON(Message{
					Type:1,
					Body:fmt.Sprintf("User %v disconnected", client),
				})
			}
			break

		case message := <- pool.Broadcast:
			fmt.Println("Sending message to all clients in pool")
			for cli, _ := range pool.Clients{
				if err := cli.Conn.WriteJSON(message); err != nil{
					fmt.Println(err)
					return
				}
			}
		}
	}
}

func NewPool() *Pool{
	return &Pool{
		Register:		make(chan* Client),
		Unregister: 	make(chan* Client),
		Clients:		make(map[*Client]bool),
		Broadcast:		make(chan Message),
	}
}