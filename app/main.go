package main

import (
	"github.com/gin-gonic/gin"
	"log"
	netHttp "net/http"
	"rabbit-test/app/drivers"
	"rabbit-test/app/env"
)

func main() {
	env.Init()

	ginEngine := gin.New()

	dbConn := drivers.ConnectDB()
	defer func() {
		_ = dbConn.Close()
	}()
	drivers.DBMigration()
	_ = drivers.ConnectRedis()

	srv := &netHttp.Server{
		Addr:    ":" + env.AppPort,
		Handler: ginEngine,
	}

	err := srv.ListenAndServe()
	if err != netHttp.ErrServerClosed {
		log.Fatal(err)
	}
}
