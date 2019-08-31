package main

import (
	"bitbucket.org/SlothNinja/user"
	"github.com/gin-gonic/gin"
)

const (
	authPath = "/auth"
)

func addRoutes(host, port, prefix string, secure bool, engine *gin.Engine) {
	// User Group
	g1 := engine.Group(prefix)

	// New User
	g1.GET("/new", newAction(prefix))

	// Create User
	g1.PUT("/new", create(prefix))

	// Current User
	g1.GET("/current", current)

	// User
	g1.GET("/json/:id", json(prefix))

	// Update User
	g1.PUT("edit/:uid", update(prefix))

	g1.GET("/login", user.Login(authPath))

	// authHandler
	g1.GET("/auth", user.Auth(authPath))
}
