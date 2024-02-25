package initializers

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectDatabase() (*mongo.Client, error) {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}
	uri := os.Getenv("MONGODB_URI")

	log.Println(uri)

	if uri == "" {
		log.Fatal("You must set your 'MONGODB_URI' environment variable. See\n\t https://www.mongodb.com/docs/drivers/go/current/usage-examples/#environment-variable")
	}

	clientOptions := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(context.Background(), clientOptions)

	if err != nil {
		log.Println(err)

		return nil, err
	}

	// Ping the MongoDB server to ensure connectivity
	err = client.Ping(context.Background(), nil)
	if err != nil {
		return nil, err
	}

	fmt.Println("Connected to MongoDB!")

	return client, nil

}

func InsertAMovie(client *mongo.Client) {
	coll := client.Database("test").Collection("movies")

	type Movie struct {
		Name string
		Hour int
	}

	doc := Movie{Name: "Kong vs Godzilla", Hour: 1}

	result, err := coll.InsertOne(context.TODO(), doc)

	// result, err := coll.InsertOne(
	// 	context.TODO(),
	// 	bson.D{
	// 		{"animal", "Dog"},
	// 		{"bread", "Beagle"},
	// 	},
	// )

	if err != nil {
		log.Fatal("INSERT ERR: ", err)
		return
	}

	log.Println(result)

}
