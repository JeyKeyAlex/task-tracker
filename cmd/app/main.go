package main

import (
	"fmt"
	"net/http"

	"github.com/JKasus/go_final_project/pkg/config"
)

func main() {
	cfg := config.NewConfig()
	http.Handle("/", http.FileServer(http.Dir(cfg.WebDir)))
	if err := http.ListenAndServe(cfg.Port, nil); err != nil {
		fmt.Println(err)
	}
}
