package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Request struct {
	Nombre  string
	Fecha   time.Time
	Monto   int
	Plazo   time.Time
	Empresa string
}

func main() {
	r := gin.Default()

	serverAPIOptions := options.ServerAPI(options.ServerAPIVersion1)
	clientOptions := options.Client().
		ApplyURI("mongodb+srv://admin:w41I0DHA3EkkYY8V@sprint4.jcenge9.mongodb.net/?retryWrites=true&w=majority").
		SetServerAPIOptions(serverAPIOptions)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	r.PUT("/requests", func(c *gin.Context) {
		var request Request
		c.BindJSON(&request)

		c.JSON(200, request)
		fmt.Println("Success parsing Request! Sending to DB...")
		cl := client.Database("requests").Collection("requests")
		cl.InsertOne(ctx, request)
		fmt.Println("Sent to mongo!")
	})

	r.Run("0.0.0.0:8080")
}
