package connector

import (
	"fmt"
	"os"

	"github.com/abel-cosmic/streamy-api/config"
	"github.com/abel-cosmic/streamy-api/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Connect to the database
func Connect() *gorm.DB {
	dbName := config.Config("DB_NAME")
	dbUser := config.Config("DB_USER")
	dbPass := config.Config("DB_PASS")
	dbHost := config.Config("DB_HOST")
	dbPort := config.Config("DB_PORT")
	conn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbUser, dbPass, dbHost, dbPort, dbName)

	fmt.Println("Attempting to connect to the database...")
	db, err := gorm.Open(mysql.Open(conn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	fmt.Println("Database connection successful!")

	if os.Getenv("DB_MIGRATE") == "true" {
		fmt.Println("Migrating database tables...")
		Migrate(db)
		fmt.Println("Database migration completed successfully!")
	}

	return db
}

// Migrate tables
func Migrate(db *gorm.DB) {
	// Auto-migrate all the models
	db.AutoMigrate(&model.Album{}, &model.Artist{}, &model.Interaction{}, &model.PasswordReset{}, &model.Playlist{}, &model.PlaylistSong{}, &model.Song{}, &model.User{})
}
