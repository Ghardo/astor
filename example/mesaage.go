package main

import (
	"encoding/json"
	"regexp"

	"github.com/asticode/go-astilectron"
	bootstrap "github.com/asticode/go-astilectron-bootstrap"
)

type PayloadAstorTestEvent struct {
	Text string `json:"text"`
}

type PayloadAstorBtnText struct {
	Button int    `json:"button"`
	Text   string `json:"text"`
}

type PayloadAstorBtnClick struct {
	ButtonID int `json:"buttonId"`
}

type AstorEvent struct {
	Name    string      `json:"name"`
	Payload interface{} `json:"payload"`
}

// handleMessages handles messages
func handleMessages(w *astilectron.Window, m bootstrap.MessageIn) (payload interface{}, err error) {

	re, _ := regexp.Compile("^(.*?)(#.*?)?$")
	match := re.FindStringSubmatch(m.Name)

	switch string(match[1]) {
	case "test.event":
		answerPayloadAstorTestEvent := PayloadAstorTestEvent{Text: "Hello World"}
		return answerPayloadAstorTestEvent, nil
		break
	case "btn.click":
		var payloadAstorBtnClick PayloadAstorBtnClick
		if err = json.Unmarshal(m.Payload, &payloadAstorBtnClick); err == nil {
			anserPayloaddAstorBtnText := PayloadAstorBtnText{Button: payloadAstorBtnClick.ButtonID, Text: "clicked"}
			astorEventBtnText := AstorEvent{Name: "butten.text", Payload: anserPayloaddAstorBtnText}
			w.SendMessage(astorEventBtnText)
		}
		break
	}

	return nil, err
}
