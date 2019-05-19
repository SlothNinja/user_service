package main

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
)

func addRoutes(prefix string, engine *gin.Engine, conf *oauth2.Config) {
	// User Group
	g1 := engine.Group(prefix)

	// New User
	g1.GET("/new",
		newAction(prefix),
	)

	// Create User
	g1.PUT("/new",
		create(prefix),
	)

	// Current User
	g1.GET("/current",
		current(prefix),
	)

	// User
	g1.GET("/json/:id",
		json(prefix),
	)

	// // Show User
	// g1.GET("show/:uid",
	// 	Show,
	// )

	// // Edit User
	// g1.GET("edit/:uid",
	// 	user.RequireLogin(),
	// 	Edit,
	// )

	// Update User
	g1.PUT("edit/:uid",
		update(prefix),
	)

	// // User Ratings
	// g1.POST("show/:uid/ratings/json",
	// 	rating.JSONIndexAction,
	// )

	// g1.POST("edit/:uid/ratings/json",
	// 	user.RequireLogin(),
	// 	rating.JSONIndexAction,
	// )

	// // User Games
	// g1.POST("show/:uid/games/json",
	// 	game.JSONIndexAction(gtype.All),
	// )

	// g1.POST("edit/:uid/games/json",
	// 	user.RequireLogin(),
	// 	game.JSONIndexAction(gtype.All),
	// )

	g1.GET("/login",
		Login(conf),
	)

	g1.GET("/logout",
		Logout,
	)

	// authHandler
	g1.GET("/auth", Auth(conf))

	// devauthHandler
	g1.POST("/auth", DevAuth)

	// // Users group
	// g2 := engine.Group(prefix + "s")

	// // Index
	// g2.GET("",
	// 	user.RequireAdmin(),
	// 	Index,
	// )

	// // json data for Index
	// g2.POST("/json",
	// 	user.RequireAdmin(),
	// 	JSON,
	// )
}
