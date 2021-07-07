package model

import "time"

type Account struct {
	ID        int
	Amount    int
	Status    int
	Type      int
	UpdatedAt time.Time
	CreatedAt time.Time
}
