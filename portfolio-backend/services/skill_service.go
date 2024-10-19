package services

import (
	"portfolio-backend/models"

	"github.com/gin-gonic/gin"
)

func GetSkills(c *gin.Context) {
	skills := []models.Skill{
		{
			Category: "Programming Languages",
			Items:    []string{"Go", "JavaScript", "Python", "C", "C++", "React", "Express"},
		},
		{
			Category: "Tech Stacks",
			Items:    []string{"MERN", "MEAN", "PERN"},
		},
		{
			Category: "Tools",
			Items:    []string{"Redis", "Apache Kafka", "RabbitMQ", "AWS", "PostgreSQL", "MongoDB"},
		},
	}
	c.JSON(200, skills)
}
