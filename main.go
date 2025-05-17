package main

import (
	"fmt"
	"net/http"

	"social-network/backend/pkg/api/auth"
	"social-network/backend/pkg/api/groups"
	"social-network/backend/pkg/db/sqlite"
)

func main() {
	db := sqlite.OpenDB()
	defer db.Close()
	fmt.Println("Database opened successfully")
	router := http.NewServeMux()

	router.Handle("/api/", http.StripPrefix("/api", auth.AuthMux()))
	router.Handle("/api/groups/", http.StripPrefix("/api/groups", groups.GroupMux()))
	// log.Fatal(http.ListenAndServe(":8080", middleware.CheckCORS(router))
}
