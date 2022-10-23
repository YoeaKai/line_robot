// db_operation package

package db_operation

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type message struct {
	MessageType string
	Text        string
}

// userMessage is the type saved in MongoDB.
type userMessage struct {
	UserId    string
	Timestamp time.Time
	Message   message
}

var (
	dbURI        string
	database     string
	dbCollection string
)

func init() {
	// Set config.
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("../")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatal("Failed to read in config: ", err)
	}

	dbURI = viper.GetString("application.dbURI")
	database = viper.GetString("application.database")
	dbCollection = viper.GetString("application.dbCollection")
}

// insertMessageToDB insert "document" into the collection which is set from the config file.
// messageId is the message ID get from each event.
func InsertMessageToDB(ctx context.Context, document userMessage, messageId string) error {
	client, err := ConnectToDatabase(ctx)
	if err != nil {
		return fmt.Errorf("failed to connect to database for %s: %v", messageId, err)
	}
	defer client.Disconnect(ctx)

	collection := client.Database(database).Collection(dbCollection)

	insertOneResult, err := collection.InsertOne(ctx, document)
	if err != nil {
		return fmt.Errorf("failed to insert message id %s: %v", messageId, err)
	}

	log.Printf("Success insert message id %s in inserted id: %s", messageId, insertOneResult.InsertedID)

	return nil
}

// connectToDatabase connects to the database and
// verifies whether the connection is available or not.
func ConnectToDatabase(ctx context.Context) (*mongo.Client, error) {
	// Connect to database.
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(dbURI))
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %v", err)
	}

	// Verify whether the connection is available or not.
	if err := client.Ping(ctx, nil); err != nil {
		log.Fatal("Failed to verify database is available: ", err)
	}

	log.Println("Success connect to database")

	return client, nil
}
