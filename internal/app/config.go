package app

import (
	"fmt"

	"github.com/paincake00/order-service/internal/env"
)

type Config struct {
	addr  string
	db    DBConfig
	cache CacheConfig
}

type DBConfig struct {
	addr        string
	driver      string
	maxOpenCons int
	maxIdleCons int
	maxIdleTime string
}

type CacheConfig struct {
	capacity int
}

func LoadConfig() Config {
	cfg := Config{
		addr: env.GetString("ADDR", ":8080"),
		db: DBConfig{
			addr:        getUriDB(),
			driver:      env.GetString("DB_DRIVER", "postgres"),
			maxOpenCons: env.GetInt("DB_MAX_OPEN_CONS", 30),
			maxIdleCons: env.GetInt("DB_MAX_IDLE_CONS", 30),
			maxIdleTime: env.GetString("DB_MAX_IDLE_TIME", "15m"),
		},
		cache: CacheConfig{
			capacity: env.GetInt("CACHE_MAX_CAPACITY", 10),
		},
	}

	return cfg
}

func getUriDB() string {
	schema := env.GetString("DB_DRIVER", "postgres")
	user := env.GetString("POSTGRES_USER", "postgres")
	password := env.GetString("POSTGRES_PASSWORD", "postgres")
	localhost := env.GetString("POSTGRES_HOST", "localhost")
	port := env.GetString("POSTGRES_PORT", "5432")
	dbName := env.GetString("POSTGRES_DB", "some-db-name")

	return fmt.Sprintf("%s://%s:%s@%s:%s/%s?sslmode=disable", schema, user, password, localhost, port, dbName)
}
