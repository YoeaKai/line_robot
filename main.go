// This project implements 3 APIs to receive,
// reply and send messages to Line's official account.
// Use to CINNOX interview
// Made by Victor Chen

package main

import (
	"fmt"
	"log"

	"github.com/line/line-bot-sdk-go/v7/linebot"
	"github.com/spf13/viper"
)

var (
	bot                *linebot.Client // Line bot client instance
	address            string          // Server address
	channelSecret      string
	channelAccessToken string
	dbURI              string
	database           string
	dbCollection       string
)

func init() {
	// Set config.
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.SetDefault("application.port", 8080)

	if err := viper.ReadInConfig(); err != nil {
		log.Fatal("Failed to read in config: ", err)
	}

	address = fmt.Sprintf(":%s", viper.GetString("application.port"))
	channelSecret = viper.GetString("application.channelSecret")
	channelAccessToken = viper.GetString("application.channelAccessToken")
	dbURI = viper.GetString("application.dbURI")
	database = viper.GetString("application.database")
	dbCollection = viper.GetString("application.dbCollection")
}

func main() {
	var err error
	if bot, err = linebot.New(channelSecret, channelAccessToken); err != nil {
		log.Fatal("Failed to create a new bot client instance: ", err)
	}
}
