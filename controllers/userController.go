package controllers

import (
	"github.com/Suman196pokhrel/go-jwt-auth/database"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/mongo"
)

var userCollection *mongo.Collection = database.OpenCollection(database.Client, "user")
var validate = validator.New()

func HashPassword() {

}

func VerifyPassword() {

}

func Signup(c *gin.Context) {

}

func Login(c *gin.Context) {

}

func GetUsers(c *gin.Context) {

}

func GetUser(c *gin.Context) {

}
