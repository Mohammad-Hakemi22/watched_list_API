package connections

import (
	"context"
	"fmt"
	"log"

	"github.com/Mohammad-Hakemi22/mongoAPI/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

func updateOneMovie(movieID string) {
	id, err := primitive.ObjectIDFromHex(movieID)
	checkError(err)
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"whatched": true}}
	result, err := collection.UpdateOne(context.Background(), filter, update)
	checkError(err)
	fmt.Println("updated one movie: OK; count: ",result.ModifiedCount)
}

func deleteOneMovie(movieID string) {
	id, err := primitive.ObjectIDFromHex(movieID)
	checkError(err)
	filter := bson.M{"_id": id}
	result, err := collection.DeleteOne(context.Background(), filter)
	checkError(err)
	fmt.Println("deleted one movie: OK; count: ", result.DeletedCount)
}

func deleteAllMovies() {
	result, err := collection.DeleteMany(context.Background(), bson.D{{}})
	checkError(err)
	fmt.Println("deleted all movies: OK; count: ", result.DeletedCount)
}

func GetAllMovies_helper() []bson.M {
	cursor, err := collection.Find(context.Background(), bson.D{{}})
	defer cursor.Close(context.Background())
	checkError(err)
	// var movies []primitive.M
	var movies []bson.M

	for cursor.Next(context.Background()) {
		var movie bson.M
		err := cursor.Decode(&movie)
		checkError(err)
		movies = append(movies, movie)
	}
	return movies
}

func checkError(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}
