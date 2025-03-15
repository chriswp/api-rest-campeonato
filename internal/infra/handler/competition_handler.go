package handler

import (
	"fmt"
	"github.com/chriswp/api-rest-campeonato/internal/constants"
	"github.com/chriswp/api-rest-campeonato/internal/usecase"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type ErrorResponse struct {
	Error string `json:"error"`
}

type CompetitionHandler struct {
	CompetitionUsecase *usecase.CompetitionUseCase
}

func NewCompetitionHandler(competitionUsecase *usecase.CompetitionUseCase) *CompetitionHandler {
	return &CompetitionHandler{
		CompetitionUsecase: competitionUsecase,
	}
}

func (h *CompetitionHandler) RegisterRoutes(router *gin.RouterGroup) {
	router.GET("/competitions", h.GetCompetitions)
	router.GET("/competitions/:id/matches", h.GetMatchesByCompetitions)
}

// GetCompetitions godoc
// @Summary Obtém todas as competições
// @Description Retorna uma lista de competições disponíveis
// @Tags Competitions
// @Accept  json
// @Produce  json
// @Success 200 {array} entity.Competition
// @Failure 500 {object} ErrorResponse
// @Security ApiKeyAuth
// @Router /api/v1/competitions [get]
func (h *CompetitionHandler) GetCompetitions(c *gin.Context) {
	competitions, err := h.CompetitionUsecase.GetCompetitions()
	if err != nil {
		c.JSON(500, ErrorResponse{Error: err.Error()})
		return
	}
	c.JSON(200, competitions)
}

// GetMatchesByCompetitions retorna as partidas de uma competição específica.
// @Summary      Retorna partidas de uma competição
// @Description  Retorna as partidas de uma competição pelo ID, podendo filtrar por rodada e equipe.
// @Tags         Competitions
// @Accept       json
// @Produce      json
// @Param        id       path      int     true  "ID da Competição"
// @Param        rodada   query     int     false "Número da Rodada"
// @Param        equipe   query     string  false "Nome da Equipe"
// @Success      200      {object}  []entity.Match
// @Failure      400      {object}  map[string]string "Erro nos parâmetros da requisição"
// @Failure      500      {object}  map[string]string "Erro interno do servidor"
// @Router       /api/v1/competitions/{id}/matches [get]
func (h *CompetitionHandler) GetMatchesByCompetitions(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": constants.InvalidID})
		return
	}

	paramMatchday := c.Query("rodada")
	var matchday *int
	if paramMatchday != "" {
		paramMatchdayInt, err := strconv.Atoi(paramMatchday)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("%s: %s", constants.InvalidID, "rodada")})
			return
		}
		matchday = &paramMatchdayInt
	}

	var teamFilter *string
	paramTeam := c.Query("equipe")
	if paramTeam != "" {
		teamFilter = &paramTeam
	}

	matches, err := h.CompetitionUsecase.GetMatches(id, matchday, teamFilter)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, matches)
}
