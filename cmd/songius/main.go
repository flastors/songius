package main

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"time"

	_ "github.com/flastors/songius/docs"
	"github.com/flastors/songius/internal/api"
	"github.com/flastors/songius/internal/config"
	"github.com/flastors/songius/internal/database/migration"
	"github.com/flastors/songius/internal/music"
	musicDB "github.com/flastors/songius/internal/music/db"
	musicService "github.com/flastors/songius/internal/music/service"
	"github.com/flastors/songius/pkg/client/postgresql"
	"github.com/flastors/songius/pkg/utils/logging"
	"github.com/julienschmidt/httprouter"
)

// @title           Songius
// @version         1.0
// @description     This is a simple music library.
// @termsOfService  http://swagger.io/terms/

// @contact.name   Flastor

// @host      localhost:8080
// @BasePath  /api/v1

// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/
func main() {
	logger := logging.GetLogger()
	conf := config.GetConfig()

	logger.Info("Connecting to database")
	migrator, err := migration.NewMigration(conf.Storage)
	if err != nil {
		logger.Fatal(fmt.Errorf("failed to create migration instance: %v", err))
	}
	err = migrator.Up()
	if err != nil {
		logger.Warn(fmt.Errorf("failed to migrate: %v", err))
	}
	migrator.Close()

	var postgreSQLClient postgresql.Client
	postgreSQLClient, err = postgresql.NewClient(context.Background(), 3, conf.Storage)
	if err != nil {
		logger.Fatal(err)
	}

	repo := musicDB.NewRepository(postgreSQLClient, logger)
	clientAPI := api.NewAPIClient(&conf.ExternalAPI)
	service := musicService.NewService(repo, clientAPI, logger)

	router := httprouter.New()
	logger.Info("Registering handlers")
	handler := music.NewHandler(service, logger)
	handler.Register(router)

	start(router, logger, conf)
}

func start(router *httprouter.Router, logger *logging.Logger, conf *config.Config) {
	logger.Info("Starting application")
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%s", conf.Http.Host, conf.Http.Port))
	if err != nil {
		logger.Fatal(err)
	}
	server := &http.Server{
		Handler:      router,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	logger.Info(fmt.Sprintf("Server is listening on %s:%s", conf.Http.Host, conf.Http.Port))
	logger.Fatal(server.Serve(listener))
}
