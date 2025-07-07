package models

type Actor struct {
	ID            uint    `json:"id"`
	FirstName     string  `json:"first_name"`
	LastName      string  `json:"last_name"`
	NationalityID uint    `json:"nationality_id"`
	FilmCount     int     `json:"film_count"`
	TotalEarnings float64 `json:"total_earnings"`
	Nationality   string  `json:"nationality"`
}
