package database

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
)

func DBinstance() *mongo.Client{

	err :=godotenv.Load(".env")
	if err!=nil{
		log.Fatal("Error load .env file")
	}

}