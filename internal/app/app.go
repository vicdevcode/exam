package app

import (
	"fmt"
	"io/fs"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/vicdevcode/exam/frontend"
	"github.com/vicdevcode/exam/internal/app/config"
	"github.com/vicdevcode/exam/internal/controller"
	"github.com/vicdevcode/exam/internal/sqlite"
	"github.com/vicdevcode/exam/internal/usecase"
)

func Run(cfg *config.Config) {
	db, err := sqlite.New(&cfg.Postgres)
	if err != nil {
		panic(err)
	}

	uc := usecase.New(cfg, db)

	gin.SetMode(gin.ReleaseMode)
	if cfg.Env == "local" {
		gin.SetMode(gin.DebugMode)
	} else {
		Migrate("create", db)
		// addData(uc)
	}

	handler := gin.Default()

	staticHandler(handler)

	controller.NewRouter(handler, cfg, uc)

	handler.Run(fmt.Sprintf("%s:%s", cfg.Host, cfg.Port))
}

func staticHandler(engine *gin.Engine) {
	dist, _ := fs.Sub(frontend.Dist, "dist")
	fileServer := http.FileServer(http.FS(dist))

	engine.Use(func(c *gin.Context) {
		if !strings.HasPrefix(c.Request.URL.Path, "/api") {
			_, err := fs.Stat(dist, strings.TrimPrefix(c.Request.URL.Path, "/"))
			if os.IsNotExist(err) {
				fmt.Println("File not found, serving index.html")
				c.Request.URL.Path = "/"
			} else {
				fmt.Println("Serving other static files")
			}

			fileServer.ServeHTTP(c.Writer, c.Request)
			c.Abort()
		}
	})
}
