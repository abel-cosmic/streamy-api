package handler

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// PlaylistHandler type
type PlaylistHandler struct {
	DB *gorm.DB
}

// Index to list all playlists
func (h PlaylistHandler) Index(ctx *fiber.Ctx) error {
	return ctx.JSON(fiber.Map{"message": "This is a placeholder for listing all playlist"})
}

// Show an playlist
func (h PlaylistHandler) Show(ctx *fiber.Ctx) error {
	return ctx.JSON(fiber.Map{"message": "This is a placeholder for showing an playlist"})
}

// Store a new playlist
func (h PlaylistHandler) Store(ctx *fiber.Ctx) error {
	return ctx.JSON(fiber.Map{"message": "This is a placeholder for storing a new playlist"})
}

// Update an playlist
func (h PlaylistHandler) Update(ctx *fiber.Ctx) error {
	return ctx.JSON(fiber.Map{"message": "This is a placeholder for updating an playlist"})
}

// Destroy an playlist
func (h PlaylistHandler) Destroy(ctx *fiber.Ctx) error {
	return ctx.JSON(fiber.Map{"message": "This is a placeholder for destroying an playlist"})
}
