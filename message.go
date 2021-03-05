package main

import (
	"github.com/asticode/go-astilectron"
	bootstrap "github.com/asticode/go-astilectron-bootstrap"
)

// handleMessages handles messages
func handleMessages(_ *astilectron.Window, m bootstrap.MessageIn) (payload interface{}, err error) {
	// Handle Templating from here on out.
	// TODO: Add templating...


	//switch m.Name {
	//case "explore":
	//	// Unmarshal payload
	//	var path string
	//	if len(m.Payload) > 0 {
	//		// Unmarshal payload
	//		if err = json.Unmarshal(m.Payload, &path); err != nil {
	//			payload = err.Error()
	//			return
	//		}
	//	}
	//
	//	// Explore
	//	if payload, err = explore(path); err != nil {
	//		payload = err.Error()
	//		return
	//	}
	//}
	//return
	return
}