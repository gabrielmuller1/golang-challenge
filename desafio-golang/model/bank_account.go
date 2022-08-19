package model

import "time"

type BankAccount struct {
	ID        int64     `json:"id"`
	Number    string    `json:"account_number"`
	CreatedAt time.Time `json:"created_at"`
}
