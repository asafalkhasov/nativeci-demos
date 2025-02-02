package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Course struct {
	Name  string `json:"name"`
	Price int    `json:"price"`
}

func main() {
	r := gin.Default()
	r.GET("/api", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
		})
	})
	r.GET("/api/products", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"products": []string{"Computer", "Iphone"},
		})
	})
	r.GET("/api/courses", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"courses": []Course{
				{
					Name:  "Rust for beginners",
					Price: 10,
				},
				{
					Name:  "Bootcamp 12 weeks",
					Price: 200,
				},
				{
					Name:  "Bootcamp for 24 weeks",
					Price: 300,
				},
				{
					Name:  "Bootcamp for 1 year",
					Price: 20203,
				},
				{
					Name:  "University 5 years",
					Price: 373235,
				},
			},
		})
	})
	r.GET("/api/locations", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"locations": []string{
				"UK",
				"USA",
				"Azerbaijan",
				"Russia",
				"Canada",
				"New",
				"New Zealand",
			},
		})
	})
	r.Run("0.0.0.0:8000")
}
