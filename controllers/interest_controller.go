package controllers

import (
	"portfolio-modular-go/models"

	"github.com/gofiber/fiber/v2"
)

func GetInterest(c *fiber.Ctx) error {
	interest := []models.Interest{
		{
			ID:    1,
			Title: "Arquitetura de Software",
		},
		{
			ID:    2,
			Title: "Desenvolvimento Back-end",
		},
		{
			ID:    3,
			Title: "Web-Development",
		},
		{
			ID:    4,
			Title: "Linux e DevOps",
		},
		{
			ID:    5,
			Title: "Projetos Open Source",
		},
		{
			ID:    6,
			Title: "Inovação Tecnológica",
		},
	}
	return c.JSON(interest)
}
