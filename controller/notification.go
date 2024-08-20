package controller

import (
	"sync"

	"github.com/gofiber/contrib/websocket"
)

var (
	clients   = make(map[*websocket.Conn]bool)
	broadcast = make(chan Message)
	mu        sync.Mutex
)

type Message struct {
	Type    string `json:"type"`    // "chat" ou "notification"
	UserID  string `json:"user_id"` // ID de l'utilisateur (pour les messages de chat)
	Content string `json:"content"` // Contenu du message ou de la notification
	Price   int    `json:"price"`   // Prix proposé pour les notifications
}

func HandleConnections(c *websocket.Conn) {
	defer func() {
		mu.Lock()
		delete(clients, c)
		mu.Unlock()
		c.Close()
	}()

	mu.Lock()
	clients[c] = true
	mu.Unlock()

	for {
		var msg Message
		err := c.ReadJSON(&msg)
		if err != nil {
			break
		}
		broadcast <- msg
	}
}

func HandleMessages() {
	for {
		msg := <-broadcast
		mu.Lock()
		for client := range clients {
			err := client.WriteJSON(msg)
			if err != nil {
				client.Close()
				delete(clients, client)
			}
		}
		mu.Unlock()
	}
}

/*
func messageNotification(ctx *fiber.Ctx, userNotified uint) error {
	//la fonction doit prevenir les participant d'un tout nouveau message

}
*/
