package main

import (
	"fmt"
	"social-network/backend/pkg/db/sqlite"
)

func main() {

	db := sqlite.OpenDB()
	defer db.Close()
	fmt.Println("Database opened successfully")

}
