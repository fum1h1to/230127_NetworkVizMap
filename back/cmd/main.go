package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
    WriteBufferSize: 1024,
    // Add this lines
    CheckOrigin: func(r *http.Request) bool {
        return true
    },
}

func webSocketHandleFunc(w http.ResponseWriter, r *http.Request) {
    ws, err := upgrader.Upgrade(w, r, nil)
    if err != nil {
        log.Println("upgrade:", err)
        return
    }
    defer ws.Close()

    err = ws.WriteMessage(websocket.TextMessage, []byte(`Server (gorilla): Hello, Client!`))
    if err != nil {
        log.Println("WriteMessage:", err)
        return
    }

    for {
        mt, message, err := ws.ReadMessage()
        if err != nil {
            log.Println("ReadMessage:", err)
            break
        }
        err = ws.WriteMessage(mt, []byte(fmt.Sprintf("Server (gorilla): '%s' received.", message)))
        if err != nil {
            log.Println("WirteMessage:", err)
            break
        }
    }
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello from Closure")
	})
    http.HandleFunc("/ws", webSocketHandleFunc)

    if err := http.ListenAndServe(":8080", nil); err != nil {
        log.Fatal(err)
    }
}