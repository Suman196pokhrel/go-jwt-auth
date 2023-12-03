package controllers

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/Suman196pokhrel/go-jwt-auth/database"
	"github.com/Suman196pokhrel/go-jwt-auth/helpers"
	"github.com/Suman196pokhrel/go-jwt-auth/models"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var userCollection *mongo.Collection = database.OpenCollection(database.Client, "user")
var validate = validator.New()

func HashPassword() {

}

func VerifyPassword() {

}

func Signup() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var c, cancle = context.WithTimeout(context.Background(), 100*time.Second)
		var user models.User

		if err := ctx.BindJSON(&user); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		validationErr := validate.Struct(user)
		if validationErr != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error()})
			return
		}

		// GET ACCOUNT WITH THIS EMAIL IF ALREADY EXISTS
		countEmail, err := userCollection.CountDocuments(c, bson.M{"email": user.Email})
		defer cancle()
		if err != nil {
			log.Panic(err)
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		// GET ACCOUNT WITH THIS Phone NUM IF ALREADY EXISTS
		countPhone, err := userCollection.CountDocuments(c, bson.M{"phone": user.Phone})
		defer cancle()
		if err != nil {
			log.Panic(err)
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		if countEmail > 0 || countPhone > 0 {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Account with this email or phone number already esists"})
			return
		}

		// Adding data into our User instance to later insert the whole entity as a row into the database
		currentTime := time.Now()
		user.Created_at = currentTime
		user.Updated_at = currentTime
		user.ID = primitive.NewObjectID()
		user.User_id = user.ID.Hex()

		// GETTING TOKEN, REFRESH_TOKEN

	}
}

func Login(c *gin.Context) {

}

func GetUsers(c *gin.Context) {

}

func GetUser() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userId := ctx.Param("user_id")

		if err := helpers.MatchUserTypeToUid(ctx, userId); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		var c, cancle = context.WithTimeout(context.Background(), 100*time.Second)
		var user models.User
		err := userCollection.FindOne(c, bson.M{"user_id": userId}).Decode(&user)
		defer cancle()

		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, user)
	}

}
