package main

import (
	"fmt"
	"godoku/game"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		fmt.Println(r)
		return true // Allow all origins for this example
	},
}

func handleConnection(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)

	if err != nil {
		fmt.Println("Error while upgrading connection:", err)
		return
	}
	defer conn.Close()
	handle := game.NewWsHandle(conn)
	fmt.Println("Client connected")

	var g = game.GameFromFile("s1.txt")
	g.WsHandle = handle
	g.PrintGame()
	fmt.Println(g.IsFinished())

	handle.SendToWs(g)

	messageType, data, err := conn.ReadMessage()
	if err != nil {
		fmt.Println(err)
		return
	}
	if messageType != websocket.BinaryMessage {
		fmt.Println("Not a binary message", messageType)
		return
	}
	if handle.ParseMessage(data) {

		fmt.Println(g.IsFinished())
		g.Solve(conn)
		fmt.Println(g.IsFinished())
	}

}

func main() {
	http.HandleFunc("/ws", handleConnection)
	fmt.Println("Server started on :8123")
	err := http.ListenAndServe(":8123", nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}

}
