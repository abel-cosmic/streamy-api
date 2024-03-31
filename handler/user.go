package handler

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// UserHandler type
type UserHandler struct {
	DB *gorm.DB
}

// Index to list all users
func (h UserHandler) Index(ctx *fiber.Ctx) error {
	return ctx.JSON(fiber.Map{"message": "This is a placeholder for listing all user"})
}

// Show an user
func (h UserHandler) Show(ctx *fiber.Ctx) error {
	return ctx.JSON(fiber.Map{"message": "This is a placeholder for showing an user"})
}

// Store a new user
func (h UserHandler) Store(ctx *fiber.Ctx) error {
	return ctx.JSON(fiber.Map{"message": "This is a placeholder for storing a new user"})
}

// Update an user
func (h UserHandler) Update(ctx *fiber.Ctx) error {
	return ctx.JSON(fiber.Map{"message": "This is a placeholder for updating an user"})
}

// Destroy an user
func (h UserHandler) Destroy(ctx *fiber.Ctx) error {
	return ctx.JSON(fiber.Map{"message": "This is a placeholder for destroying an user"})
}
