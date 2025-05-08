package database

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	//Carregando o arquivo .env com os dados para conexão ao Postgres
	err := godotenv.Load()

	//Mensagem de erro se houver falhar ao tentar ler o arquivo .env
	if err != nil {
		log.Fatal("Houve um erro ao carregador os dados do arquivo .env", err)
	}

	//Usando a lib godotenv para buscar os dados do banco de dados
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)

	//Iniciando a conexão ao banco de dados Postgress
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	//Mensagem de erro se houver uma falha ao tentar realizar a conexão
	if err != nil {
		log.Fatal("Houve um erro ao tentar conectar ao banco de dados:", err)
	}

	//Log informando ao usuario sobre conexão bem sucedida
	DB = db
	log.Println("Conexão ao banco de dados bem sucedida")
}
