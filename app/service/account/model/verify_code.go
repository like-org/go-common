package model

import "time"

type VerifyCode struct {
	ID        int
	Key       string
	Code      int
	ExpiredAt time.Time
	CreatedAt time.Time
}
