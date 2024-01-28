package main

import (
	"fmt"
	"github.com/SaaranshShandilya/URLShortner/initialise"
	"go.mongodb.org/mongo-driver/mongo"
	"github.com/gin-gonic/gin"
	"github.com/SaaranshShandilya/URLShortner/controllers"

)


var Client * mongo.Client
// var UsersCollection * mongo.Collection

func init(){
	Client  = initialise.ConnectToDb()
	// UsersCollection = Client.Database("testing").Collection("users")
	// fmt.Println(UsersCollection)
}

func main(){

	r := gin.Default()
	r.Use(func(c *gin.Context) {
		fmt.Println()
        c.Set("db", Client)
        c.Next()
    })

	r.POST("/shorten",controllers.Testandsave)
	r.GET("/:url", controllers.Route)
	


	r.Run() 

}