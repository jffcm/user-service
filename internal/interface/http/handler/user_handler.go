package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jffcm/user-service/internal/application/usecase"
)

type UserHandler struct {
	registerUseCase usecase.RegisterUseCase
}

func NewUserHandler(registerUseCase usecase.RegisterUseCase) *UserHandler {
	return &UserHandler{registerUseCase: registerUseCase}
}

func (u *UserHandler) Register(ctx *gin.Context) {
	input := usecase.RegisterUseCaseInput{}
	if err := ctx.BindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"type":   "about:blank",
			"title":  "Invalid JSON",
			"status": http.StatusBadRequest,
			"detail": "The request body could not be parsed. Please ensure it is valid JSON and try again.",
		})
		return
	}

	output, err := u.registerUseCase.Execute(&input)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"type":   "about:blank",
			"title":  "Internal Server Error",
			"status": http.StatusInternalServerError,
			"detail": "An unexpected error occurred while registering the user.",
		})
		return
	}

	ctx.JSON(http.StatusCreated, output)
}
