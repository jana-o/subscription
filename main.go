package main

import (
	"database/sql"
	"fmt"
	"github.com/jana-o/subscription/app"
	"github.com/jana-o/subscription/config"
	"github.com/jana-o/subscription/db"
	_ "github.com/lib/pq"
	"log"
)

func main() {
	cfg, err := config.LoadConfig(".")
	if err != nil {
		log.Fatal("can't load config:", err)
	}

	conn, err := sql.Open(cfg.DBDriver, cfg.DBSource)
	if err != nil {
		log.Fatal("can't connect to db:", err)
	}
	if err := conn.Ping(); err != nil {
		log.Fatal("can't connect to db:", err)
	}

	store := db.NewStore(conn)
	fmt.Println("connected to db")

	server, err := app.NewServer(cfg, store)
	if err != nil {
		log.Fatal("can't create server:", err)
	}

	err = server.Start(cfg.ServerAddress)
	if err != nil {
		log.Fatal("can't start server:", err)
	}
}
