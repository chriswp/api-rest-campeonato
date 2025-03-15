package api

import (
	"github.com/chriswp/api-rest-campeonato/internal/infra"
	"github.com/chriswp/api-rest-campeonato/internal/infra/registry"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"log"
	"net/http"
)

type Server struct {
	addr     string
	registry *registry.Registry
}

func NewAPIServer(addr string, registry *registry.Registry) *Server {
	return &Server{
		addr:     addr,
		registry: registry,
	}
}

func (s *Server) Run() error {
	router := gin.Default()
	subrouter := router.Group("/api/v1")

	userHandler := infra.NewUserHandler(s.registry.Database)
	userHandler.RegisterRoutes(subrouter)

	competitionHandler := infra.NewCompetitionHandler()
	competitionHandler.RegisterRoutes(subrouter)

	subrouter.GET("/docs/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	log.Println("Listening on", s.addr)
	return http.ListenAndServe(s.addr, router)
}
