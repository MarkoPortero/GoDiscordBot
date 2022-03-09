package DataAccess

import (
	"GoDiscordBot/GlobalVariables"
	"context"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"strconv"
	"strings"
	"time"
)

type CaptainLog struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	CaptainsLog string             `bson:"CaptainsLog"`
	AuthorID    string             `bson:"authorId"`
	Date        primitive.DateTime `bson:"$date"`
}

func MongoDbStoreCaptainsLogInDatabase(message *discordgo.MessageCreate) {
	client, err := mongo.NewClient(options.Client().ApplyURI(GlobalVariables.MongoUri))
	if err != nil {
		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer func(client *mongo.Client, ctx context.Context) {
		err := client.Disconnect(ctx)
		if err != nil {
			log.Fatal(err)
		}
	}(client, ctx)

	quickstartDatabase := client.Database("MyDatabase")
	captainsLogCollection := quickstartDatabase.Collection("CaptainsLog")

	breakDown := strings.Split(strings.ToLower(message.Content), "captainslog")

	captainLog := CaptainLog{
		CaptainsLog: breakDown[1],
		AuthorID:    message.Author.ID,
		Date:        primitive.NewDateTimeFromTime(time.Now()),
	}
	captainsLogResult, err := captainsLogCollection.InsertOne(ctx, captainLog)

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(captainsLogResult)

	fmt.Println()
}

func MongoDbReadCaptainsLogInDatabase(message *discordgo.MessageCreate) string {
	client, err := mongo.NewClient(options.Client().ApplyURI(GlobalVariables.MongoUri))
	if err != nil {
		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer func(client *mongo.Client, ctx context.Context) {
		err := client.Disconnect(ctx)
		if err != nil {
			log.Fatal(err)
			return
		}
	}(client, ctx)

	quickstartDatabase := client.Database("MyDatabase")
	captainsLogCollection := quickstartDatabase.Collection("CaptainsLog")

	var logs []CaptainLog
	var id string
	id = message.Author.ID
	cursor, err := captainsLogCollection.Find(ctx, bson.M{"authorId": id})
	if err != nil {
		log.Fatal(err)
	}

	if err = cursor.All(ctx, &logs); err != nil {
		panic(err)
	}

	fmt.Println(logs)

	var totalLogMessage string
	for i, s := range logs {
		totalLogMessage = totalLogMessage + " " + strconv.Itoa(i) + " " + s.CaptainsLog
	}
	return "```" + totalLogMessage + "```"
}
