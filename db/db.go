package db

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

var DB *sql.DB

// инициализация бд
func InitDB() (*sql.DB, error) {
	connStr := "host=localhost user=admin password=admin dbname=movieactors sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Ошибка при подключении к базе данных: %v", err)
		return nil, err
	}
	if err := db.Ping(); err != nil {
		log.Fatalf("Ошибка при подключении к базе данных: %v", err)
		return nil, err
	}
	log.Println("Подключение к базе данных установлено.")
	DB = db
	return db, nil
}
