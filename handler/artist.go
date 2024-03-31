package handler

import (
	"errors"
	"net/http"
	"strconv"
	"time"

	"path/filepath"
	"strings"

	"github.com/abel-cosmic/streamy-api/model"
	"github.com/abel-cosmic/streamy-api/validation"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// ArtistHandler type
type ArtistHandler struct {
	DB *gorm.DB
}

type ArtistResponse struct {
	ID        uint      `json:"id"`
	Name      string    `json:"name"`
	Image     string    `json:"image"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// Index to list all artists
func (h ArtistHandler) Index(ctx *fiber.Ctx) error {
	var artists []model.Artist
	result := h.DB.Find(&artists)
	if result.Error != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Internal server error",
		})
	}
	var response []ArtistResponse
	for _, artist := range artists {
		response = append(response, ArtistResponse{
			ID:        artist.ID,
			Name:      artist.Name,
			Image:     artist.Image,
			CreatedAt: artist.CreatedAt,
			UpdatedAt: artist.UpdatedAt,
		})
	}

	return ctx.JSON(fiber.Map{"message": response})
}

// Show to list a single artist
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
	response := ArtistResponse{
		ID:    artist.ID,
		Name:  artist.Name,
		Image: artist.Image,
	}

	return ctx.JSON(fiber.Map{"artist": response})
}

// Store a new artist
func (h ArtistHandler) Store(ctx *fiber.Ctx) error {
	artist := new(model.Artist)
	artist.Name = ctx.FormValue("name")

	file, err := ctx.FormFile("image")
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Problem with image upload"})
	}
	filename := filepath.Base(file.Filename)
	filename = strings.ReplaceAll(filename, " ", "-")
	imagePath := filepath.Join("./public", filename)
	if err := ctx.SaveFile(file, imagePath); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Could not save file"})
	}
	artist.Image = imagePath
	if validate := validation.CreateArtist(*artist); len(validate) > 0 {
		return ctx.Status(422).JSON(fiber.Map{"errors": validate})
	}
	if result := h.DB.Create(&artist); result.Error != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to save artist in database", "details": result.Error.Error()})
	}
	response := ArtistResponse{
		ID:    artist.ID,
		Name:  artist.Name,
		Image: artist.Image,
	}
	return ctx.JSON(fiber.Map{
		"message": "Artist Created Successfully",
		"res":     response,
	})
}

// Update an artist
func (h ArtistHandler) Update(ctx *fiber.Ctx) error {
	// Parse artist ID from URL parameter
	id, err := strconv.ParseUint(ctx.Params("id"), 10, 64)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid artist ID"})
	}

	// Retrieve the existing artist from the database
	var existingArtist model.Artist
	if err := h.DB.First(&existingArtist, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Artist not found"})
		}
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Internal Server Error"})
	}

	// Update artist fields based on request data
	artist := &model.Artist{
		Name: ctx.FormValue("name"),
	}

	// Handle image upload if provided
	file, err := ctx.FormFile("image")
	if err != nil && err != http.ErrMissingFile {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Problem with image upload"})
	}
	if file != nil {
		filename := filepath.Base(file.Filename)
		filename = strings.ReplaceAll(filename, " ", "-")
		imagePath := filepath.Join("./public", filename)
		if err := ctx.SaveFile(file, imagePath); err != nil {
			return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Could not save file"})
		}
		artist.Image = imagePath
	}

	// Validate the updated artist
	if validate := validation.CreateArtist(*artist); len(validate) > 0 {
		return ctx.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{"errors": validate})
	}

	// Update the artist in the database
	if err := h.DB.Model(&existingArtist).Updates(artist).Error; err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to update artist in database", "details": err.Error()})
	}

	// Construct response
	response := ArtistResponse{
		ID:    existingArtist.ID,
		Name:  artist.Name,
		Image: artist.Image, // Use updated image if provided, otherwise keep the existing image
	}

	return ctx.JSON(fiber.Map{"message": "Artist Updated Successfully", "res": response})
}

// Destroy an artist
func (h ArtistHandler) Destroy(ctx *fiber.Ctx) error {
	return ctx.JSON(fiber.Map{"message": "This is a placeholder for destroying an artist"})
}
