package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"

	"github.com/nlopes/slack"
)

// interactive message response
type interactionHandler struct {
	slackClient       *slack.Client
	verificationToken string
}

// receive result action for interactive message
func (h interactionHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Printf("ServerHTTP")
	if r.Method != http.MethodPost {
		log.Printf("[Error] Invalid method: %s", r.Method)
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	buf, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("[Error] Failed to read request body: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	jsonData, err := url.QueryUnescape(string(buf)[8:])
	if err != nil {
		log.Printf("[Error] Failed to unescape request body: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var message slack.AttachmentActionCallback
	if err := json.Unmarshal([]byte(jsonData), &message); err != nil {
		log.Printf("[Error] Failed decode json message: %s", jsonData)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if message.Token != h.verificationToken {
		log.Printf("[Error] Invalid token: %s", message.Token)
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	action := message.Actions[0]
	switch action.Name {
	case actionSelect:
		value := action.SelectedOptions[0].Value

		// Overwrite original message
		originalMessage := message.OriginalMessage
		originalMessage.Attachments[0].Text = fmt.Sprintf("overwrited %s", strings.Title(value))
		originalMessage.Attachments[0].Actions = []slack.AttachmentAction{
			{
				Name:  actionStart,
				Text:  "Yes",
				Type:  "button",
				Value: "start",
				Style: "primary",
			},
			{
				Name:  actionCancel,
				Text:  "No",
				Type:  "button",
				Style: "danger",
			},
		}

		w.Header().Add("Content-type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(&originalMessage)

	case actionStart:
		title := ":ok: submitted!"
		responseMessage(w, message.OriginalMessage, title, "")
		return

	case actionCancel:
		title := fmt.Sprintf(":x: canceled by %s", message.User.Name)
		responseMessage(w, message.OriginalMessage, title, "")
		return

	default:
		log.Printf("[Error] Invalid action \"%s\"", action.Name)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func responseMessage(w http.ResponseWriter, original slack.Message, title, value string) {
	original.Attachments[0].Actions = []slack.AttachmentAction{}
	original.Attachments[0].Fields = []slack.AttachmentField{
		{
			Title: title,
			Value: value,
			Short: false,
		},
	}

	w.Header().Add("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&original)
}
