package web

import (
	"time"
)

type Pet struct {
	ID       int `json:"id" example:"1"`
	Category struct {
		ID   int    `json:"id" example:"1"`
		Name string `json:"name" example:"category_name"`
	} `json:"category"`
	Name      string   `json:"name" example:"poti"`
	PhotoUrls []string `json:"photo_urls"`
	Tags      []Tag    `json:"tags"`
	Status    string   `json:"status"`
	Price     float32  `json:"price" example:"3.25"`
	IsAlive   bool     `json:"is_alive" example:"true"`
}

type Tag struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Pet2 struct {
	ID int `json:"id"`
}

type APIError struct {
	ErrorCode    int
	ErrorMessage string
	CreatedAt    time.Time
}

type RevValueBase struct {
	Status bool `json:"Status"`

	Err int32 `json:"Err"`
}
type RevValue struct {
	RevValueBase

	Data int `json:"Data"`
}
