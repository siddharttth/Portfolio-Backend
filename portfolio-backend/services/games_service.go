package services

import (
	"portfolio-backend/models"

	"github.com/gin-gonic/gin"
)

func GetGames(c *gin.Context) {
	games := []models.Game{
		{
			ID:          "memory",
			Name:        "Memory Match",
			Description: "Test your memory by matching pairs of cards",
			Rules: []string{
				"Click cards to reveal them",
				"Match pairs of identical cards",
				"Complete the game with minimum moves",
			},
			URL: "/games/memory-match", // Add this for routing
		},
		{
			ID:          "snake",
			Name:        "Snake Game",
			Description: "Classic snake game with modern visuals",
			Rules: []string{
				"Use arrow keys to control the snake",
				"Eat food to grow longer",
				"Avoid hitting walls and yourself",
			},
			URL: "/games/snake-game", // Add this for routing
		},
		{
			ID:          "puzzle",
			Name:        "Sliding Puzzle",
			Description: "Arrange the tiles in numerical order",
			Rules: []string{
				"Click tiles to slide them",
				"Arrange numbers in order",
				"Complete with minimum moves",
			},
			URL: "/games/sliding-puzzle", // Add this for routing
		},
	}
	c.JSON(200, games)
}
