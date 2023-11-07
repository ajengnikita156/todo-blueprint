package model

import (
	"time"
)

type User struct {  
	ID        int
	Nama       string
	Email      string
	Umur       int
	CreatedAt  time.Time
	UpdatedAt   time.Time
}