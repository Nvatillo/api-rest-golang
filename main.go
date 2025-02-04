package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type album struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Artist string `json:"artist"`
	Year   int    `json:"year"`
}

var albums = []album{
	{ID: "1", Title: "Familia", Artist: "Camila Cabello", Year: 2022},
	{ID: "2", Title: "salsa", Artist: "Camila Cabello", Year: 2021},
	{ID: "3", Title: "Fofos", Artist: "Camila Cabello", Year: 2024},
}

func getAlbums(c *gin.Context) {
	c.JSON(http.StatusOK, albums)
}

func postAlbums(c *gin.Context) {
	var newAlbum album

	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	albums = append(albums, newAlbum)

	c.JSON(http.StatusCreated, albums)
}

func getAlbumById(c *gin.Context) {
	id := c.Param("id_album")

	for _, a := range albums {
		if a.ID == id {
			c.JSON(http.StatusOK, a)
			return
		}
	}
	c.JSON(http.StatusBadRequest, nil)
}

func main() {
	router := gin.Default()

	// TODO define routes

	router.GET("/albums", getAlbums)
	router.POST("/postalbums", postAlbums)
	router.GET("/search/:id", getAlbumById)
	router.Run("localhost:8080")

}
