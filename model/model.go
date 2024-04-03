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
	DeletedAt time.Time `json:"deleted_at"`
}

// interaction type
type Interaction struct {
	gorm.Model
}

// password reset type
type PasswordReset struct {
	gorm.Model
}

// album type
type Album struct {
	gorm.Model
}

// playlist songs type
type PlaylistSong struct {
	gorm.Model
}

// playlist type
type Playlist struct {
	gorm.Model
}

// song type
type Song struct {
	gorm.Model
}

// user type
type User struct {
	gorm.Model
}
