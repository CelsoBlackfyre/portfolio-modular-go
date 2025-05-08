package controllers

import (
	"portfolio-modular-go/models"

	"github.com/gofiber/fiber/v2"
)

func GetEducation(c *fiber.Ctx) error {
	education := []models.Education{
		{
			ID:           1,
			StudyProgram: "Jogos Digitais",
			Institution:  "PUC - MG",
			Duration:     "02/2018 - 08/2019",
		},
		{
			ID:           2,
			StudyProgram: "Jogos Digitais",
			Institution:  "Universidade FUMEC",
			Duration:     "08/2019 - 08/2020",
		},
		{
			ID:           3,
			StudyProgram: "Técnico em Desenvolvimento de Sistemas",
			Institution:  "Senai - MG",
			Duration:     "08/2023 - 11/2024",
		},
		{
			ID:           4,
			StudyProgram: "Engenharia de Software",
			Institution:  "GranFaculdade",
			Duration:     "08/2024 - Presente",
		},
	}
	return c.JSON(education)
}
