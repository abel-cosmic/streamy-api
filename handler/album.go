package handler

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// AlbumHandler type
type AlbumHandler struct {
	DB *gorm.DB
}

// Index to list all albums
func (h AlbumHandler) Index(ctx *fiber.Ctx) error {
	return ctx.JSON(fiber.Map{"message": "This is a placeholder for listing all albums"})
}

// Show an album
func (h AlbumHandler) Show(ctx *fiber.Ctx) error {
	return ctx.JSON(fiber.Map{"message": "This is a placeholder for showing an album"})
}

// Store a new album
func (h AlbumHandler) Store(ctx *fiber.Ctx) error {
	return ctx.JSON(fiber.Map{"message": "This is a placeholder for storing a new album"})
}

// Update an album
func (h AlbumHandler) Update(ctx *fiber.Ctx) error {
	return ctx.JSON(fiber.Map{"message": "This is a placeholder for updating an album"})
}

// Destroy an album
func (h AlbumHandler) Destroy(ctx *fiber.Ctx) error {
	return ctx.JSON(fiber.Map{"message": "This is a placeholder for destroying an album"})
}
