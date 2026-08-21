package app

import (
	"context"
	"database/sql"
	"fmt"
	"log/slog"
	"net"
	"net/http"
	"strconv"

	"github.com/VladiTNT/calorie-tracker-api/internal/config"
	"github.com/VladiTNT/calorie-tracker-api/internal/database"
	"github.com/VladiTNT/calorie-tracker-api/internal/middleware"
	"github.com/VladiTNT/calorie-tracker-api/pkg/auth"
)

type CalorieTrackerAPI struct {
	Config *config.Config
	Logger *slog.Logger

	Database    *sql.DB
	AuthService *auth.Service

	Router *http.ServeMux
	Server *http.Server
}

func (ctapi *CalorieTrackerAPI) Init(cfg *config.Config) error {
	// Config
	ctapi.Config = cfg

	// Logger
	ctapi.Logger = slog.New(slog.NewTextHandler(cfg.LoggerWriter, cfg.LoggerOpts))

	// Database
	db, err := database.New(cfg.DatabaseURL)
	if err != nil {
		return err
	}
	ctapi.Database = db

	// Auth Service
	ctapi.AuthService = auth.New()

	// Router
	ctapi.Router = http.NewServeMux()
	ctapi.RegisterRoutes()

	// Server
	ctapi.Server = &http.Server{
		Addr: net.JoinHostPort(cfg.Host, strconv.Itoa(cfg.Port)),
		Handler: middleware.Stack(
			middleware.Logger,
			ctapi.AuthService.Middleware,
		)(ctapi.Router),
	}

	return nil
}

func (ctapi *CalorieTrackerAPI) Run(ctx context.Context) error {
	go func() {
		fmt.Println("HTTP Server running on", ctapi.Server.Addr)
		if err := ctapi.Server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			panic(err)
		}
	}()

	<-ctx.Done()
	fmt.Println("Context canceled, shutting down...")

	shutdownCtx, cancel := context.WithTimeout(context.Background(), ctapi.Config.ShutdownTime)
	defer cancel()

	if err := ctapi.Server.Shutdown(shutdownCtx); err != nil {
		return err
	}

	fmt.Println("Server shutdown successfully!")

	return nil
}
