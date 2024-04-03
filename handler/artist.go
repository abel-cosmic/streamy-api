package handler

import (
	"errors"
	"strconv"

	"github.com/abel-cosmic/streamy-api/helper"
	"github.com/abel-cosmic/streamy-api/model"
	"github.com/abel-cosmic/streamy-api/utils"
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
	result := h.DB.Find(&artists)
	if result.Error != nil {
		return helper.HandleError(ctx, fiber.StatusInternalServerError, result.Error)
	}
	response := utils.TransformArtistsToArtistResponses(artists)
	return ctx.JSON(fiber.Map{"message": response})
}

// Show to list a single artist
func (h ArtistHandler) Show(ctx *fiber.Ctx) error {
	id, err := strconv.ParseUint(ctx.Params("id"), 10, 64)
	if err != nil {
		return helper.HandleError(ctx, fiber.StatusBadRequest, errors.New("invalid artist id"))
	}

	var artist model.Artist
	if err := h.DB.First(&artist, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return helper.HandleError(ctx, fiber.StatusNotFound, errors.New("artist not found"))
		}
		return helper.HandleError(ctx, fiber.StatusInternalServerError, err)
	}

	response := utils.TransformArtistToArtistResponse(artist)

	return ctx.JSON(fiber.Map{"artist": response})
}

// Store a new artist
func (h ArtistHandler) Store(ctx *fiber.Ctx) error {
	return ctx.JSON(fiber.Map{
		"message": "Artist Created Successfully",
	})
}

// Update an artist
func (h ArtistHandler) Update(ctx *fiber.Ctx) error {
	return ctx.JSON(fiber.Map{"message": "Artist Updated Successfully"})
}

// Destroy an artist
func (h ArtistHandler) Destroy(ctx *fiber.Ctx) error {
	return ctx.JSON(fiber.Map{"message": "This is a placeholder for destroying an artist"})
}
