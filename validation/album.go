package validation

import (
	"github.com/abel-cosmic/streamy-api/model"
)

// CreateOrUpdate to verify the required fields
func CreateOrUpdate(album model.Album) []string {
	var messages []string

	// if !isValid(album.Title) {
	// 	messages = append(messages, "Title is required")
	// }

	// if !isValid(album.Total) {
	// 	messages = append(messages, "Total should be greater than zero")
	// }

	return messages
}
