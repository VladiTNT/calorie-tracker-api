package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"

	"github.com/VladiTNT/calorie-tracker-api/internal/app"
	"github.com/VladiTNT/calorie-tracker-api/internal/config"
)

func main() {
	var a app.CalorieTrackerAPI
	if err := a.Init(config.Default()); err != nil {
		fmt.Fprintf(os.Stderr, "Error initializing api: %v\n", err)
		os.Exit(1)
	}

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	if err := a.Run(ctx); err != nil {
		fmt.Fprintf(os.Stderr, "Error shutting down api: %v\n", err)
		os.Exit(1)
	}
}
