package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/nlopes/slack"
	"github.com/spf13/viper"
)

func main() {
	service := Service{}
	service.init()

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

type Service struct {
	user         string
	pass         string
	weatherToken string
}

func (s *Service) init() {
	viper.SetEnvPrefix("OWL")
	viper.AutomaticEnv()
	s.user = os.Getenv("WIKI_USER")
	s.pass = os.Getenv("WIKI_PASS")
	s.weatherToken = viper.GetString("WEATHER_TOKEN")

	s.loadIntents()
}

func (s *Service) HandleMsg(s string) {
	s.respond(s.lookup())
}

func (s *Service) respond(rtm *slack.RTM, msg *slack.MessageEvent) {
	var response string
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
	text := msg.Text
	text = strings.TrimSpace(text)
	text = strings.ToLower(text)

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
		wiki(s.user, wiki_pass, message)
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
