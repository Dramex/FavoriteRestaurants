package routes

import 	(
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	mongo "github.com/Dramex/FavoriteRestaurants/lib/mongo"
	"fmt"
)

type DATA struct{
    USERNAME string
    PASSWORD string
	FIRSTNAME string
	LASTNAME string
	EMAIL string
}

func SignUp(c *gin.Context) {
	var data DATA
	c.BindJSON(&data)
	//var userDetails bson.M 
	userResult, err := mongo.Users.InsertOne(mongo.Ctx, bson.M{"username": data.USERNAME});
	if err == nil {
		//userResult.Decode(&userDetails)
		fmt.Println("USER", userResult);
		c.JSON(200, gin.H{"status": data.USERNAME}) 
	}

}