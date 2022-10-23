// This project implements 3 APIs to receive,
// reply and send messages to Line's official account.
// Use to CINNOX interview
// Made by Victor Chen

package main

import (
	"fmt"
	"log"

	db "github.com/YoeaKai/line_robot/db"

	"github.com/gin-gonic/gin"
	"github.com/line/line-bot-sdk-go/v7/linebot"
	"github.com/spf13/viper"
)

var (
	bot                *linebot.Client // Line bot client instance
	address            string          // Server address
	channelSecret      string
	channelAccessToken string
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
}

func main() {
	server := gin.Default()

	var err error
	if bot, err = linebot.New(channelSecret, channelAccessToken); err != nil {
		log.Fatal("Failed to create a new bot client instance: ", err)
	}

	// Receive the user's message, save it into the database, and reply to the user.
	server.POST("/message", receiveAndSaveMessage)

	// Reply message to the user.
	server.POST("/reply", replyMessage)

	// Get the user list from the database.
	server.GET("/user/list", getUserListFromDB)

	server.Run(address)
}

// receiveAndSaveMessage receives the user's message, saves it into the database,
// and finally replies with a message to the user.
func receiveAndSaveMessage(ctx *gin.Context) {
	// Parse and check the request.
	events, err := bot.ParseRequest(ctx.Request)
	if err != nil {
		if err == linebot.ErrInvalidSignature {
			ctx.Writer.WriteHeader(400)
		} else {
			ctx.Writer.WriteHeader(500)
		}
		return
	}

	// Analyze each event.
	for _, event := range events {
		if event.Type == linebot.EventTypeMessage {
			var text, messageId string
			switch message := event.Message.(type) {
			// Handle text messages.
			case *linebot.TextMessage:
				text = message.Text
				messageId = message.ID
				replyMessage := linebot.NewTextMessage("Get: " + text + ", \nThank you for your using!")

				// Reply message to the user.
				if _, err := bot.ReplyMessage(event.ReplyToken, replyMessage).Do(); err != nil {
					log.Printf("Failed to reply message for %s: %v", messageId, err)
				}
				// Handle sticker messages.
			case *linebot.StickerMessage:
				var keywords string
				for _, keyword := range message.Keywords {
					keywords = fmt.Sprint(keywords, ", ", keyword)
				}
				keywords = keywords[2:]

				replyMessage := linebot.NewTextMessage("Get your cute stickcer! \nIts keyword: " + keywords)

				if _, err = bot.ReplyMessage(event.ReplyToken, replyMessage).Do(); err != nil {
					log.Print(err)
				}
			}

			// Set the document and save it into the database.
			document := db.UserMessage{
				UserId:    event.Source.UserID,
				Timestamp: event.Timestamp,
				Message: db.Message{
					MessageType: string(event.Message.Type()),
					Text:        text,
				},
			}

			if err := db.InsertMessageToDB(ctx, document, messageId); err != nil {
				log.Printf("Failed to insert message for %s to database: %v", messageId, err)
			}
		}
	}
}

// replyMessage replies a message to the event determined by replyToken.
func replyMessage(ctx *gin.Context) {
	replyToken := ctx.PostForm("replyToken")
	messageText := ctx.PostForm("messageText")
	message := linebot.NewTextMessage("Get: \n" + messageText + ", \nThank you for your using!")

	if _, err := bot.ReplyMessage(replyToken, message).Do(); err != nil {
		log.Println("Failed to reply message: ", err)
	} else {
		ctx.JSON(200, gin.H{
			"message": messageText,
		})
	}
}

// getUserListFromDB gets all user IDs from the collection which is set from the config file.
func getUserListFromDB(ctx *gin.Context) {
	var results []interface{}
	var err error

	if results, err = db.GetUserList(ctx); err != nil {
		log.Println("Failed to get user id from database: ", err)
	}

	ctx.JSON(200, gin.H{
		"userList": results,
	})
}
