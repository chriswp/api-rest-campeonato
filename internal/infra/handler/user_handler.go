package handler

import (
	config "github.com/chriswp/api-rest-campeonato/configs"
	"github.com/chriswp/api-rest-campeonato/internal/usecase"
	"github.com/gin-gonic/gin"
)

type LoginRequest struct {
	Email    string `json:"email" example:"user@test.com"`
	Password string `json:"password" example:"123456"`
}

type UserHandler struct {
	UserUseCase *usecase.UserUseCase
}

func NewUserHandler(useCase *usecase.UserUseCase) *UserHandler {
	return &UserHandler{UserUseCase: useCase}
}

func (h *UserHandler) RegisterRoutes(router *gin.RouterGroup) {
	router.POST("/login", h.Login)
	router.GET("/refresh_token", config.Envs.TokenAuth.RefreshHandler)
}

// Login realiza a autenticação do usuário e retorna um token JWT.
// @Summary Realiza o login do usuário
// @Description Autentica o usuário e retorna um token JWT para acesso às rotas protegidas.
// @Tags Auth
// @Accept json
// @Produce json
// @Param request body LoginRequest true "Dados de login"
// @Success 200 {object} map[string]interface{} "Token gerado com sucesso"
// @Failure 401 {object} map[string]interface{} "Usuário ou senha incorretos"
// @Router /api/v1/login [post]
func (h *UserHandler) Login(c *gin.Context) {
	config.Envs.TokenAuth.LoginHandler(c)
}
