package helper

import (
	"os"

	"github.com/slack-go/slack"
)

func SendSlackNotification(message string) (string, error) {

	channelID := os.Getenv("AUB__CHANNEL_ID")
	slackToken := os.Getenv("AUB__SLACK_TOKEN")
	client := slack.New(slackToken, slack.OptionDebug(false))

	attachment := slack.Attachment{
		Pretext: "Message from Aubserver",
		Text:    "List of Failing pods",
		// Color Styles the Text, making it possible to have like Warnings etc.
		Color: "#ff0000",
		// Fields are Optional extra data!
		Fields: []slack.AttachmentField{
			{
				Title: "Failed Pods",
				Value: message,
			},
		},
	}
	_, timestamp, err := client.PostMessage(
		channelID,
		// uncomment the item below to add a extra Header to the message, try it out :)
		slack.MsgOptionText("Pods failing", false),
		slack.MsgOptionAttachments(attachment),
	)

	return timestamp, err

}
