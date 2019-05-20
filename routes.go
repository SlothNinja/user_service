package main

import (
	"github.com/gin-gonic/gin"
)

func addRoutes(prefix string, engine *gin.Engine) {
	// User Group
	g1 := engine.Group(prefix)

	// New User
	g1.GET("/new", newAction(prefix))

	// Create User
	g1.PUT("/new", create(prefix))

	// Current User
	g1.GET("/current", current(prefix))

	// User
	g1.GET("/json/:id", json(prefix))

	// Update User
	g1.PUT("edit/:uid", update(prefix))

	g1.GET("/login", Login)

	g1.GET("/logout", Logout)

	// authHandler
	g1.GET("/auth", Auth)

	// devauthHandler
	g1.POST("/auth", DevAuth)
}
