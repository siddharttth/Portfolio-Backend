package services

import (
	"portfolio-backend/models"

	"github.com/gin-gonic/gin"
)

func GetEducation(c *gin.Context) {
	education := []models.Education{
		{
			Degree:    "B. Tech (CSE)",
			Institute: "Pranveer Singh Institute of Technology, Kanpur",
			CGPA:      "7.6",
			Year:      "2022-25",
		},
		{
			Degree:    "Diploma (ME)",
			Institute: "Dayalbagh Educational Institute, Agra",
			CGPA:      "7.7",
			Year:      "2019-22",
		},
		{
			Degree:    "High School",
			Institute: "Holy Cross School (ICSE), Ballia",
			CGPA:      "8.0",
			Year:      "2018-19",
		},
	}
	c.JSON(200, education)
}
