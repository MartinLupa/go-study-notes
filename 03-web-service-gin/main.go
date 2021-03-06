package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//=====================================================================================================
func main() {

	//Initializing gin Router.
	router := gin.Default()

	//Defining routes
	router.GET("/albums", getAlbums)
	router.GET("/albums/:id", getAlbumByID)
	router.POST("/albums", postAlbums)

	//Attaching the router to an http.Server
	router.Run("localhost:3000")
}

//=====================================================================================================
//album struct represents data about a record album.
type album struct {
	ID     string  `json: "id"`
	Title  string  `json: "title"`
	Artist string  `json: "artist"`
	Price  float64 `json: "price"`
}

//album slice to seed record album data.
var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

//=====================================================================================================
//gin.Context is the most important part of Gin. It carries request details, validates and serializes JSON, and more. (Different from Go’s built-in context package.)
//getAlbums responds with the list of all albums as JSON.
func getAlbums(c *gin.Context) {
	//WARNING: use IndentedJSON only for development purposes since printing pretty JSON is more CPU and bandwidth consuming. For production use Context.JSON() instead.
	c.IndentedJSON(http.StatusOK, albums)
}

//postAlbums adds an album from JSON received in the request body.
func postAlbums(c *gin.Context) {
	var newAlbum album

	//Call BindJSON to bind the received JSON to newAlbum.
	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	//Add the new album to the slice.
	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

//getAlbumByID locates the album whose ID value matches the id parameter sent by the client, then returns that album as a response.
func getAlbumByID(c *gin.Context) {
	id := c.Param("id")

	//Loop over the list of albums, looking for an album whose ID value matches the parameter.
	for _, a := range albums {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}
