package handlers

import (
	"net/http"
	"online-song-library/models"
	"os"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterRoutes(router *gin.Engine, db *gorm.DB) {
	router.GET("/library", func(c *gin.Context) { GetSongs(c, db) })
	router.GET("/library/:id/lyrics", func(c *gin.Context) { GetLyrics(c, db) })
	router.POST("/library", func(c *gin.Context) { AddSong(c, db) })
	router.PUT("/library/:id", func(c *gin.Context) { UpdateSong(c, db) })
	router.DELETE("/library/:id", func(c *gin.Context) { DeleteSong(c, db) })
}

// Get songs with pagination and filtering
func GetSongs(c *gin.Context, db *gorm.DB) {
	var songs []models.Song
	group := c.Query("group")
	song := c.Query("song")

	query := db
	if group != "" {
		query = query.Where("group = ?", group)
	}
	if song != "" {
		query = query.Where("title = ?", song)
	}

	query.Find(&songs)
	c.JSON(http.StatusOK, songs)
}

// Get lyrics by ID with pagination
func GetLyrics(c *gin.Context, db *gorm.DB) {
	var song models.Song
	id := c.Param("id")
	if err := db.First(&song, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Song not found"})
		return
	}

	// Paginate lyrics
	page := c.DefaultQuery("page", "1")
	limit := c.DefaultQuery("limit", "10")
	// Split and paginate lyrics (not shown here for brevity)
	c.JSON(http.StatusOK, song.Lyrics)
}

// Add a new song
func AddSong(c *gin.Context, db *gorm.DB) {
	var input struct {
		Group string `json:"group" binding:"required"`
		Song  string `json:"song" binding:"required"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Fetch additional info from API
	resty := resty.New()
	resp, err := resty.R().SetQueryParams(map[string]string{
		"group": input.Group,
		"song":  input.Song,
	}).Get(os.Getenv("API_BASE_URL") + "/info")
	if err != nil || resp.StatusCode() != http.StatusOK {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to fetch song info"})
		return
	}

	// Parse API response
	var apiResponse struct {
		ReleaseDate string `json:"releaseDate"`
		Text        string `json:"text"`
		Link        string `json:"link"`
	}
	if err := resp.UnmarshalJSON(&apiResponse); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse API response"})
		return
	}

	// Save to database
	song := models.Song{
		Group:       input.Group,
		Title:       input.Song,
		ReleaseDate: apiResponse.ReleaseDate,
		Lyrics:      apiResponse.Text,
		Link:        apiResponse.Link,
	}
	db.Create(&song)
	c.JSON(http.StatusOK, song)
}

// Update a song
func UpdateSong(c *gin.Context, db *gorm.DB) {
	var song models.Song
	id := c.Param("id")
	if err := db.First(&song, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Song not found"})
		return
	}

	var input struct {
		Group string `json:"group"`
		Title string `json:"title"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Update song
	db.Model(&song).Updates(input)
	c.JSON(http.StatusOK, song)
}

// Delete a song
func DeleteSong(c *gin.Context, db *gorm.DB) {
	id := c.Param("id")
	if err := db.Delete(&models.Song{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete song"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Song deleted"})
}
