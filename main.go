package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Review struct {
	ID     string `json:"id"`
	Movie  string `json:"movie"`
	Review string `json:"review"`
	Rating int    `json:"rating"`
}

var reviews = []Review{
	{ID: "1", Movie: "Inception", Review: "Mind-blowing!", Rating: 5},
	{ID: "2", Movie: "Interstellar", Review: "Epic space scenes!", Rating: 4},
}

func main() {
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Welcome to Movie API 🚀"})
	})

	router.GET("/reviews", func(c *gin.Context) {
		c.JSON(http.StatusOK, reviews)
	})

	router.POST("/reviews", func(c *gin.Context) {
		var newReview Review
		if err := c.BindJSON(&newReview); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		reviews = append(reviews, newReview)
		c.JSON(http.StatusCreated, newReview)
	})

	router.PUT("/reviews/:id", func(c *gin.Context) {
		id := c.Param("id")
		var updatedReview Review
		if err := c.BindJSON(&updatedReview); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		for i, r := range reviews {
			if r.ID == id {
				reviews[i] = updatedReview
				c.JSON(http.StatusOK, updatedReview)
				return
			}
		}
		c.JSON(http.StatusNotFound, gin.H{"message": "Review not found"})
	})

	router.DELETE("/reviews/:id", func(c *gin.Context) {
		id := c.Param("id")
		for i, r := range reviews {
			if r.ID == id {
				reviews = append(reviews[:i], reviews[i+1:]...)
				c.JSON(http.StatusOK, gin.H{"message": "Deleted"})
				return
			}
		}
		c.JSON(http.StatusNotFound, gin.H{"message": "Review not found"})
	})

	router.Run(":8080")
}
