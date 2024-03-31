package handler

import (
	"errors"
	"strconv"

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
	id, err := strconv.ParseUint(ctx.Params("id"), 10, 64)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid artist ID"})
	}

	var artist model.Artist
	if err := h.DB.First(&artist, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Artist not found"})
		}
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Internal Server Error"})
	}

	return ctx.JSON(fiber.Map{"artist": artist})
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
