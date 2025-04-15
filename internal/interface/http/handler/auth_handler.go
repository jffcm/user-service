package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jffcm/user-service/internal/application/usecase"
)

type AuthHandler interface {
	Login(ctx *gin.Context)
}

type authHandler struct {
	loginUseCase usecase.LoginUseCase
}

func NewAuthHandler(loginUseCase usecase.LoginUseCase) AuthHandler {
	return &authHandler{loginUseCase: loginUseCase}
}

func (a *authHandler) Login(ctx *gin.Context) {
	input := usecase.LoginUseCaseInput{}
	if err := ctx.BindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"type": "about:blank",
			"title": "Invalid JSON",
			"status": http.StatusInternalServerError,
			"detail": "The request body could not be parsed. Please ensure it is valid JSON and try again.",
		}) 
		return
	}

	output, err := a.loginUseCase.Execute(&input)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"type":   "about:blank",
			"title":  "Internal Server Error",
			"status": http.StatusInternalServerError,
			"detail": "An unexpected error occurred while trying to log in.",
		})
		return
	}

	ctx.JSON(http.StatusOK, output)
}

