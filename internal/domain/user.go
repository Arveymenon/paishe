package domain

import "gorm.io/gorm"

type User struct {
	gorm.Model        // ID, CreatedAt, UpdatedAt, and DeletedAt (Needs to added if gorm is being removed)
	Name       string `json:"name"`
	Email      string `json:"email"`
	Phone      string `json:"phone"`
}
