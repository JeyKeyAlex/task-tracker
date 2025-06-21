package main

import (
	"github.com/JKasus/go_final_project/pkg/api"
	"github.com/JKasus/go_final_project/pkg/db"
	"log"
	"net/http"

	"github.com/JKasus/go_final_project/pkg/config"
)

func main() {
	cfg := config.NewConfig()
	err := db.Init(cfg.DbFile)
	if err != nil {
		log.Printf("Failed to init db: %v", err)
	}
	api.Init()
	http.Handle("/", http.FileServer(http.Dir(cfg.WebDir)))
	if err = http.ListenAndServe(cfg.Port, nil); err != nil {
		log.Printf("Failed to start server: %v", err)
	}
}
