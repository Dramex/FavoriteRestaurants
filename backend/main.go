package main

import (
	"github.com/gin-gonic/gin"
	Routes "github.com/Dramex/FavoriteRestaurants/routes"
	mongo "github.com/Dramex/FavoriteRestaurants/lib/mongo"
)

func main() {
	r := gin.Default()
	r.POST("/signup", Routes.SignUp)
	mongo.Mongo() // connect to database :-)

	r.Run(":8989") // listen and serve on 0.0.0.0:8080
}