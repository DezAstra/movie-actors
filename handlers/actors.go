package handlers

import (
	"encoding/json"
	"log"
	"movieactors/db"
	"movieactors/models"
	"net/http"
)

func GetActors(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// параметры запроса
	search := r.URL.Query().Get("search")
	filmsMin := r.URL.Query().Get("films_min")
	filmsMax := r.URL.Query().Get("films_max")
	earningsMin := r.URL.Query().Get("earnings_min")
	earningsMax := r.URL.Query().Get("earnings_max")
	nationalityID := r.URL.Query().Get("nationality")

	// SQL-запрос
	query := "SELECT a.id, a.first_name, a.last_name, a.nationality_id, a.film_count, a.total_earnings, n.name AS nationality FROM actors a LEFT JOIN nationalities n ON a.nationality_id = n.id WHERE 1=1"

	// фильтрация по поисковому запросу
	if search != "" {
		query += " AND (a.first_name ILIKE '%" + search + "%' OR a.last_name ILIKE '%" + search + "%')"
	}

	// фильтрация по кол-ву фильмов
	if filmsMin != "" {
		query += " AND a.film_count >= " + filmsMin
	}
	if filmsMax != "" {
		query += " AND a.film_count <= " + filmsMax
	}

	// фильтрация по гонорарам
	if earningsMin != "" {
		query += " AND a.total_earnings >= " + earningsMin
	}
	if earningsMax != "" {
		query += " AND a.total_earnings <= " + earningsMax
	}

	// фильтрация по национальности
	if nationalityID != "" {
		query += " AND a.nationality_id = " + nationalityID
	}

	// Выполнение запроса
	rows, err := db.DB.Query(query)
	if err != nil {
		log.Println("Ошибка при выполнении запроса:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	// чтение результата
	var actors []models.Actor
	for rows.Next() {
		var actor models.Actor
		err := rows.Scan(&actor.ID, &actor.FirstName, &actor.LastName, &actor.NationalityID, &actor.FilmCount, &actor.TotalEarnings, &actor.Nationality)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		actors = append(actors, actor)
	}

	// проверка на наличие ошибок при итерации
	if err := rows.Err(); err != nil {
		log.Println("Ошибка при чтении данных:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(actors)
}
