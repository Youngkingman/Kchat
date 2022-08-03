package handler

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

// Handler struct holds required services for handler to function
type Handler struct {
	// TODO, like user service/ token servace/
	// core logic related to certain model(tables in the db) are injected here
}

// Config will hold services that will eventually be injected into this
// handler layer on handler initialization
type Config struct {
	R *gin.Engine
	// core logic related to certain model(tables in the db) are injected here
	// UserService     model.UserService
	// TokenService    model.TokenService
	BaseURL         string
	TimeoutDuration time.Duration
	MaxBodyBytes    int64
}

func NewHandler(c *Config) {
	// Create a handler (which will later have injected services)
	h := &Handler{
		// UserService:  c.UserService,
		// TokenService: c.TokenService,
		// MaxBodyBytes: c.MaxBodyBytes,
	} // currently has no properties
	g := c.R.Group(c.BaseURL)

	// logic handler here TODO
	if gin.Mode() != gin.TestMode {
		//g.Use(middleware.Timeout(c.TimeoutDuration, apperrors.NewServiceUnavailable()))
		//g.GET("/me", middleware.AuthUser(h.TokenService), h.Me)
	} else {
		//g.GET("/me", h.Me)
	}
	fmt.Print(g)
	fmt.Print(h)
}
