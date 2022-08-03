package main

import (
	"fmt"
	"log"
	"time"

	"github.com/Youngkingman/Kchat/kchat/internal/handler"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"github.com/jmoiron/sqlx"
)

type dataSources struct {
	DB          *sqlx.DB
	RedisClient *redis.Client
	//TODO the files and picture storage
}

func initDS() (ds *dataSources, err error) {
	//TODO
	// Environment Setting

	// link MySql database

	// verify the accessability of MySql database

	// link redis

	// verify the accessability of MySql database

	// others
	return
}

func (d *dataSources) closeDS() error {
	if err := d.DB.Close(); err != nil {
		return fmt.Errorf("error closing MySql: %w", err)
	}
	if err := d.RedisClient.Close(); err != nil {
		return fmt.Errorf("error closing Redis Client: %w", err)
	}
	//TODO
	return nil
}

func initEngine(d *dataSources) (*gin.Engine, error) {
	log.Println("Injecting data sources")

	// repository layer

	// services layer

	// rsa keys for jwt

	// tokenService

	// initialize gin.Engine
	router := gin.Default()

	baseURL := "" //TODO here, replace with config

	handler.NewHandler(&handler.Config{
		R: router,
		//TODO
		BaseURL:         baseURL,
		TimeoutDuration: time.Duration(time.Duration(1) * time.Second), // 1 should be replace
		MaxBodyBytes:    0,                                             //TODO
	})
	return router, nil
}
