package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"regexp"
	"time"

	astikit "github.com/asticode/go-astikit"
	astilectron "github.com/asticode/go-astilectron"
	bootstrap "github.com/asticode/go-astilectron-bootstrap"
)

// Vars injected via ldflags by bundler
var (
	BuiltAt            string
	VersionAstilectron string
	VersionElectron    string
)

// Application Vars
var (
	debug = true
	w     *astilectron.Window
)

func main() {
	// Parse flags
	flag.Parse()

	// Create logger
	l := log.New(log.Writer(), log.Prefix(), log.Flags())

	url := "index.html"

	// Debug
	if debug {
		url = "http://localhost:5173/index.html"
	}

	// Run bootstrap
	l.Printf("Running app built at %s\n", BuiltAt)
	if err := bootstrap.Run(bootstrap.Options{
		AstilectronOptions: astilectron.Options{
			AppName:            "AstorExample",
			AppIconDefaultPath: "resources/icon.png",
			SingleInstance:     true,
			VersionAstilectron: "0.37.0",
			VersionElectron:    "8.2.0",
		},
		Debug:       debug,
		Logger:      l,
		MenuOptions: nil,
		OnWait: func(_ *astilectron.Astilectron, ws []*astilectron.Window, _ *astilectron.Menu, _ *astilectron.Tray, _ *astilectron.Menu) error {
			w = ws[0]
			go func() {
				time.Sleep(2 * time.Second)
				if debug {
					w.OpenDevTools()
				}
			}()
			return nil
		},
		Windows: []*bootstrap.Window{{
			Homepage:       url,
			MessageHandler: handleMessages,
			Options: &astilectron.WindowOptions{
				BackgroundColor: astikit.StrPtr("#333333"),
				Center:          astikit.BoolPtr(true),
				Resizable:       astikit.BoolPtr(true),
				Height:          astikit.IntPtr(600),
				Width:           astikit.IntPtr(800),
			},
		}},
	}); err != nil {
		l.Fatal(fmt.Errorf("running bootstrap failed: %w", err))
	}
}

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
	case "btn.click":
		var payloadAstorBtnClick PayloadAstorBtnClick
		if err = json.Unmarshal(m.Payload, &payloadAstorBtnClick); err == nil {
			anserPayloaddAstorBtnText := PayloadAstorBtnText{Button: payloadAstorBtnClick.ButtonID, Text: "clicked"}
			astorEventBtnText := AstorEvent{Name: "butten.text", Payload: anserPayloaddAstorBtnText}
			w.SendMessage(astorEventBtnText)
		}
	}

	return nil, err
}
