package services

import (
	"portfolio-backend/models"

	"github.com/gin-gonic/gin"
)

func GetExperience(c *gin.Context) {
	experience := []models.Experience{
		{
			Title:    "Software Development Engineer Intern",
			Company:  "ScaleX",
			Location: "Bengaluru, India",
			Period:   "Nov 2023 - Nov 2024",
			Achievements: []string{
				"Developed proficiency in multiple programming languages including Golang, React, and JavaScript, while also gaining hands-on experience with MongoDB for database management",
				"Utilized various industry-standard tools and technologies such as Redis, RabbitMQ, Kafka, and Postman to enhance development workflow and application performance",
				"Contributed to backend development of diverse projects, including a high-profile astrology application \"Horocosmo\" for Astro Arun Pandit which crossed 10K+ downloads in 1 month",
				"Acquired practical experience in version control and collaborative development using Git and GitHub, implementing best practices for code management and team collaboration",
			},
		},
	}
	c.JSON(200, experience)
}
