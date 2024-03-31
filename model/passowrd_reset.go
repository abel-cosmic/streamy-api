package model

import (
	"gorm.io/gorm"
)

// password reset type
type PasswordReset struct {
	gorm.Model
}
