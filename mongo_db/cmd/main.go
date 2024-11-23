package main

import (
	"context"
	"flag"
	"fmt"
	"log"

	//"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	//"go.mongodb.org/mongo-driver/mongo/readpref"
	rout "github.com/WinnersDunice/dunice_back/mongo_db/router"
)

// You will be using this Trainer type later in the program
type Trainer struct {
	Name string
	Age  int
	City string
}

func main() {
	// Rest of the code will go here
	key := flag.String("connect_url", "connect_url", "the key used to connxt to db")
	flag.Parse()

	client, err := mongo.NewClient(options.Client().ApplyURI(*key))
	if err != nil {
		log.Fatal(err)
	}

	// Create connect
	err = client.Connect(context.TODO())
	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}
    fmt.Println("Connected to MongoDB!")
	err = rout.Rout(client)
	if err != nil {
		log.Fatal("Error creating router: ", err)
	}
	fmt.Println("Connected to MongoDB!")
}
