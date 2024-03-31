package validation

import (
	"github.com/abel-cosmic/streamy-api/model"
)

func CreateArtist(artist model.Artist) []string {
	var messages []string
	if !isValid(artist.Name) {
		messages = append(messages, "Name is required")
	}
	if !isValid(artist.Image) {
		messages = append(messages, "Image is required")
	}
	return messages
}

func UpdateArtist(artist model.Artist) []string {
	var messages []string
	if !isValid(artist.Name) {
		messages = append(messages, "Name is required")
	}
	if !isValid(artist.Image) {
		messages = append(messages, "Image is required")
	}
	return messages
}
