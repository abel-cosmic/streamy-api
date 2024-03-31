package handler

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// PasswordResetHandler type
type PasswordResetHandler struct {
	DB *gorm.DB
}

// Index to list all passwordresets
func (h PasswordResetHandler) Index(ctx *fiber.Ctx) error {
	return ctx.JSON(fiber.Map{"message": "This is a placeholder for listing all passwordreset"})
}

// Show an passwordreset
func (h PasswordResetHandler) Show(ctx *fiber.Ctx) error {
	return ctx.JSON(fiber.Map{"message": "This is a placeholder for showing an passwordreset"})
}

// Store a new passwordreset
func (h PasswordResetHandler) Store(ctx *fiber.Ctx) error {
	return ctx.JSON(fiber.Map{"message": "This is a placeholder for storing a new passwordreset"})
}

// Update an passwordreset
func (h PasswordResetHandler) Update(ctx *fiber.Ctx) error {
	return ctx.JSON(fiber.Map{"message": "This is a placeholder for updating an passwordreset"})
}

// Destroy an passwordreset
func (h PasswordResetHandler) Destroy(ctx *fiber.Ctx) error {
	return ctx.JSON(fiber.Map{"message": "This is a placeholder for destroying an passwordreset"})
}
