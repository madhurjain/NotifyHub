package main

import (
	"encoding/json"
	"log"
)

type message struct {
	uid     string
	payload []byte
}

// SocketConnManager maintains the set of active
// connections and sends messages to connections
type SocketConnManager struct {
	// registered connections
	// links connection to uid
	connections map[string]*connection
	// Adds new peer
	add chan *connection
	// Removes peer
	remove chan *connection
	// Sends notification message to peer
	send chan message
}

var socketConnManager = SocketConnManager{
	connections: make(map[string]*connection),
	add:         make(chan *connection),
	remove:      make(chan *connection),
	send:        make(chan message),
}

func InitWebSockets() {
	// Start Socket Connection Manager
	go socketConnManager.run()
}

// runs as a goroutine and manages adding / removing
// connections to connection manager and sending response
func (socketConnManager *SocketConnManager) run() {
	for {
		select {
		case conn := <-socketConnManager.add:
			//log.Println("Add connection", conn.uid)
			socketConnManager.connections[conn.uid] = conn
			UpdateActiveConnections()
		case conn := <-socketConnManager.remove:
			//log.Println("Remove connection", conn.uid)
			if _, ok := socketConnManager.connections[conn.uid]; ok {
				delete(socketConnManager.connections, conn.uid)
				close(conn.send)
			}
			UpdateActiveConnections()
		case msg := <-socketConnManager.send:
			//log.Println("sending notification", msg.uid, string(msg.payload))
			if conn, ok := socketConnManager.connections[msg.uid]; ok {
				select {
				case conn.send <- msg.payload:
				default:
					delete(socketConnManager.connections, conn.uid)
					close(conn.send)
				}
			}
		}
	}
}

func SendMessage(msg *message) {
	go func(m *message) {
		socketConnManager.send <- *m
	}(msg)
}

// Send list of peers to dashboard via WebSockets
func UpdateActiveConnections() {
	peers := []string{}
	for uid, _ := range socketConnManager.connections {
		if uid == DASHBOARD_UID {
			continue
		} // ignore dashboard user in list
		peers = append(peers, uid)
	}

	data, err := json.Marshal(peers)
	if err != nil {
		log.Println(err)
		return
	}
	// send list of peers to dashboard websocket connection
	msg := &message{uid: DASHBOARD_UID, payload: data}
	SendMessage(msg)
}
