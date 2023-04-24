package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Album struct for storing albums
type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

// Creating dummy albums
var albums = []album{
	{
		ID:     "1",
		Title:  "Blue Train",
		Artist: "John Coltrane",
		Price:  56.99,
	},
	{
		ID:     "2",
		Title:  "Jeru",
		Artist: "Gerry Mulligan",
		Price:  17.99,
	},
}

func main() {
	fmt.Println("Starting the API...")
	// Initiating the router
	router := gin.Default()

	// Registering all endpoints
	router.GET("/albums", getAlbums)
	router.POST("/albums", postAlbums)

	// Running the API
	router.Run("localhost:8080")
}

// getAlbums function that responds with a list of all albums as JSON
func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

// postAlbums function that adds an album from received JSON
func postAlbums(c *gin.Context) {
	var newAlbum album

	// Binding the received JSON to the album struct
	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)

}
