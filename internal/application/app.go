package application

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/leetcode-golang-classroom/golang-stock-api/internal/config"
	"github.com/leetcode-golang-classroom/golang-stock-api/internal/db"
	"github.com/leetcode-golang-classroom/golang-stock-api/internal/util"
	_ "github.com/lib/pq"
)

type App struct {
	router *mux.Router
	config *config.Config
	db     *sql.DB
}

func New(config *config.Config) *App {
	dbConn, err := db.Connect(config.DbURL)
	if err != nil {
		util.FailOnError(err, "failed to connect")
	}
	router := NewRouter()
	app := &App{
		router: router,
		config: config,
		db:     dbConn,
	}
	// setup routes
	app.loadRoutes()
	return app
}

func (app *App) Start(ctx context.Context) error {
	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", app.config.Port),
		Handler: app.router,
	}
	log.Printf("Starting server on %d\n", app.config.Port)
	errCh := make(chan error, 1)
	go func() {
		err := server.ListenAndServe()
		if err != nil {
			errCh <- fmt.Errorf("failed to start server: %w", err)
		}
		util.CloseChannel(errCh)
	}()
	select {
	case err := <-errCh:
		return err
	case <-ctx.Done():
		timeout, cancel := context.WithTimeout(context.Background(), time.Second*10)
		defer cancel()
		return server.Shutdown(timeout)
	}
}
