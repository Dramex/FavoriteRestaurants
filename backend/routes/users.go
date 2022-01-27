package routes

import 	(
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	mongo "github.com/Dramex/FavoriteRestaurants/lib/mongo"
	"fmt"
	"crypto/md5"
    "encoding/hex"
	"github.com/golang-jwt/jwt"

)
var hmacSampleSecret []byte

type REGISTER_DATA struct{
	EMAIL string
    PASSWORD string
	FIRSTNAME string
	LASTNAME string
}
type LOGIN_DATA struct{
	EMAIL string
    PASSWORD string
}

func SignUp(c *gin.Context) {
	var data REGISTER_DATA
	c.BindJSON(&data)
	fmt.Println("USER DATA", data);
	//var userDetails bson.M 
	h := md5.New()
    h.Write([]byte(data.PASSWORD))
	userResult, err := mongo.Users.InsertOne(mongo.Ctx, bson.M{
		"email": data.EMAIL,
		"password": hex.EncodeToString(h.Sum(nil)),
		"firstname": data.FIRSTNAME,
		"lastname": data.LASTNAME,
	});

	if err == nil {
		//userResult.Decode(&userDetails)
		fmt.Println("USER db", userResult);
		c.JSON(200, gin.H{"status": data.EMAIL}) 
	}

}



func SignIn(c *gin.Context) {
	var data LOGIN_DATA
	c.BindJSON(&data)

	// hash pass
	h := md5.New()
    h.Write([]byte(data.PASSWORD))

	fmt.Println("LOGIN DATA", bson.M{
		"email": data.EMAIL,
		"password": hex.EncodeToString(h.Sum(nil)),
	});
	var user bson.M
	mongo.Users.FindOne(mongo.Ctx, bson.M{
		"email": data.EMAIL,
		"password": hex.EncodeToString(h.Sum(nil)),
	}).Decode(&user);

	
	// JSONData := struct {
	// 	Path string `json:"Path"`
	// }{}
	//userResult.Decode(&JSONData);
	fmt.Println("decode", user)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id": user["_id"],
		"email": user["email"],
		"firstname":user["firstname"],
		"lastname:":user["lastname"],
	})

		// Sign and get the complete encoded token as a string using the secret
		tokenString, err := token.SignedString(hmacSampleSecret)
		
		fmt.Println(tokenString, err)
		c.JSON(200, gin.H{"token": tokenString}) 
}