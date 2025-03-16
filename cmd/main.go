package main

import (
	"context"
	"github.com/chriswp/api-rest-campeonato/cmd/api"
	config "github.com/chriswp/api-rest-campeonato/configs"
	_ "github.com/chriswp/api-rest-campeonato/internal/docs"
	"github.com/chriswp/api-rest-campeonato/internal/infra"
	"github.com/chriswp/api-rest-campeonato/internal/infra/db"
	"github.com/chriswp/api-rest-campeonato/internal/usecase"
	"log"
)

// @title         	API Rest Competition
// @version         1.0
// @description     Competition API with auhtentication
// @termsOfService  http://swagger.io/terms/

// @contact.name   Christopher Fernandes
// @host      localhost:8080
// @BasePath  /
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	ctx := context.Background()
	config.LoadConfig()

	connectionDB, err := db.NewPostgresConnection(ctx)
	if err != nil {
		log.Fatalf("Erro ao conectar ao banco: %v", err)
	}

	reg, err := infra.NewRegistry(connectionDB)
	if err != nil {
		log.Fatal("Erro ao inicializar o Registry:", err)
	}
	defer func() {
		if err := reg.Close(); err != nil {
			log.Println("Erro ao fechar conexão com o banco:", err)
		}
	}()

	authUseCase := usecase.NewAuthUseCase(reg.UserRepo)
	authMiddleware, err := config.NewAuthMiddleware(authUseCase)
	if err != nil {
		log.Fatal("Erro ao criar middleware JWT:", err)
	}
	config.Envs.TokenAuth = authMiddleware

	server := api.NewAPIServer(":8080", reg)
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}
