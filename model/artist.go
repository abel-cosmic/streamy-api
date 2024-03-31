package model

import (
	"time"

	"gorm.io/gorm"
)

// artist type
type Artist struct {
	gorm.Model
	Id        uint      `json:"id" gorm:"primaryKey"`
	Name      string    `json:"name"`
	Image     string    `json:"image"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
