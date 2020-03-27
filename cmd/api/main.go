package main

import (
	"context"
	"github.com/ckukadiya/go-gin-boilerplate/cmd/api/config"
	"github.com/ckukadiya/go-gin-boilerplate/cmd/api/router"
	"github.com/ckukadiya/go-gin-boilerplate/internal/database/mongodb"
	"github.com/ckukadiya/go-gin-boilerplate/internal/modules/friendship"
	"github.com/gin-gonic/gin"
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
	// TODO: Find out the way to add dynamic router
	router.NewFriendship(new(friendship.FriendshipController), v1Route)
	// router.NewPerson(person.New(db, cfg), v1Route)
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
