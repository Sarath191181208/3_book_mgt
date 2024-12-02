package main

import (
	"context"
	"database/sql"
	"log"
	"net/http"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"

	"sarath/3_book_mgt/cmd/api"
	"sarath/3_book_mgt/internal/data"
	"sarath/3_book_mgt/internal/logger"
)

func main() {
	dsn := os.Getenv("DB_DSN")
	conn, err := connectDB(dsn)
	if err != nil {
		log.Fatal(err)
	}

	app := api.Application{
		Logger: logger.NewSysOutLogger(),
		Db:     data.New(conn),
	}

	server := http.Server{
		Addr:    ":8000",
		Handler: app.Routes(),
	}

	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}

func connectDB(dsn string) (*sql.DB, error) {
	// open the conn
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	// ping the db
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	err = db.PingContext(ctx)
	if err != nil {
		return nil, err
	}

  db.SetConnMaxLifetime(-1)
	db.SetMaxOpenConns(200) // Limit max open connections to the DB
	db.SetMaxIdleConns(0)  // Control idle connections

	return db, nil
}
