package controllers

import (
	"portfolio-modular-go/models"

	"github.com/gofiber/fiber/v2"
)

func GetExperience(c *fiber.Ctx) error {
	experience := []models.Experience{
		{ID: 1,
			Position: "Atendente Operador de Chat",
			Company:  "AEC - Centro de Contatos",
			Tasks:    "Realizei atendimentos via chat para sistemas de recompensas, utilizando Salesforce e CLM para gestão de clientes. Participei de treinamentos internos em fluxos corporativos e ferramentas de CRM, sempre mantendo foco em empatia, clareza e agilidade no atendimento.",
			Duration: "10/2024 - 01/2025",
		},
	}
	return c.JSON(experience)
}
