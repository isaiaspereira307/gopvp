package main

import (
	"log"

	"github.com/isaiaspereira307/gopvp/configs"
	"github.com/isaiaspereira307/gopvp/internal/db"
	"github.com/isaiaspereira307/gopvp/routes"
)

func main() {
	err := configs.Load()
	if err != nil {
		log.Fatal("Failed to load configurations:", err)
		panic(err)
	}
	conn, err := db.OpenConnection()
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	defer conn.Close()

	queries := db.New(conn)
	routes.Initialize(queries)
}
