package main

import (
	"context"
	"fmt"
	"github.com/ckukadiya/go-gin-boilerplate/cmd/api/config"
	"github.com/ckukadiya/go-gin-boilerplate/cmd/api/router"
	"github.com/ckukadiya/go-gin-boilerplate/internal/database/mongodb"
	"github.com/ckukadiya/go-gin-boilerplate/internal/modules/person"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
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

	cfg, err := config.Load(mode)
	if err != nil {
		fmt.Println("error in getting configuration")
		fmt.Println(err)
	}
	initLog(cfg)
	r := gin.Default()

	// Create database connection
	db := mongodb.New(cfg)
	err = db.Ping(context.Background(), readpref.Primary())
	if err != nil {
		log.Fatal("Couldn't connect to the database", err)
	} else {
		log.Println("Connected!")
	}

	addV1Services(r, db, cfg)
	port := "8080"
	if cfg.Server.Port != 0 {
		port = strconv.Itoa(cfg.Server.Port)
	}

	r.Run(":" + port)
}

func addV1Services(r *gin.Engine, db *mongo.Client, cfg *config.Configuration) {
	v1Route := r.Group("/v1")
	// TODO: Find out the way to add dynamic router
	router.NewPerson(person.New(db, cfg), v1Route)
}

func initLog(cfg *config.Configuration) {
	// create logfile
	date := time.Now().Format("2006-01-02")
	n := cfg.Logger.Path + "/" + date + ".log"
	f, err := os.OpenFile(n, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0666)
	if err != nil {
		os.Exit(1)
	}
	log.SetOutput(f)

	// flags to log timestamp and lineno
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}
