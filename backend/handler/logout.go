package handler

import (
	"fmt"
	"net/http"

	db "social-network/Database/cration"
)

func Logout(w http.ResponseWriter, r *http.Request) {
	tocken, _ := r.Cookie("SessionToken")
	usernameFromToken := db.GetUsernameByToken(tocken.Value)

	_ = db.UpdateTocken(tocken.Value)

	// Safely remove the WebSocket connection from the clients map
	clientsMutex.Lock()

	for conn, username := range clients {
		fmt.Println("usernaem:", username, "usernameFromToken:", usernameFromToken)
		if username == usernameFromToken {
			conn.Close() // Close the WebSocket connection
			delete(clients, conn)
			fmt.Println("3", clients)
			break
		}
	}

	clientsMutex.Unlock()

	BroadcastUsers()

	// Respond with a success message
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"error": "Logout successful", "status":true}`))
}
