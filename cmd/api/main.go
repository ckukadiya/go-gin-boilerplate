// Go Gin Boilerplate API Demo.
//
// The purpose of this application is to provide boilerplate
// of REST full API using MongoDB
//
// Terms Of Service:
//
// there are no TOS at this moment, use at your own risk we take no responsibility
//
//     Schemes: http, https
//     Host: localhost:8080
//     BasePath: /v1
//     Version: 0.0.1
//     License: MIT http://opensource.org/licenses/MIT
//     Contact: Chintan Kukadiya<contact.chintankukadiya@gmail.com>
//
//     Consumes:
//     - application/json
//
//     Produces:
//     - application/json
//
// swagger:meta
package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"go-gin-boilerplate/cmd/api/config"
	"go-gin-boilerplate/cmd/api/router"
	"go-gin-boilerplate/internal/database/mongodb"
	"go-gin-boilerplate/internal/modules/person"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
	"os"
	"strconv"
	"time"
)

func main() {
	// Getting configuration base on environment
	mode := os.Getenv("MODE")
	if mode == "" {
		mode = "dev"
	}

	if mode == "prod" {
		gin.SetMode(gin.ReleaseMode)
	}

	config.Load(mode)
	initLog()
	r := gin.Default()

	// Create database connection
	mongodb.NewMongoDB()
	err := mongodb.GetMongoDBClient().Ping(context.Background(), readpref.Primary())
	if err != nil {
		log.Fatal("Couldn't connect to the database", err)
	}

	addV1Services(r)
	port := "8080"
	cfg := config.GetConfig()
	if cfg.Server.Port != 0 {
		port = strconv.Itoa(cfg.Server.Port)
	}

	r.Run(":" + port)
}

func addV1Services(r *gin.Engine) {
	v1Route := r.Group("/v1")
	r.Static("/swagger/", "cmd/api/docs")
	// TODO: Find out the way to add dynamic router
	router.NewPerson(new(person.PersonController), v1Route)
}

func initLog() {
	// create logfile
	date := time.Now().Format("2006-01-02")
	cfg := config.GetConfig()

	n := cfg.Logger.Path + "/" + date + ".log"
	f, err := os.OpenFile(n, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0666)
	if err != nil {
		os.Exit(1)
	}
	log.SetOutput(f)

	// flags to log timestamp and lineno
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}
