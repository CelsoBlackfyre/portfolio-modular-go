package controllers

import (
	"portfolio-modular-go/models"

	"github.com/gofiber/fiber/v2"
)

func GetCertification(c *fiber.Ctx) error {
	certification := []models.Certification{
		{
			ID:     1,
			Title:  "Introdução a Cibersegurança",
			Issuer: "Cisco Skills for All",
			Date:   "2024",
		},
		{
			ID:    2,
			Title: "Python Essentials",
			Date:  "2024",
		},
	}
	return c.JSON(certification)
}
