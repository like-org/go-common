package model

import "time"

type Member struct {
	ID        int
	Username  string
	Mobile    string
	Email     string
	Password  string
	Status    string
	CreatedAt time.Time
	Avatar    string
	Signature string
	Birthday  time.Time
	Area      string
}
