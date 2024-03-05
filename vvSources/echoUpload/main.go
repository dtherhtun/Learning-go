package main

import (
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/lib/pq" // Import the pq library for handling PostgreSQL array types
	"gorm.io/gorm"
)

type RoomType struct {
	gorm.Model
	ID          uuid.UUID      `json:"id" gorm:"primary_key;type:uuid;default:uuid_generate_v4()"`
	Name        string         `json:"name"`
	Description string         `json:"description"`
	Photo       pq.StringArray `json:"photo" gorm:"type:text[]"`
	Price       float64        `json:"price"`
}

// CustomBinder is a custom Echo binder to handle multipart/form-data and JSON in the same request
type CustomBinder struct{}

// Bind implements the echo.Binder interface
func (cb *CustomBinder) Bind(i interface{}, c echo.Context) error {
	if c.Request().Header.Get("Content-Type") == "application/json" {
		return c.Bind(i)
	} else {
		formBinder := new(echo.DefaultBinder)
		return formBinder.Bind(i, c)
	}
}

func main() {
	e := echo.New()
	e.Binder = &CustomBinder{} // Use the custom binder

	// Middleware

	// Routes
	e.POST("/create-room", createRoomHandler)

	// Start server
	e.Start(":8080")
}

// Handler for creating a room with photo upload
func createRoomHandler(c echo.Context) error {
	room := new(RoomType)
	if err := c.Bind(room); err != nil {
		return err
	}

	// Handle photo uploads
	form, err := c.MultipartForm()
	if err != nil {
		return err
	}

	files := form.File["photo"]
	for _, file := range files {
		src, err := file.Open()
		if err != nil {
			return err
		}
		defer src.Close()

		// Process the file (e.g., save to storage, database, etc.)
		// In this example, we print the file name
		fmt.Println("Uploaded File:", file.Filename)

		// Append the file URL to the Photo field
		room.Photo = append(room.Photo, "path/to/uploaded/"+file.Filename)
	}

	// Save the room to the database (adjust this based on your database setup)
	// For simplicity, this example uses Gorm with a PostgreSQL database

	return c.JSON(http.StatusOK, room)
}
