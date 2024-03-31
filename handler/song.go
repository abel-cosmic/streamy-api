package handler

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// SongHandler type
type SongHandler struct {
	DB *gorm.DB
}

// Index to list all songs
func (h SongHandler) Index(ctx *fiber.Ctx) error {
	return ctx.JSON(fiber.Map{"message": "This is a placeholder for listing all song"})
}

// Show an song
func (h SongHandler) Show(ctx *fiber.Ctx) error {
	return ctx.JSON(fiber.Map{"message": "This is a placeholder for showing an song"})
}

// Store a new song
func (h SongHandler) Store(ctx *fiber.Ctx) error {
	return ctx.JSON(fiber.Map{"message": "This is a placeholder for storing a new song"})
}

// Update an song
func (h SongHandler) Update(ctx *fiber.Ctx) error {
	return ctx.JSON(fiber.Map{"message": "This is a placeholder for updating an song"})
}

// Destroy an song
func (h SongHandler) Destroy(ctx *fiber.Ctx) error {
	return ctx.JSON(fiber.Map{"message": "This is a placeholder for destroying an song"})
}
