package dto

type MovieInput struct {
	Name     string `json:"name"`
	Year     string `json:"year"`
	Director string `json:"director"`
	Writer   string `json:"writer"`
}