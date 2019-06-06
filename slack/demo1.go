package main

import (
	"encoding/json"
	"fmt"
	"github.com/nlopes/slack"
	"strconv"
	"time"
)

func main() {
	attachment := slack.Attachment{
		Color:         "good",
		Fallback:      "You successfully posted by Incoming Webhook URL!",
		AuthorName:    "GoCn",
		AuthorSubname: "News",
		AuthorLink:    "https://gocn.vip/explore/category-14",
		AuthorIcon:    "https://gocn.vip/uploads/nav_menu/12.jpg",
		Text:          "<!channel>GoCn-News :smile:\nSee <https://api.slack.com/docs/message-formatting#linking_to_channels_and_users>",
		Footer:        "Go-Cn",
		FooterIcon:    "https://gocn.vip/static/common/avatar-max-img.png",
		Ts:            json.Number(strconv.FormatInt(time.Now().Unix(), 10)),
	}
	msg := slack.WebhookMessage{
		Attachments: []slack.Attachment{attachment},
	}
	err := slack.PostWebhook("", &msg)
	if err != nil {
		fmt.Println(err)
	}
}
