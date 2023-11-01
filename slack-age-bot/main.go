package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/shomali11/slacker"
)

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
	os.Setenv("SLACK_BOT_TOKEN", "xoxb-6101182274725-6125533944708-bn3Tw255YZROGrhxMlmVewlo")
	os.Setenv("SLACK_APP_TOKEN", "xapp-1-A06323ET118-6116761573425-1cf2b081e0f06046f59ef734e85412ed1c084001ec2d669cd0f31f9c91c55bbe")

	bot := slacker.NewClient(os.Getenv("SLACK_BOT_TOKEN"), os.Getenv("SLACK_APP_TOKEN"))
	go printCommandEvent(bot.CommandEvents())

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	err := bot.Listen(ctx)
	if err != nil {
		log.Fatal(err)
	}
}
