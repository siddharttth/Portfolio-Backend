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
			TechStack: "Go, MongoDB, React, JWT, SMTP, Docker, Kubernetes",
			Points: []string{
				"Built backend in Go, managing job posts, profiles, and applications with MongoDB",
				"Integrated JWT authentication and email verification for secure user access",
				"Developed 75% of the frontend with React for a responsive UI",
				"Implemented features for job posting, browsing, and applying, with data stored in MongoDB",
				"Configured Docker containers and used Kubernetes for deployment",
			},
		},

		{
			Name:      "Real-Time Multiplayer Chess Game with Socket.IO",
			Period:    "Sept 2024",
			TechStack: "Node.js, Express, Socket.IO, Chess.js, HTML/CSS, WebSocket, Heroku",
			Points: []string{
				"Developed real-time chess game with Socket.IO for seamless communication",
				"Implemented drag-and-drop moves with validation using Chess.js",
				"Added spectator mode for real-time viewing without interaction",
				"Synchronized player moves and auto-assigned roles (White/Black)",
				"Deployed the game on Heroku with WebSocket support",
			},
		},

		{
			Name:      "Real-Time Eye-Tracking",
			Period:    "Aug 2023 - Oct 2023",
			TechStack: "Python, OpenCV, Pycharm, Webcam, TensorFlow, Flask",
			Points: []string{
				"Developed an eye-tracking system to monitor pupil movement and set coordinates for tracking",
				"Implemented segmentation of eyes and contour detection using Python and OpenCV",
				"Performed feasibility studies covering technical, operational, economic, and legal aspects",
				"Ensured compatibility across platforms like Linux, Windows, and MacOS",
				"Added Flask for serving the results through a web-based interface",
			},
		},

		{
			Name:      "Go Quiz App",
			Period:    "Jan 2024",
			TechStack: "Go, JSON, HTTP, REST, Postman",
			Points: []string{
				"Built a Go-based quiz application for creating and taking quizzes",
				"Implemented RESTful APIs to manage questions and responses",
				"Stored quiz data in JSON format with HTTP endpoints for interaction",
				"Tested APIs using Postman to ensure proper functionality",
			},
		},

		{
			Name:      "Go Weather App",
			Period:    "Feb 2024",
			TechStack: "Go, OpenWeather API, HTTP, JSON, Docker",
			Points: []string{
				"Developed a weather application using Go and OpenWeather API",
				"Fetched real-time weather data for any location through API requests",
				"Displayed weather information such as temperature, humidity, and conditions",
				"Containerized the app using Docker for easy deployment",
			},
		},

		{
			Name:      "Go Simple HTTP Server",
			Period:    "Mar 2024",
			TechStack: "Go, HTTP, HTML/CSS, JSON, AWS EC2",
			Points: []string{
				"Implemented a basic HTTP server using Go to serve static files",
				"Handled form submissions and routed requests to specific endpoints",
				"Processed GET and POST requests for simple web interactions",
				"Deployed the server on AWS EC2 for scalability and availability",
			},
		},
	}
	c.JSON(200, projects)
}
