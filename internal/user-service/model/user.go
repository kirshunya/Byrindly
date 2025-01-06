package model

import "time"

type User struct {
	ID               int64     `json:"id"`
	Name             string    `json:"name"`
	Age              uint8     `json:"age"`
	Gender           uint8     `json:"gender"`
	Latitude         float64   `json:"latitude"`
	Longitude        float64   `json:"longitude"`
	RegistrationDate time.Time `json:"registration-date"`
	About            string    `json:"about"`

	Username string `json:"username"`
	Photo    string `json:"photo"`
	TgId     int64  `json:"tg-id"`
	Password string `json:"password"`
}
