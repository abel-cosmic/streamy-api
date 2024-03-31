package router

import (
	"github.com/abel-cosmic/streamy-api/handler"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// Expenses route
func Routes(app *fiber.App, db *gorm.DB) {
	album := &handler.AlbumHandler{
		DB: db,
	}
	r := app.Group("/albums")
	r.Get("/", album.Index)
	r.Get("/:id", album.Show)
	r.Post("/", album.Store)
	r.Put("/:id", album.Update)
	r.Delete("/:id", album.Destroy)

	artist := &handler.ArtistHandler{
		DB: db,
	}
	r = app.Group("/artists")
	r.Get("/", artist.Index)
	r.Get("/:id", artist.Show)
	r.Post("/", artist.Store)
	r.Put("/:id", artist.Update)
	r.Delete("/:id", artist.Destroy)

	interaction := &handler.InteractionHandler{
		DB: db,
	}
	r = app.Group("/interactions")
	r.Get("/", interaction.Index)
	r.Get("/:id", interaction.Show)
	r.Post("/", interaction.Store)
	r.Put("/:id", interaction.Update)
	r.Delete("/:id", interaction.Destroy)

	passwordReset := &handler.PasswordResetHandler{
		DB: db,
	}
	r = app.Group("/passwordresets")
	r.Get("/", passwordReset.Index)
	r.Get("/:id", passwordReset.Show)
	r.Post("/", passwordReset.Store)
	r.Put("/:id", passwordReset.Update)
	r.Delete("/:id", passwordReset.Destroy)

	playlistSong := &handler.PlaylistSongHandler{
		DB: db,
	}
	r = app.Group("/playlist_songs")
	r.Get("/", playlistSong.Index)
	r.Get("/:id", playlistSong.Show)
	r.Post("/", playlistSong.Store)
	r.Put("/:id", playlistSong.Update)
	r.Delete("/:id", playlistSong.Destroy)

	playlist := &handler.PlaylistHandler{
		DB: db,
	}
	r = app.Group("/playlists")
	r.Get("/", playlist.Index)
	r.Get("/:id", playlist.Show)
	r.Post("/", playlist.Store)
	r.Put("/:id", playlist.Update)
	r.Delete("/:id", playlist.Destroy)

	song := &handler.SongHandler{
		DB: db,
	}
	r = app.Group("/songs")
	r.Get("/", song.Index)
	r.Get("/:id", song.Show)
	r.Post("/", song.Store)
	r.Put("/:id", song.Update)
	r.Delete("/:id", song.Destroy)

	user := &handler.UserHandler{
		DB: db,
	}
	r = app.Group("/users")
	r.Get("/", user.Index)
	r.Get("/:id", user.Show)
	r.Post("/", user.Store)
	r.Put("/:id", user.Update)
	r.Delete("/:id", user.Destroy)

}
