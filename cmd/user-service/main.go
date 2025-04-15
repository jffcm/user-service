package main

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jffcm/user-service/internal/application/usecase"
	"github.com/jffcm/user-service/internal/domain/service"
	"github.com/jffcm/user-service/internal/infrastructure/database/postgres/repository"
	"github.com/jffcm/user-service/internal/interface/http/handler"
	_ "github.com/lib/pq"
)

func main() {
	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	port := os.Getenv("DB_PORT")

	driverName := "postgres"
	dataSourceName := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		host, user, password, dbname, port,
	)

	db, err := sql.Open(driverName, dataSourceName)
	if err != nil {
		return
	}

	router := gin.Default()

	postgresUserRepository := repository.NewPostgresUserRepository(db)
	bcryptHasher := service.NewBcryptPasswordHasher()

	secretKey := os.Getenv("SECRET_KEY")
	jwtGenerator := service.NewJWTTokenGenerator(secretKey)

	registerUseCase := usecase.NewRegisterUseCase(postgresUserRepository, bcryptHasher)
	loginUseCase := usecase.NewLoginUseCase(postgresUserRepository, bcryptHasher, jwtGenerator)

	userHandler := handler.NewUserHandler(registerUseCase)
	authHandler := handler.NewAuthHandler(loginUseCase)

	relativePath := "/api/v1"
	v1 := router.Group(relativePath)
	{
		v1.POST("/auth/login", authHandler.Login)
		v1.POST("/users", userHandler.Register)
	}

	router.Run(":8080")
}
