package model

import (
	"time"
	"mime/multipart"
)


type User struct {
	ID         int64                 `json:"id"`
	Name        string                `json:"name"`
	Email       string                `json:"email"`
	Age        int64                 `json:"age"`
	Address     *string               `json:"address"`
	PhoneNumber *string               `json:"phone_number"`
	Gender      *string               `json:"gender"`
	Status      *string               `json:"status"`
	City        *string               `json:"city"`
	Province    *string               `json:"province"`
	CreatedAt   time.Time             `json:"created_at"`
	UpdatedAt   *time.Time            `json:"updated_at"`
	Image       *multipart.FileHeader `json:"image"`
}
