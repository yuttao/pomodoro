package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/tomyutao/pomodoro/pkg/config"
	"github.com/tomyutao/pomodoro/pkg/handlers"
)

func main() {
	cfg, err := config.ReadConfig()
	if err != nil {
		log.Println("Config file not found, use default port: 8080")
	} else {
		log.Println("Listening on port:", cfg.Port)
	}
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	_ = http.ListenAndServe(fmt.Sprintf(":%d", cfg.Port), nil)
}
