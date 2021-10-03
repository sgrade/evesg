package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// album represents data about a record album.
type item struct {
	ID    string  `json:"id"`
	Name  string  `json:"name"`
	Type  string  `json:"type"`
	Price float64 `json:"price"`
}

// albums slice to seed record album data.
var items = []item{
	{ID: "1", Name: "Hoarder", Type: "Ship", Price: 56.99},
	{ID: "2", Name: "Mammoth", Type: "Ship", Price: 17.99},
	{ID: "3", Name: "Iterion Mark V", Type: "Ship", Price: 39.99},
}

func main() {
	router := gin.Default()
	router.GET("/items", getItems)
	router.POST("/items", postItems)

	router.Run("localhost:8080")
}

// getAlbums responds with the list of all albums as JSON.
func getItems(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, items)
}

// postAlbums adds an album from JSON received in the request body.
func postItems(c *gin.Context) {
	var newItem item

	// Call BindJSON to bind the received JSON to
	// newAlbum.
	if err := c.BindJSON(&newItem); err != nil {
		return
	}

	// Add the new album to the slice.
	items = append(items, newItem)
	c.IndentedJSON(http.StatusCreated, newItem)
}
