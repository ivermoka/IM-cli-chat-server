package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
    CheckOrigin: func(r *http.Request) bool {
        return true
    },
}

var clients = make(map[*websocket.Conn]bool)
var broadcast = make(chan Message)

type Message struct {
    Date  string `json:"date"`
    Message  string `json:"message"`
}

func main() {
    port := ":8080"
    http.HandleFunc("/", handleConnections)

    go handleMessages()

    fmt.Printf("Server started on %s", port)
    err := http.ListenAndServe(port, nil)
    if err != nil {
        panic("Error starting server: " + err.Error())
    }

}

func handleConnections(w http.ResponseWriter, r *http.Request) {
    conn, err := upgrader.Upgrade(w, r, nil)
    if err != nil {
        fmt.Println(err)
        return
    }
    defer conn.Close()

    clients[conn] = true
    fmt.Printf("\nActive connections: %d", activeConnections())

    for {
        var msg Message
        err := conn.ReadJSON(&msg)
        if err != nil {
            fmt.Println(err)
            delete(clients, conn)
            return
        }

        broadcast <- msg
    }
}

func handleMessages() {
    for {
        msg := <-broadcast

        for client := range clients {
            err := client.WriteJSON(msg)
            if err != nil {
                fmt.Println(err)
                client.Close()
                delete(clients, client)
            }
        }
    }
}

// ikke nødvendig, men morro å se tenker jeg :p
func activeConnections() int {
    count := 0
    for _, isConnected := range clients {
        if isConnected {
            count++
        }
    }
    return count
}

