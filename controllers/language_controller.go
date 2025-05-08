package controllers

import (
	"portfolio-modular-go/models"

	"github.com/gofiber/fiber/v2"
)

func GetLanguage(c *fiber.Ctx) error {
	language := []models.Language{
		{
			ID:        1,
			Title:     "Português",
			Profiency: "Nativo",
		},
		{
			ID:        2,
			Title:     "Inglês",
			Profiency: "Avançado",
		},
	}
	return c.JSON(language)
}
