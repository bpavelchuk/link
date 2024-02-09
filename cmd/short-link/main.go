package main

import (
	"fmt"
	"short-link/internal/config"
)

func main() {
	cfg := config.MustLoad()
	fmt.Println(cfg)
	// TODO: init config: cleanenv
	// TODO: inti logger: slog
	// TODO: init storage: sqlite
	// TODO: init router: chi, render
	// TODO: run server

}
