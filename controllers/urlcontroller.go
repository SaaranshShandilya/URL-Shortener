package controllers

import (
	"fmt"

	"github.com/SaaranshShandilya/URLShortner/models"
	// "fmt"

	"context"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"log"

	// "encoding/json"

	"go.mongodb.org/mongo-driver/bson"
	"net/http"

	"math/rand"
)


var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func RandStringRunes(n int) string {
    b := make([]rune, n)
    for i := range b {
        b[i] = letterRunes[rand.Intn(len(letterRunes))]
    }
    return string(b)
}

func Testandsave(c *gin.Context) {
	var body struct{
		URL string
	}

	c.Bind(&body)

	random:=RandStringRunes(10)

	client:=c.MustGet("db").(*mongo.Client)
	usersCollection := client.Database("testing").Collection("urls")

	url := models.Table{
		FullUrl:body.URL,
		ShortUrl: random,
	}

	result, err := usersCollection.InsertOne(context.TODO(), url)

	if err != nil{
		fmt.Println(result.InsertedID)
		return 
	}

	c.JSON(200,gin.H{
		"table":url,
	})
}

func Route(c *gin.Context){
	url_name := c.Param("url")

	fmt.Println(url_name)
	filter:=bson.D{{"shorturl",url_name}}
	client:=c.MustGet("db").(*mongo.Client)
	usersCollection := client.Database("testing").Collection("urls")

	var result models.Table
	err := usersCollection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
    // ErrNoDocuments means that the filter did not match any documents in the collection
    if err == mongo.ErrNoDocuments {
		log.Println("NO doc found")
        return
    }
    	log.Fatal(err)
	}

	fmt.Println(result)

	c.Redirect(http.StatusMovedPermanently,result.FullUrl)

	
	c.JSON(200,gin.H{
		"url":result,
	})


}