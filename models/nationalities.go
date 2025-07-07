package models

import (
	"database/sql"
	"log"
)

type Nationality struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

// Получение всех национальностей из базы данных
func GetAllNationalities(db *sql.DB) ([]Nationality, error) {
	var nationalities []Nationality

	// SQL-запрос
	rows, err := db.Query("SELECT id, name FROM nationalities")
	if err != nil {
		log.Println("Ошибка при запросе национальностей:", err)
		return nil, err
	}
	defer rows.Close()

	// чтение строк результата
	for rows.Next() {
		var nationality Nationality
		err := rows.Scan(&nationality.ID, &nationality.Name)
		if err != nil {
			log.Println("Ошибка при сканировании строки:", err)
			return nil, err
		}
		nationalities = append(nationalities, nationality)
	}

	// Проверка на наличие ошибок при итерации
	if err := rows.Err(); err != nil {
		log.Println("Ошибка при обработке строк:", err)
		return nil, err
	}

	return nationalities, nil
}
