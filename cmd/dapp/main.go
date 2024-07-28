package main

import (
	"context"
	"log/slog"

	"github.com/rollmelette/rollmelette"
)

func main() {
	//////////////////////// Setup Application //////////////////////////
	app := NewApp()

	///////////////////////// Rollmelette //////////////////////////
	ctx := context.Background()
	opts := rollmelette.NewRunOpts()
	err := rollmelette.Run(ctx, opts, app)
	if err != nil {
		slog.Error("application error", "error", err)
	}
}
