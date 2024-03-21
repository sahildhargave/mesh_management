package helpers

import (
	"context"
	"log"
	"os"
	"restruant-management/database"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/mongo"
)

//😁😂😎😍😘🥰
//Use jwt.StandardClaims:
//You can create a JWT token with standard claims such as
//expiration time, issuing time, and issuer using
//jwt.StandardClaims.
type SignedDetails struct{
  Email         string 
  First_name    string 
  Last_name     string
  Uid           string
  jwt.StandardClaims
}


var userCollection *mongo.Collection= database.OpenCollection(database.Client,"user")

var SECRET_KEY string = os.Getenv("SECRET_KEY")

func GeneraterAllTokens(email string , firstName string , lastName string, uid string)(signedToken string , signedrefreshToken string , err error){
	claims:= &SignedDetails{
		Email      :     email,
		First_name :     firstName,
		Last_name  :     lastName,
		Uid        :     uid,
		StandardClaims:  jwt.StandardClaims{
			ExpiresAt : time.Now().Local().Add(time.Hour * time.Duration(24)).Unix(),


		},
	}

	refreshClaims := &SignedDetails{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt : time.Now().Local().Add(time.Hour * time.Duration(168)).Unix(),
		},
	}
	
	// 
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(SECRET_KEY))
    refreshToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims).SignedString([]byte(SECRET_KEY))


	if err != nil{
		log.Panic(err)
		return
	}
	return token, refreshToken ,err
}

func UpdateAllTokens(signedToken string , signedrefreshToken string , userId string){
	var context,cancel = context.WithTimeout(context.Background(), 100 *time.Second)
    
	var updateObj primitive.D
    
	up

}