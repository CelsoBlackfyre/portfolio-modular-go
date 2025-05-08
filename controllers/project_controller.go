package controllers

import (
	"portfolio-modular-go/database"
	"portfolio-modular-go/models"

	"github.com/gofiber/fiber/v2"
)

func GetProjects(c *fiber.Ctx) error {
	projects := []models.Project{
		{
			ID:          1,
			Title:       "Blog pessoal com Next.js",
			Description: "Blog pessoal desenvolvido utilizando Next.js, ou seja, contém uso de React e Tailwind para o frontend e Golang no backend.",
			Link:        "github.com/",
		},
		{
			ID:          2,
			Title:       "Portfolio modular utilizando Golang",
			Description: "Portfolio modular criado utilizando Golang. Utilizado PostgreSQL e terá estrutura escalavel",
			Link:        "github.com/",
		},
	}
	return c.JSON(projects)
}

func CreateProject(c *fiber.Ctx) error {
	project := new(models.Project)

	if err := c.BodyParser(project); err != nil {
		return c.Status(400).JSON(fiber.Map{"erro": "Entrada invalida"})
	}

	database.DB.Create(&project)
	return c.JSON(project)
}
