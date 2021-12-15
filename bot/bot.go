package main

import (
	"fmt"
	"log"

	slack "github.com/slack-go/slack"
)

var api *slack.Client

func startBot() {
	api = slack.New(appToken)

	sendMessage(testChannelId, "Hi there!\nThis is GoBot")

	att1 := getButtons()
	api.PostMessage(
		testChannelId,
		slack.MsgOptionAttachments(att1),
	)

}

func sendMessage(channel string, text string) {
	message := fmt.Sprintf(text)
	msgText := slack.NewTextBlockObject("mrkdwn", message, false, false)
	msgSection := slack.NewSectionBlock(msgText, nil, nil)

	msg := slack.MsgOptionBlocks(
		msgSection,
	)

	channelId, timestamp, something, err := api.SendMessage(channel, msg)

	if err != nil {
		log.Fatalf("%s\n", err)
	}

	log.Println("channelId:", channelId, "timestamp", timestamp, "something", something)
}

func getButtons() slack.Attachment {
	attachment := slack.Attachment{
		Title:      "Choice",
		Text:       "make your choice",
		CallbackID: "ACTION1",
		Fallback:   "Unable to choose",
		Color:      "#3AA3E3",
		Actions: []slack.AttachmentAction{
			slack.AttachmentAction{
				Name:  "Yes_Action",
				Text:  "Yes",
				Value: "YES",
				Type:  "button",
				Style: "",
				Confirm: &slack.ConfirmationField{
					OkText:      "Are you sure",
					DismissText: "Ditch it",
				},
			},
			slack.AttachmentAction{
				Name:  "No_Action",
				Text:  "No",
				Value: "NO",
				Type:  "button",
				Style: "",
			},
		},
	}

	return attachment
}
