// package main

// import (
// 	"log"
// 	"net/http"
// 	"time"

// 	"golang.org/x/net/http2"
// )

// // Chatroom is a type that contains a map of clients and a channel for messages
// type Chatroom struct {
// 	clients map[*Client]bool
// 	message chan string
// }

// // Client is a type that represents a connection to the chatroom
// type Client struct {
// 	conn      *http.Conn
// 	send      chan string
// 	room      *Chatroom
// 	flusher   http2.Flusher
// 	connected bool
// }

// // NewChatroom returns a new instance of Chatroom
// func NewChatroom() *Chatroom {
// 	return &Chatroom{
// 		clients: make(map[*Client]bool),
// 		message: make(chan string),
// 	}
// }

// // ServeHTTP serves the chatroom
// func (room *Chatroom) ServeHTTP(w http.ResponseWriter, r *http.Request) {
// 	if r.Method != http.MethodGet {
// 		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
// 		return
// 	}

// 	// Create a new client
// 	c := &Client{
// 		conn:      r.Context().Value("conn").(*http.Conn),
// 		send:      make(chan string),
// 		room:      room,
// 		flusher:   w.(http2.Flusher),
// 		connected: true,
// 	}

// 	// Add the client to the list of clients
// 	room.clients[c] = true

// 	defer func() {
// 		room.remove(c)
// 		c.conn.Close()
// 	}()

// 	// Send any messages that come in
// 	go c.sendMessage()

// 	// Read any messages that come in
// 	go c.readMessage()

// 	// Set the write timeout so we don't block the response
// 	c.conn.SetWriteDeadline(time.Now().Add(time.Second * 10))

// 	// Write the response
// 	if err := c.writeResponse(); err != nil {
// 		log.Println(err)
// 	}
// }

// // Send sends a message to the chatroom
// func (room *Chatroom) Send(msg string) {
// 	room.message <- msg
// }

// // Remove removes a client from the chatroom
// func (room *Chatroom) remove(c *Client) {
// 	delete(room.clients, c)
// }

// // ReadMessage reads messages from the client
// func (c *Client) readMessage() {
// 	for {
// 		select {
// 		case message := <-c.room.message:
// 			c.send <- message
// 		case <-c.conn.Context().Done():
// 			c.connected = false
// 			return
// 		}
// 	}
// }

// // SendMessage sends messages to the client
// func (c *Client) sendMessage() {
// 	for {
// 		select {
// 		case message := <-c.send:
// 			if _, err := c.conn.Write([]byte(message)); err != nil {
// 				continue
// 			}
// 			c.flusher.Flush()
// 		case <-c.conn.Context().Done():
// 			c.connected = false
// 			return
// 		}
// 	}
// }

// // writeResponse writes the initial response
// func (c *Client) writeResponse() error {
// 	_, err := c.conn.Write([]byte("Welcome to the chatroom!\n"))
// 	if err != nil {
// 		return err
// 	}
// 	c.flusher.Flush()
// 	return nil
// }

// func main() {
// 	room := NewChatroom()
// 	http.HandleFunc("/", room.ServeHTTP)
// 	if err := http.ListenAndServe(":8080", nil); err != nil {
// 		log.Fatal(err)
// 	}
// }
package main

import (
	"log"
	"net/http"
)

func main() {
	// 在 8000 端口启动服务器
	// 确切地说，如何运行HTTP/1.1服务器。

	srv := &http.Server{Addr: ":8000", Handler: http.HandlerFunc(handle)}

	// 用TLS启动服务器，因为我们运行的是http/2，它必须是与TLS一起运行。
	// 确切地说，如何使用TLS连接运行HTTP/1.1服务器。
	log.Printf("Serving on https://0.0.0.0:8000")
	log.Fatal(srv.ListenAndServeTLS("server.crt", "server.key"))
}

// func (fw flushWriter) Write(p []byte) (n int, err error) {
// 	n, err = fw.w.Write(p)
// 	// Flush - send the buffered written data to the client
// 	if f, ok := fw.w.(http.Flusher); ok {
// 		f.Flush()
// 	}
// 	return
// }

func handle(w http.ResponseWriter, r *http.Request) {
	//Allow CORS here By * or specific origin
	w.Header().Set("Access-Control-Allow-Origin", "*")

	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	log.Printf("Got connection: %s", r.Proto)

	w.Write([]byte("Hello"))
	// if f, ok := w.(http.Flusher); ok {
	// 	f.Flush()
	// }

}
