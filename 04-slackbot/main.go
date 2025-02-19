package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/shomali11/slacker"
)

func printCommandEvents(analyticsChannel <-chan *slacker.CommandEvent) {

	for event := range analyticsChannel {
		fmt.Println("Command Events")
		fmt.Println("Timestamp: ", event.Timestamp)
		fmt.Println("Command: ", event.Command)
		fmt.Println("Parameters: ", event.Parameters)
		fmt.Println("Channel: ", event.Event)
	}
}

func main() {
	fmt.Println("Loading slackbot...")
	os.Setenv("SLACK_BOT_TOKEN", "GET_YOUR_OWN")
	os.Setenv("SLACK_APP_TOKEN", "GET_YOUR_OWN")

	bot := slacker.NewClient(os.Getenv("SLACK_BOT_TOKEN"), os.Getenv("SLACK_APP_TOKEN"))

	go printCommandEvents(bot.CommandEvents())

	bot.Command("My YOB is <year>", &slacker.CommandDefinition{
		Description: "Your Age Calculator",
		// Example:      "Your Year of Birth 1990",
		Handler: func(botCtx slacker.BotContext, request slacker.Request, response slacker.ResponseWriter) {
			year := request.Param("year")
			yob, err := strconv.Atoi(year)
			if err != nil {
				response.Reply("Please enter a valid year")
				return
			}

			response.Reply(fmt.Sprintf("You are %d years old", 2025-yob))
		},
	})

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	err := bot.Listen(ctx)
	if err != nil {
		log.Fatal(err)
	}

}
