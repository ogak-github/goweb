package app

import (
	"context"
	"fmt"
	"log"
	"sync"

	"github.com/jackc/pgx/v5/pgxpool"
)

var (
	instance *pgxpool.Pool
	once     sync.Once
)

func appDB(connString string) *pgxpool.Pool {
	pool, err := pgxpool.New(context.Background(), connString)
	if err != nil {
		log.Fatal("Cannot connect to database: ", err)
	}

	err = pool.Ping(context.Background())
	if err != nil {
		log.Fatal("Can't ping to DB:", err)
	}
	fmt.Println("Database ready!")
	return pool
}

func GetDatabaseConnection(connString string) *pgxpool.Pool {
	once.Do(func() {
		instance = appDB(connString)
	})
	return instance
}
