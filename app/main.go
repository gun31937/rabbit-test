package main

import (
	"github.com/gin-gonic/gin"
	"log"
	netHttp "net/http"
	"rabbit-test/app/drivers"
	"rabbit-test/app/env"
	"rabbit-test/app/handlers/http"
	"rabbit-test/app/repositories"
	"rabbit-test/app/usecases"
)

func main() {
	env.Init()

	ginEngine := gin.New()

	dbConn := drivers.ConnectDB()
	defer func() {
		_ = dbConn.Close()
	}()
	drivers.DBMigration()

	rdbConn := drivers.ConnectRedis()

	http.NewRouterHealth(ginEngine, dbConn, rdbConn)

	databaseRepo := repositories.InitDatabase(dbConn)
	redisRepo := repositories.InitRedis(rdbConn)

	shortURLUseCase := usecases.InitShortURL(databaseRepo, redisRepo)
	http.NewRouterShortURL(ginEngine, shortURLUseCase)
	http.NewRouterAdmin(ginEngine)

	srv := &netHttp.Server{
		Addr:    ":" + env.AppPort,
		Handler: ginEngine,
	}

	err := srv.ListenAndServe()
	if err != netHttp.ErrServerClosed {
		log.Fatal(err)
	}
}
