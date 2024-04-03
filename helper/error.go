package helper

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func HandleError(ctx *fiber.Ctx, statusCode int, err error) error {
	// Define the error message based on the most common HTTP status codes.
	var errorMessage string
	switch statusCode {
	case fiber.StatusBadRequest: // 400
		errorMessage = "Bad request. Please check your request and try again."
	case fiber.StatusUnauthorized: // 401
		errorMessage = "Unauthorized. Please provide valid credentials."
	case fiber.StatusForbidden: // 403
		errorMessage = "Forbidden. You do not have permission to access this resource."
	case fiber.StatusNotFound: // 404
		errorMessage = "Resource not found. The requested resource could not be found."
	case fiber.StatusConflict: // 409
		errorMessage = "Conflict. The request could not be completed due to a conflict with the current state of the resource."
	default:
		// This handles any other codes and defaults to internal server error
		statusCode = fiber.StatusInternalServerError // Ensure the status code is 500
		errorMessage = "Internal server error. Please try again later."
		// Optional: Log the actual error internally for diagnostics
		fmt.Printf("Internal Server Error: %v\n", err)
	}

	// Respond with the determined status code and error message
	return ctx.Status(statusCode).JSON(fiber.Map{
		"error": errorMessage,
	})
}
