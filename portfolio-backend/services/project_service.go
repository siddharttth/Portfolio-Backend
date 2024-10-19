package services

import (
	"portfolio-backend/models"

	"github.com/gin-gonic/gin"
)

func GetProjects(c *gin.Context) {
	projects := []models.Project{
		{
			Name:      "Hiring Platform",
			Period:    "Aug 2024 - Oct 2024",
			TechStack: "Go, MongoDB, React, JWT, SMTP email service",
			Points: []string{
				"Built backend in Go, managing job posts, profiles, and applications with MongoDB",
				"Integrated JWT authentication and email verification for secure user access",
				"Developed 75% of the frontend with React for a responsive UI",
				"Implemented features for job posting, browsing, and applying, with data stored in MongoDB",
			},
		},

		{
			Name:      "Real-Time Multiplayer Chess Game with Socket.IO",
			Period:    "Sept 2024",
			TechStack: "Node.js, Express, Socket.IO, Chess.js, HTML/CSS",
			Points: []string{
				"Developed real-time chess game with Socket.IO for seamless communication",
				"Implemented drag-and-drop moves with validation using Chess.js",
				"Added spectator mode for real-time viewing without interaction",
				"Synchronized player moves and auto-assigned roles (White/Black)",
				"Planned future features: player authentication, AI, and improved error handling",
			},
		},
	}
	c.JSON(200, projects)
}
