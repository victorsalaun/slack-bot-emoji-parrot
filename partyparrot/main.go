package partyparrot

import (
	"encoding/json"
	"github.com/nlopes/slack"
	"log"
	"io/ioutil"
)

type Token struct {
	Token string `json:"token"`
}

var (
	api    *slack.Client
	botKey Token
)

func init() {
	file, err := ioutil.ReadFile("../token.json")

	if err != nil {
		log.Fatal("File doesn't exist")
	}

	if err := json.Unmarshal(file, &botKey); err != nil {
		log.Fatal("Cannot parse token.json")
	}
}

func repeat() {
	for {
		ac := <-botReplyChannel
		params := slack.PostMessageParameters{}
		params.AsUser = true
		params.Attachments = []slack.Attachment{*ac.Attachment}
		_, _, errPostMessage := api.PostMessage(ac.Channel.Name, ac.DisplayTitle, params)
		if errPostMessage != nil {
			log.Fatal(errPostMessage)
		}
	}
}

func main() {
	api = slack.New(botKey.Token)

	rtm := api.NewRTM()

	go rtm.ManageConnection()
	go repeat()

	for {
		select {
		case msg := <-rtm.IncomingEvents:
			switch ev := msg.Data.(type) {
			case *slack.MessageEvent:
				break
			}
		}

	}
}
