package handler

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// PlaylistSongHandler type
type PlaylistSongHandler struct {
	DB *gorm.DB
}

// Index to list all playlistsongs
func (h PlaylistSongHandler) Index(ctx *fiber.Ctx) error {
	return ctx.JSON(fiber.Map{"message": "This is a placeholder for listing all playlistsong"})
}

// Show an playlistsong
func (h PlaylistSongHandler) Show(ctx *fiber.Ctx) error {
	return ctx.JSON(fiber.Map{"message": "This is a placeholder for showing an playlistsong"})
}

// Store a new playlistsong
func (h PlaylistSongHandler) Store(ctx *fiber.Ctx) error {
	return ctx.JSON(fiber.Map{"message": "This is a placeholder for storing a new playlistsong"})
}

// Update an playlistsong
func (h PlaylistSongHandler) Update(ctx *fiber.Ctx) error {
	return ctx.JSON(fiber.Map{"message": "This is a placeholder for updating an playlistsong"})
}

// Destroy an playlistsong
func (h PlaylistSongHandler) Destroy(ctx *fiber.Ctx) error {
	return ctx.JSON(fiber.Map{"message": "This is a placeholder for destroying an playlistsong"})
}
