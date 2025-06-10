package main

import (
	"fmt"
	"net/http"

	data "social-network/db/sqlite"
	"social-network/handler"
)

func main() {

	Db := data.GetDB()
	defer Db.Close()

	router := http.NewServeMux()

	// router.HandleFunc("/", handler.First)
	router.HandleFunc("/resgester", handler.Register)
	router.HandleFunc("/login", handler.Login)

	router.Handle("/statuts", handler.AuthMiddleware(http.HandlerFunc(handler.Statuts)))
	router.Handle("/pubpost", handler.AuthMiddleware(http.HandlerFunc(handler.Post)))
	router.Handle("/getpost", handler.AuthMiddleware(http.HandlerFunc(handler.Getpost)))
	// router.HandleFunc("/getpost", handler.Getpost)
	router.Handle("/getChats", handler.AuthMiddleware(http.HandlerFunc(handler.Getchats)))
	router.Handle("/sendcomment", handler.AuthMiddleware(http.HandlerFunc(handler.Sendcomment)))
	router.Handle("/getcomment", handler.AuthMiddleware(http.HandlerFunc(handler.Comments)))
	router.Handle("/logout", handler.AuthMiddleware(http.HandlerFunc(handler.Logout)))

	//follow system 
	router.Handle("/followRequest", handler.AuthMiddleware(http.HandlerFunc(handler.Followreq)))
	router.Handle("/unfollowRequest", handler.AuthMiddleware(http.HandlerFunc(handler.Unfollowreq)))

	// router.HandleFunc("/online-users", handler.OnlineUsers)
	router.HandleFunc("/ws", handler.WebSocketHandler) // Add WebSocket route
	router.Handle("/static/", http.HandlerFunc(handler.Sta))
	router.Handle("/javascript/", http.HandlerFunc(handler.Sta))
	go handler.HandleMessages() // Start WebSocket message handler in a goroutine
	go handler.Typing()

	// Wrap router with CORS middleware
	corsRouter := handler.CorsMiddleware(router)

	fmt.Println("✅ Server running on: http://localhost:8080")
	err := http.ListenAndServe(":8080", corsRouter)
	if err != nil {
		fmt.Println(err)
		return
	}
}
