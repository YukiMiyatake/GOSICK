package Slack

import (
	"log"
	"plugin"
	"strings"

	"InputPlugins"

	"util"

	"github.com/nlopes/slack"
)

const (
	actionSelect = "select"
	actionStart  = "satrt"
	actionCancel = "cancel"
)

//

type SlackListener struct {
	client    *slack.Client
	botID     *string
	botName   *[]string
	channelID *string
	rtm       *slack.RTM

	//	allmsg		map[string]func([]string)string
	//	mention		map[string]func([]string)string
	// TODO: オブジェクト化する
	// TODO: mentionと通常メッセージ分けるかも
	promiscous  *map[string]plugin.Symbol
	mention     *map[string]plugin.Symbol
	slackConfig InputPlugins.ISlackConfig
}

func NewSlackListener(sc InputPlugins.ISlackConfig) (InputPlugins.IChatListener, error) {
	log.Printf("[Info] Start slack event listening ")
	client := slack.New(*sc.GetBotToken())

	pm, err := util.NewPluginManager("./plugin.json")
	if err != nil {
		log.Printf("[Error] %s", err)
		return nil, err
	}

	s := &SlackListener{
		client:      client,
		botID:       sc.GetBotID(),
		botName:     sc.GetBotName(),
		channelID:   sc.GetChannelID(),
		promiscous:  &pm.Promiscuous,
		mention:     &pm.Mention,
		slackConfig: sc,
	}

	return s, nil
}

// MessageEvent を待ち受け
func (s *SlackListener) Listen() error {
	s.rtm = s.client.NewRTM()

	go s.rtm.ManageConnection()

	for msg := range s.rtm.IncomingEvents {
		switch ev := msg.Data.(type) {
		case *slack.MessageEvent:
			log.Printf(ev.Channel)
			if err := s.handleMessageEvent(ev); err != nil {
				log.Printf("[Error] Failed to handle message %s", err)
			}
		default:
		}
	}
	return nil
}

// MessageEvent を処理
func (s *SlackListener) handleMessageEvent(ev *slack.MessageEvent) error {
	// Validate section
	if ev.Channel != *s.channelID {
		log.Printf("%s %s", ev.Channel, ev.Msg.Text)
		return nil
	}

	// TODO: BOTからのメッセージ除外したい
	// TODO: メンションだとIDがはいってくるので変換・・
	//	msgs := strings.Split(strings.TrimSpace(ev.Msg.Text), " ")
	msgs := strings.Fields(ev.Msg.Text)
	//if (strings.HasPrefix(ev.Msg.Text, s.botID) ||
	//	strings.HasPrefix(ev.Msg.Text, "<@" + s.botID + "%s>" )) {
	log.Printf(ev.Msg.Text)
	md := util.NewMessageDispatcher(s.botName, s.botID)
	if md.GetMessageType(msgs[0]) == util.Mention {

		log.Printf("*mention*")
		for key, value := range *s.mention {

			if msgs[1] == key {
				// TODO: 引数をjoinしておく？
				// TODO: 引数をポインタにする
				s.rtm.SendMessage(s.rtm.NewOutgoingMessage(value.(func([]string) string)(msgs[2:]), *s.channelID))
				return nil
			}

		}
		//		s.rtm.SendMessage(s.rtm.NewOutgoingMessage("なぁに？ " + strings.Join(msgs[1:], " " ), s.channelID))

	}

	// 該当なし
	//		log.Printf( s.botID )
	//		log.Printf( fmt.Sprintf("<@%s>", s.botID) )
	//		log.Printf( ev.Msg.Text )

	/*
		m := strings.Split(strings.TrimSpace(ev.Msg.Text), " ")[1:]
		if len(m) == 0 || m[0] != "keyword" {
			return fmt.Errorf("invalid message")
		}

		// Display menu
		attachment := slack.Attachment{
			Text:       "What can I do",
			Color:      "#f9a41b",
			CallbackID: "something",
			Actions: []slack.AttachmentAction{
				{
					Name: actionSelect,
					Type: "select",
					Options: []slack.AttachmentActionOption{
						{
							Text:  "content 1",
							Value: "content 1 value",
						},
						{
							Text:  "content 2",
							Value: "content 2 value",
						},
					},
				},

				{
					Name:  actionCancel,
					Text:  "Cancel",
					Type:  "button",
					Style: "danger",
				},
			},
		}

		params := slack.PostMessageParameters{
			Attachments: []slack.Attachment{
				attachment,
			},
		}

		if _, _, err := s.client.PostMessage(ev.Channel, "", params); err != nil {
			return fmt.Errorf("Failed to post message: %s", err)
		}
	*/
	return nil
}
