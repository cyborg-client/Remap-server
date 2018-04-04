// Package websocketserver implements the WebSocket connection for serving data to users.
package websocketserver

import (
	"github.com/satori/go.uuid"
)

// splitterRequest defines a request to join the PUBSUB splitterMain. ID
type splitterRequest struct {
	ID     uuid.UUID // A unique ID which no other subscriber has yet acquired
	DataCh chan []int64 // A channel you wish to receive the subscriber data.
}
