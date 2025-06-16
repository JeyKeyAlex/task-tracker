package main

import (
	"fmt"
	"github.com/JKasus/go_final_project/pkg/db"
	"net/http"

	"github.com/JKasus/go_final_project/pkg/config"
)

func main() {
	cfg := config.NewConfig()
	db.Init(cfg.DbFile)
	http.Handle("/", http.FileServer(http.Dir(cfg.WebDir)))
	if err := http.ListenAndServe(cfg.Port, nil); err != nil {
		fmt.Println(err)
	}
}
