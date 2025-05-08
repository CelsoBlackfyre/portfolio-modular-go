package main

import (
	"portfolio-modular-go/database"
	"portfolio-modular-go/models"
	"portfolio-modular-go/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	//Call para a função de conexão ao banco de dados
	database.Connect()

	//Fazendo migração automatica dos modelos para o banco de dados, um modelo de cada vez
	database.DB.AutoMigrate(&models.Project{})
	database.DB.AutoMigrate(&models.Certification{})
	database.DB.AutoMigrate(&models.Education{})
	database.DB.AutoMigrate(&models.Experience{})
	database.DB.AutoMigrate(&models.Interest{})
	database.DB.AutoMigrate(&models.Language{})
	database.DB.AutoMigrate(&models.Skills{})
	database.DB.AutoMigrate(&models.User{})

	//Iniciando app com a lib fiber
	app := fiber.New()

	//Definindo as rotas do app
	routes.SetupRoutes(app)

	//Definindo em qual porta o aplicativo irá funcionar
	app.Listen(":3000")
}
