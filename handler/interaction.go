package handler

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// InteractionHandler type
type InteractionHandler struct {
	DB *gorm.DB
}

// Index to list all interactions
func (h InteractionHandler) Index(ctx *fiber.Ctx) error {
	return ctx.JSON(fiber.Map{"message": "This is a placeholder for listing all interaction"})
}

// Show an interaction
func (h InteractionHandler) Show(ctx *fiber.Ctx) error {
	return ctx.JSON(fiber.Map{"message": "This is a placeholder for showing an interaction"})
}

// Store a new interaction
func (h InteractionHandler) Store(ctx *fiber.Ctx) error {
	return ctx.JSON(fiber.Map{"message": "This is a placeholder for storing a new interaction"})
}

// Update an interaction
func (h InteractionHandler) Update(ctx *fiber.Ctx) error {
	return ctx.JSON(fiber.Map{"message": "This is a placeholder for updating an interaction"})
}

// Destroy an interaction
func (h InteractionHandler) Destroy(ctx *fiber.Ctx) error {
	return ctx.JSON(fiber.Map{"message": "This is a placeholder for destroying an interaction"})
}
