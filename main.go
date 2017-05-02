package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/nlopes/slack"
	"github.com/spf13/viper"
)

func main() {
	viper.SetEnvPrefix("OWL")
	viper.AutomaticEnv()

	slackToken := viper.GetString("SLACK_TOKEN")

	debug := viper.GetBool("DEBUG")

	api := slack.New(slackToken)
	api.SetDebug(debug)

	rtm := api.NewRTM()
	go rtm.ManageConnection()
	fmt.Println(slackToken)

Loop:
	for {
		select {
		case msg := <-rtm.IncomingEvents:
			fmt.Print("Event Received: ")
			switch ev := msg.Data.(type) {
			case *slack.ConnectedEvent:
				fmt.Println("Connection counter:", ev.ConnectionCount)

			case *slack.MessageEvent:
				fmt.Printf("Message: %v\n", ev)

				respond(rtm, ev)

			case *slack.RTMError:
				fmt.Printf("Error: %s\n", ev.Error())

			case *slack.InvalidAuthEvent:
				fmt.Printf("Invalid credentials")
				break Loop

			default:
				//Take no action
			}
		}
	}
}

func respond(rtm *slack.RTM, msg *slack.MessageEvent) {

	// TODO: Figure out better way to pass these environment variables around
	viper.SetEnvPrefix("OWL")
	viper.AutomaticEnv()
	wiki_user := os.Getenv("WIKI_USER")
	wiki_pass := os.Getenv("WIKI_PASS")
	weatherToken := viper.GetString("WEATHER_TOKEN")

	var response string
	text := msg.Text
	text = strings.TrimSpace(text)
	text = strings.ToLower(text)

	acceptedGreetings := map[string]bool{
		"what's up?": true,
		"hey!":       true,
		"hello":      true,
		"yo":         true,
	}
	acceptedHowAreYou := map[string]bool{
		"how's it going?": true,
		"how are ya?":     true,
		"feeling okay?":   true,
	}

	if acceptedGreetings[text] {
		response = "What's up buddy!?!?!"
		rtm.SendMessage(rtm.NewOutgoingMessage(response, msg.Channel))
	} else if acceptedHowAreYou[text] {
		response = "Good. How are you?"
		rtm.SendMessage(rtm.NewOutgoingMessage(response, msg.Channel))
	} else if strings.HasPrefix(text, ":construction:") {
		response = "Sending message to wiki!"
		rtm.SendMessage(rtm.NewOutgoingMessage(response, msg.Channel))
		message := fmt.Sprintln("<p>", msg.Text, "</p>")
		wiki(wiki_user, wiki_pass, message)
	} else if strings.HasPrefix(text, "weather") {
		s := strings.Split(text, " ")

		if len(s) == 2 {
			zipCode := s[1]
			summary, temperature := getWeather(weatherToken, zipCode)
			response := fmt.Sprintf("It is %s with a temperature of %.f°F\n", strings.ToLower(summary), temperature)
			rtm.SendMessage(rtm.NewOutgoingMessage(response, msg.Channel))
		} else {
			response := fmt.Sprintln("Usage: `weather zipcode`")
			rtm.SendMessage(rtm.NewOutgoingMessage(response, msg.Channel))
		}

	}

}
