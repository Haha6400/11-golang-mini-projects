package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type album struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Artist string `json:"artist"`
	// Println string `json:"println"`
}

var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane"},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan"},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan"},
}

/* When the client makes a request at GET /albums
Write the following:
- Logic to prepare a response
- Code to map the request path to your logic
*/

// gin.Context carries request details, validates and serializes JSON,....
func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums) //can replace c.IndentedJSON with c.JSON
}

func addANewAlbum(c *gin.Context) {
	var newAlbum album
	//Receive JSON to newAlbum by calling BindJSON
	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}
	//Add newAlbum to database
	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

func getAlbumByID(c *gin.Context) {
	id := c.Param("id")

	for _, a := range albums {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

func main() {
	router := gin.Default()

	//Get all albums
	router.GET("/albums", getAlbums)
	//Post a new album
	router.POST("/albums", addANewAlbum)
	//Get album by ID
	router.GET("/albums/:id", getAlbumByID)

	router.Run("localhost:8080")
}
