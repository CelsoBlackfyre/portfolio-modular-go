package controllers

import (
	"portfolio-modular-go/models"

	"github.com/gofiber/fiber/v2"
)

func GetSkills(c *fiber.Ctx) error {
	skills := []models.Skills{
		{
			ID:    1,
			Title: "C++",
		},
		{
			ID:    2,
			Title: "Python",
		},
		{
			ID:    3,
			Title: "Golang",
		},
		{
			ID:    4,
			Title: "HTML",
		},
		{
			ID:    5,
			Title: "CSS",
		},
		{
			ID:    6,
			Title: "Javascript",
		},
		{
			ID:    7,
			Title: "Next.js",
		},
		{
			ID:    8,
			Title: "Tailwind",
		},
		{
			ID:    9,
			Title: "Git",
		},
		{
			ID:    10,
			Title: "Linux",
		},
		{
			ID:    11,
			Title: "VSCode",
		},
		{
			ID:    12,
			Title: "Fundamentos de Cibersegurança",
		},
		{
			ID:    13,
			Title: "PHP",
		},
		{
			ID:    14,
			Title: "Laravel",
		},
	}
	return c.JSON(skills)
}
