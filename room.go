package main

type room struct {
	forward chan []byte
	join    chan *client
	leave   chan *client
	clients map[*client]bool
}

func (r *room) run() {
	for {
		select {
		case client := <-r.join:
			//join
			r.clients[client] = true
		case client := <-r.leave:
			//leave
			delete(r.clients, client)
			close(client.send)
		case msg := <-r.forward:
			// forwarding the message to all clients
			for client := range r.clients {
				select {
				case client.send <- msg:
					//send a message
				default:
					//fail to send
					delete(r.clients, client)
					close(client.send)
				}
			}
		}
	}
}
