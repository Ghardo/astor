package main

import (
	"fmt"
	"log"
	"time"

	"github.com/asticode/go-astikit"
	"github.com/asticode/go-astilectron"
)

type AstorEvent struct {
	Name    string      `json:"name"`
	Payload interface{} `json:"payload"`
}

func main() {
	// Set logger
	l := log.New(log.Writer(), log.Prefix(), log.Flags())

	// Create astilectron
	a, err := astilectron.New(l, astilectron.Options{
		AppName:            "Test",
		VersionAstilectron: "0.37.0",
		VersionElectron:    "8.2.0",
	})
	if err != nil {
		l.Fatal(fmt.Errorf("main: creating astilectron failed: %w", err))
	}
	defer a.Close()

	// Handle signals
	a.HandleSignals()

	// Start
	if err = a.Start(); err != nil {
		l.Fatal(fmt.Errorf("main: starting astilectron failed: %w", err))
	}

	// New window
	var w *astilectron.Window
	if w, err = a.NewWindow("build/resources/index.html", &astilectron.WindowOptions{
		Center: astikit.BoolPtr(true),
		Height: astikit.IntPtr(700),
		Width:  astikit.IntPtr(1000),
	}); err != nil {
		l.Fatal(fmt.Errorf("main: new window failed: %w", err))
	}

	// Create windows
	if err = w.Create(); err != nil {
		l.Fatal(fmt.Errorf("main: creating window failed: %w", err))
	}

	w.OpenDevTools()
	// This will listen to messages sent by Javascript
	w.OnMessage(func(m *astilectron.EventMessage) interface{} {
		var r AstorEvent
		// Unmarshal
		var e AstorEvent
		m.Unmarshal(&e)

		l.Printf("%v", e)

		if e.Name == "app-ready" {
			r = AstorEvent{Name: e.Name, Payload: "I got your message. Thanks."}
			cm := AstorEvent{Name: "go-custom-message", Payload: "custom Message from go."}

			for i := 0; i < 5; i++ {
				w.SendMessage(cm)
				time.Sleep(1 * time.Second)
			}

			return r
		}

		r = AstorEvent{Name: e.Name, Payload: "unknown message"}
		return r
	})

	// Blocking pattern
	a.Wait()
}
