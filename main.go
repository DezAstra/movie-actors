package main

import (
	"fmt"
	"log"
	"movieactors/db"
	"movieactors/handlers"
	"net/http"
)

func main() {
	// Инициализация базы данных
	database, err := db.InitDB()
	if err != nil {
		log.Fatalf("Ошибка при подключении к базе данных: %v", err)
	}
	defer database.Close()

	fmt.Println("БД инициализирована.")

	// Настройка маршрутов
	http.HandleFunc("/actors", handlers.GetActors)
	http.HandleFunc("/nationalities", handlers.GetNationalities)

	// Подключение статических файлов
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// Запуск сервера
	log.Println("Сервер запущен.")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
