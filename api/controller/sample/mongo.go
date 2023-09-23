package sample

import (
	"context"
	"github.com/gin-gonic/gin"
	"log"
	"pss/pkg/setting"
)

func MongoInsert(c *gin.Context) {

	collection := setting.MongoClient.Database("logs").Collection("logs")
	_, err := collection.InsertOne(context.TODO(), "hello")
	if err != nil {
		log.Println("error writing to mongo", err)
	}

}

func MongoGet(c *gin.Context) {

	mongoData := setting.MongoClient
	log.Println(mongoData)
	c.JSON(200, gin.H{
		"message": "success",
	})
}

func MongoUpdate(c *gin.Context) {

	c.JSON(200, gin.H{
		"message": "dbinsert",
	})
}

func MongoDelete(c *gin.Context) {

	c.JSON(200, gin.H{
		"message": "dbinsert",
	})
}
