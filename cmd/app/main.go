package main

import (
	"github.com/vicdevcode/exam/internal/app"
	"github.com/vicdevcode/exam/internal/app/config"
)

func main() {
	cfg := config.MustLoad()

	app.Run(cfg)
}
