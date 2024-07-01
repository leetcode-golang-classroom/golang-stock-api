package db

import (
	"context"
	"database/sql"
	"time"
)

func Connect(url string) (*sql.DB, error) {
	db, err := sql.Open("postgres", url)
	if err != nil {
		return nil, err
	}
	// setup limit
	db.SetConnMaxIdleTime(0)
	db.SetMaxIdleConns(3)
	db.SetMaxIdleConns(3)
	// test ping
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err = db.PingContext(ctx); err != nil {
		return nil, err
	}
	return db, nil
}
