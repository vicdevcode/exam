package main

import (
	"flag"

	"github.com/vicdevcode/exam/internal/app"
	"github.com/vicdevcode/exam/internal/app/config"
	"github.com/vicdevcode/exam/internal/sqlite"
)

func main() {
	cfg := config.MustLoad()
	db, err := sqlite.New(cfg.DatabasePath)
	if err != nil {
		panic(err)
	}
	runType := flag.String("run", "", "")

	flag.Parse()

	app.Migrate(*runType, db)
}
