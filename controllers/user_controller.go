package controllers

import (
	"portfolio-modular-go/models"

	"github.com/gofiber/fiber/v2"
)

func GetUser(c *fiber.Ctx) error {
	user := []models.User{
		{
			ID:          1,
			Name:        "Celso Carvalho Jr.",
			Title:       "Desenvolvedor Full Stack em formação",
			Bio:         "Entusiasta da tecnologia com sólida formação em Desenvolvimento de Sistemas e forte interesse por Cibersegurança. Experiência acadêmica em Jogos Digitais pela PUC Minas e Universidade FUMECddddd. Destaque atual no curso técnico de Desenvolvimento de Sistemas no SENAI MG. Possuo certificados da Cisco em Introdução à Cibersegurança e Python Essentials. Experiência prática com C++, C#, Python, Golang, HTML e CSS. Vivência profissional como operador de chat, lidando com sistemas de recompensas, atendimento ao cliente e resolução de problemas em ambiente dinâmico.",
			Email:       "celsomartinscarvalho@gmail.com",
			Github:      "github.com/CelsoBlackfyre",
			Linkedin:    "linkedin/in/celsomartins1998",
			PhoneNumber: "31 98497-7467",
			Address:     "Belo Horizonte, MG, Brasil",
		},
	}
	return c.JSON(user)
}
