package main

import (
	"fmt"
	"log"

	"github.com/febriansr/jwt-auth/controller"
	"github.com/febriansr/jwt-auth/repository"
	"github.com/febriansr/jwt-auth/usecase"
	"github.com/febriansr/jwt-auth/utils"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func main() {
	dbHost := utils.DotEnv("DB_HOST")
	dbPort := utils.DotEnv("DB_PORT")
	dbUser := utils.DotEnv("DB_USER")
	dbPassword := utils.DotEnv("DB_PASSWORD")
	dbName := utils.DotEnv("DB_NAME")
	sslMode := utils.DotEnv("SSL_MODE")
	serverPort := utils.DotEnv("SERVER_PORT")

	dataSourceName := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s", dbUser, dbPassword, dbHost, dbPort, dbName, sslMode)
	db, err := sqlx.Connect("postgres", dataSourceName)
	defer func() {
		err := db.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("Connected to DB")
	}

	router := gin.Default()

	userRepo := repository.NewUserRepo(db)
	userUsecase := usecase.NewUserUsecase(userRepo)
	userController := controller.NewUserController(userUsecase)

	router.POST(`/auth/login`, userController.Login)

	if err := router.Run(serverPort); err != nil {
		log.Fatal(err)
	}
}
