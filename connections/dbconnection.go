package connections

import (
	"context"
	"fmt"
	"log"
	"github.com/Mohammad-Hakemi22/mongoAPI/models"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const connectionString = "localhost:27017"
const dbName = "Netflix"
const collectionName = "Movies"

var collection *mongo.Collection

// Connection with mongoDB
func init() {
	// client opions
	clientOption := options.Client().ApplyURI(connectionString)

	// connect to mongo
	client, err := mongo.Connect(context.TODO(), clientOption)
	checkError(err)

	fmt.Println("Database Connection: OK")

	collection = client.Database(dbName).Collection(collectionName)
	fmt.Println("Collection instance: OK")

}

// mongoDB helper - insert
func insertOneMovie(movie models.Netflix) {
	inserted, err := collection.InsertOne(context.Background(), movie)
	checkError(err)
	fmt.Println("inserted movie: OK; movieID: ", inserted.InsertedID)
}

func checkError(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}
