package model

import "time"

type Transfer struct {
	ID        int64     `json:"id"`
	FromId    int64     `json:"from"`
	ToId      int64     `json:"to"`
	Amount    float64   `json:"amount"`
	CreatedAt time.Time `json:"created_at"`
}
