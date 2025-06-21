package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"

	"github.com/JKasus/go_final_project/pkg/api"
	"github.com/JKasus/go_final_project/pkg/config"
	"github.com/JKasus/go_final_project/pkg/db"
)

func main() {
	cfg := config.NewConfig()
	err := db.Init(cfg.DbFile)
	if err != nil {
		log.Printf("Failed to init db: %v", err)
	}
	defer db.Close()
	r := chi.NewRouter()
	api.Init(r)
	r.Handle("/*", http.FileServer(http.Dir(cfg.WebDir)))
	if err = http.ListenAndServe(cfg.Port, r); err != nil {
		log.Printf("Failed to start server: %v", err)
	}
}
