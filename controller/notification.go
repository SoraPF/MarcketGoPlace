package controller

import (
	"Marcketplace/data/request"
	"sync"

	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
)

var (
	clients   = make(map[*websocket.Conn]bool)
	broadcast = make(chan string)
	mu        sync.Mutex
)

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
		_, msg, err := c.ReadMessage()
		if err != nil {
			break
		}
		broadcast <- string(msg)
	}
}

func HandleMessages() {
	for {
		msg := <-broadcast
		mu.Lock()
		for client := range clients {
			err := client.WriteMessage(websocket.TextMessage, []byte(msg))
			if err != nil {
				client.Close()
				delete(clients, client)
			}
		}
		mu.Unlock()
	}
}

func NotifiedAdminNewArticle(ctx *fiber.Ctx, req *request.CreateObjRequest) error {
	notification := "Un nouvel article a été créé : " + req.Title
	broadcast <- notification

	return ctx.SendStatus(fiber.StatusOK)
}

func NotifiedUserNewArticle(ctx *fiber.Ctx, req *request.CreateObjRequest) error {
	notification := "Article " + req.Title + " Valider par un Admin"
	broadcast <- notification

	return ctx.SendStatus(fiber.StatusOK)
}

/*
func messageNotification(ctx *fiber.Ctx, userNotified uint) error {
	//la fonction doit prevenir les participant d'un tout nouveau message

}
*/
