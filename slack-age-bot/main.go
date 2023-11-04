package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/shomali11/slacker"
)

// Use analyticChannel to receive commans events from the Slack bot.
func printCommandEvent(analyticChannel <-chan *slacker.CommandEvent) {
	for event := range analyticChannel {
		fmt.Println("Command Events")
		fmt.Println(event.Timestamp)
		fmt.Println(event.Command)
		fmt.Println(event.Event)
		fmt.Println()
	}

}

func main() {
	// Set environment variables
	os.Setenv("SLACK_BOT_TOKEN", "xoxb-6101182274725-6142744918517-IQSEfEO5Mt4BXnKmpBfI2Msv")
	os.Setenv("SLACK_APP_TOKEN", "xapp-1-A063V3HA1HD-6139013502054-34e0e9cb687dc81173a58b851cda7826addd09019827e57c044d3f5c32f09ec1")

	bot := slacker.NewClient(os.Getenv("SLACK_BOT_TOKEN"), os.Getenv("SLACK_APP_TOKEN"))
	go printCommandEvent(bot.CommandEvents())
	bot.Command("my year of birth is <year>", &slacker.CommandDefinition{
		Description: "yob calculator",
		Handler: func(botCtx slacker.BotContext, request slacker.Request, response slacker.ResponseWriter) {
			year := request.Param("year")
			yob, err := strconv.Atoi(year)
			if err != nil {
				response.Reply("Error")
				return
			}
			age := 2023 - yob
			r := fmt.Sprintf("age is %d", age)
			response.Reply(r)
		},
	})

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	err := bot.Listen(ctx)
	if err != nil {
		log.Fatal(err)
	}
}
