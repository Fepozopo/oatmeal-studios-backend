package service

import (
	"time"
)

type CreateOrUpdateProductInput struct {
	Type           string    `json:"type"`
	Sku            string    `json:"sku"`
	Upc            string    `json:"upc"`
	Status         string    `json:"status"`
	Cost           float64   `json:"cost"`
	Price          float64   `json:"price"`
	Envelope       string    `json:"envelope,omitempty"`
	Artist         string    `json:"artist,omitempty"`
	Category       string    `json:"category,omitempty"`
	ReleaseDate    time.Time `json:"release_date,omitempty"`
	LastBoughtDate time.Time `json:"last_bought_date,omitempty"`
	Description    string    `json:"description,omitempty"`
	TextFront      string    `json:"text_front,omitempty"`
	TextInside     string    `json:"text_inside,omitempty"`
}
