package handler

import (
	"github.com/abel-cosmic/streamy-api/model"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// ArtistHandler type
type ArtistHandler struct {
	DB *gorm.DB
}

// Index to list all artists
func (h ArtistHandler) Index(ctx *fiber.Ctx) error {
	var artists []model.Artist
	h.DB.Find(&artists)
	return ctx.JSON(fiber.Map{"message": artists})
}

// Show an artist
func (h ArtistHandler) Show(ctx *fiber.Ctx) error {
	return ctx.JSON(fiber.Map{"message": "This is a placeholder for showing an artist"})
}

// Store a new artist
func (h ArtistHandler) Store(ctx *fiber.Ctx) error {
	return ctx.JSON(fiber.Map{"message": "This is a placeholder for storing a new artist"})
}

// Update an artist
func (h ArtistHandler) Update(ctx *fiber.Ctx) error {
	return ctx.JSON(fiber.Map{"message": "This is a placeholder for updating an artist"})
}

// Destroy an artist
func (h ArtistHandler) Destroy(ctx *fiber.Ctx) error {
	return ctx.JSON(fiber.Map{"message": "This is a placeholder for destroying an artist"})
}
