package main

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

type message struct {
	Username string `json:"username"`
	Text     string `json:"text"`
}

type room struct {
	ID          string
	Connections map[*websocket.Conn]bool
	Broadcast   chan message
	Messages    []message
}

var rooms = make(map[string]*room)

func handleWebSocket(w http.ResponseWriter, r *http.Request) {
	// Upgrade HTTP connection to WebSocket
	upgrader := websocket.Upgrader{}
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}

	// Get room ID from URL query parameter
	roomID := r.URL.Query().Get("room")

	// Create new room if it doesn't exist
	if _, ok := rooms[roomID]; !ok {
		rooms[roomID] = &room{
			ID:          roomID,
			Connections: make(map[*websocket.Conn]bool),
			Broadcast:   make(chan message),
			Messages:    make([]message, 0),
		}
		go rooms[roomID].run()
	}

	// Add connection to room
	rooms[roomID].Connections[conn] = true
}

func (r *room) run() {
	for {
		// Wait for incoming message on broadcast channel
		msg := <-r.Broadcast

		// Add message to room history
		r.Messages = append(r.Messages, msg)

		// Broadcast message to all connections in room
		for conn := range r.Connections {
			err := conn.WriteJSON(msg)
			if err != nil {
				log.Println(err)
				delete(r.Connections, conn)
			}
		}
	}
}

func handleChatMessage(w http.ResponseWriter, r *http.Request) {
	// Get room ID from URL query parameter
	roomID := r.URL.Query().Get("room")

	// Get chat room
	chatRoom, ok := rooms[roomID]
	if !ok {
		http.NotFound(w, r)
		return
	}

	// Parse incoming message from JSON
	var msg message
	err := json.NewDecoder(r.Body).Decode(&msg)
	if err != nil {
		log.Println(err)
		return
	}

	// Add message to room history
	chatRoom.Messages = append(chatRoom.Messages, msg)

	// Broadcast message to all connections in room
	chatRoom.Broadcast <- msg
}

func handleChatRoom(w http.ResponseWriter, r *http.Request) {
	// Get room ID from URL query parameter
	roomID := r.URL.Query().Get("room")

	// Get chat room
	chatRoom, ok := rooms[roomID]
	if !ok {
		http.NotFound(w, r)
		return
	}

	// Define template data
	type templateData struct {
		RoomID   string
		Messages []message
	}
	data := templateData{
		RoomID:   roomID,
		Messages: make([]message, 0),
	}

	// Collect chat messages
	for _, msg := range chatRoom.Messages {
		data.Messages = append(data.Messages, msg)
	}

	// Execute template
	tmpl, err := template.ParseFiles("chat.html")
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	err = tmpl.Execute(w, data)
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
}

func main() {
	// Register HTTP handlers
	http.HandleFunc("/ws", handleWebSocket)
	http.HandleFunc("/chat", handleChatRoom)
	http.HandleFunc("/message", handleChatMessage)

	// Start HTTP server
	log.Println("Starting server on :8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
