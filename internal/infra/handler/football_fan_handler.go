package handler

import (
	"github.com/chriswp/api-rest-campeonato/internal/domain/entity"
	"github.com/chriswp/api-rest-campeonato/internal/dto"
	"github.com/chriswp/api-rest-campeonato/internal/usecase"
	"github.com/gin-gonic/gin"
	"net/http"
)

type FootballFanResponse struct {
	Mensagem string              `json:"mensagem"`
	Fan      *entity.FootballFan `json:"fan"`
}

type FootballFanHandler struct {
	FootballFanUseCase *usecase.FootballFanUseCase
}

func NewFootballFanHandler(useCase *usecase.FootballFanUseCase) *FootballFanHandler {
	return &FootballFanHandler{FootballFanUseCase: useCase}
}

func (h *FootballFanHandler) RegisterRoutes(router *gin.RouterGroup) {
	router.POST("/football-fan", h.CreateFootballFan)
}

// CreateFootballFan godoc
// @Summary Cria um novo torcedor de futebol
// @Description Cria um novo torcedor de futebol com os dados fornecidos e retorna o torcedor criado
// @Tags FootballFan
// @Accept  json
// @Produce  json
// @Param input body dto.CreateFootballFanDTO true "Dados do Torcerdor de Futebol"
// @Success 200 {object} FootballFanResponse "Sucesso"
// @Failure 400 {object} ErrorResponse "Requisição inválida"
// @Failure 500 {object} ErrorResponse "Erro interno do servidor"
// @Security ApiKeyAuth
// @Router /api/v1/football-fan [post]
func (h *FootballFanHandler) CreateFootballFan(c *gin.Context) {
	var input dto.CreateFootballFanDTO
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	fan := &entity.FootballFan{
		Name:  input.Name,
		Email: input.Email,
		Team:  input.Team,
	}

	createdFan, err := h.FootballFanUseCase.CreateFootballFan(c, fan)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"mensagem": "registration completed successfully",
		"fan":      createdFan,
	})
}
