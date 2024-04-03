package utils

import (
	"github.com/abel-cosmic/streamy-api/model"
	"github.com/abel-cosmic/streamy-api/schema"
)

// TransformArtistsToArtistResponses takes a slice of Artist and transforms it to a slice of ArtistResponse.
func TransformArtistsToArtistResponses(artists []model.Artist) []schema.ArtistResponse {
	var responses []schema.ArtistResponse
	for _, artist := range artists {
		response := schema.ArtistResponse{
			ID:        artist.ID,
			Name:      artist.Name,
			Image:     artist.Image,
			CreatedAt: artist.CreatedAt,
			UpdatedAt: artist.UpdatedAt,
		}
		responses = append(responses, response)
	}
	return responses
}

// TransformArtistToArtistResponse takes an Artist and transforms it into an ArtistResponse.
func TransformArtistToArtistResponse(artist model.Artist) schema.ArtistResponse {
	return schema.ArtistResponse{
		ID:    artist.ID,
		Name:  artist.Name,
		Image: artist.Image,
	}
}
