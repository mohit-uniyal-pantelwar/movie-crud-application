package models

type Movie struct {
	Id              int     `json:"id"`
	Name            string  `json:"name"`
	Genre           string  `json:"genre"`
	Rating          float32 `json:"rating"`
	LengthInMinutes int     `json:"lengthInMinutes"`
	Language        string  `json:"language"`
}
