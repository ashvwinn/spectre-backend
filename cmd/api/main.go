package main

import (
	"flag"
	"log/slog"
	"os"

	"github.com/gin-gonic/gin"
)

const version = "1.0.0"

type config struct {
	port int
	env  string
}

type application struct {
	config config
	logger *slog.Logger
}

func main() {
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	var cfg config

	flag.IntVar(&cfg.port, "port", 8080, "API server port")
	flag.StringVar(&cfg.env, "env", "development", "Environment (development|staging|production)")
	flag.Parse()

	app := &application{
		config: cfg,
		logger: logger,
	}

	g := gin.New()
	app.registerRoutes(g)

	logger.Info("Starting server", "version", version, "port", cfg.port, "env", cfg.env)
	if err := g.Run(); err != nil {
		logger.Error("Server Error: " + err.Error())
	}
}
