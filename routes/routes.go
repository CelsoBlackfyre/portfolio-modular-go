package routes

import (
	"portfolio-modular-go/controllers"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api")

	api.Get("/projects", controllers.GetProjects)
	api.Post("/projects", controllers.CreateProject)

	api.Get("/user", controllers.GetUser)
	api.Get("/skills", controllers.GetSkills)
	api.Get("/languages", controllers.GetLanguage)
	api.Get("/certifications", controllers.GetCertification)
	api.Get("/education", controllers.GetEducation)
	api.Get("/interests", controllers.GetInterest)
	api.Get("/experience", controllers.GetExperience)
}
