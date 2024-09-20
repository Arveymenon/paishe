package domain

import "gorm.io/gorm"

type User struct {
	gorm.Model // ID, CreatedAt, UpdatedAt, and DeletedAt (Needs to added if gorm is being removed)
	// ID         uint   `json:"id"`
	Name  string `json:"name"`
	Phone string `json:"phone"`
	Email string `json:"email"`
}
