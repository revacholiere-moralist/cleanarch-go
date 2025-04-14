package main

import (
	"log"

	"github.com/lpernett/godotenv"
	"github.com/revacholiere-moralist/cleanarch-go/internal/db"
	"github.com/revacholiere-moralist/cleanarch-go/internal/env"
	"github.com/revacholiere-moralist/cleanarch-go/internal/store"
)

func main() {
	godotenv.Load()

	cfg := config{
		addr: env.GetString("ADDR", ":8080"),

		db: dbConfig{
			addr:           env.GetString("DB_ADDR", "postgres://postgres:-@localhost:5432/cleanarch?sslmode=disable"),
			maxOpenConns:   env.GetInt("DB_MAX_OPEN_CONNS", 30),
			maxIdleConns:   env.GetInt("DB_MAX_IDLE_CONNS", 30),
			maxIdleTime:    env.GetInt("DB_MAX_IDLE_TIME", 900),
			maxTimeOutTime: env.GetInt("DB_MAX_TIMEOUT_TIME", 5),
		},
	}

	db, err := db.New(
		cfg.db.addr,
		cfg.db.maxOpenConns,
		cfg.db.maxIdleConns,
		cfg.db.maxIdleTime,
		cfg.db.maxTimeOutTime,
	)

	if err != nil {
		log.Panic(err)
	}

	defer db.Close()
	log.Println("database connection pool established")
	store := store.NewStorage(db)

	app := &application{
		config: cfg,
		store:  store,
	}

	mux := app.mount()
	log.Fatal(app.run(mux))
}
