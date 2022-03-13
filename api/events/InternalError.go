package events

import (
	"fmt"
	"github.com/Kalitsune/Haddock/api/ws"
	"log"
)

type InternalErrorArgs struct {
	Err     error
	Logging bool
}

func InternalError(client *ws.Client, msg string, args InternalErrorArgs) {
	/*
		If asked send a log
	*/
	if args.Logging {
		log.Println(msg, args.Err)
	}

	/*
		Sent an event to the client
	*/
	msg = fmt.Sprintf(`{"reason": "%s"}`, msg)
	event := Event{Name: "ERROR", Args: []byte(msg)}
	event.Send(client)
}
