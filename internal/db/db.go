package db

import (
	"context"
	"database/sql"
	"time"
)

func New(addr, driver string, maxOpenCons, maxIdleCons int, maxIdleTime string) (*sql.DB, error) {
	db, err := sql.Open(driver, addr)
	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(maxOpenCons)
	db.SetMaxIdleConns(maxIdleCons)

	duration, err := time.ParseDuration(maxIdleTime)
	if err != nil {
		return nil, err
	}
	db.SetConnMaxIdleTime(duration)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := db.PingContext(ctx); err != nil {
		return nil, err
	}

	return db, nil
}
