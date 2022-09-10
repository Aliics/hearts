package main

import (
	"encoding/json"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"time"
)

type player struct {
	*websocket.Conn
	id     uuid.UUID
	hand   []Card
	points int
}

func (p player) writeClientViolation(msg string) {
	p.writeOutboundEvent(outboundEventClientViolation, map[string]any{"msg": msg})
}

func (p player) writeOutboundEvent(eventType outboundEventType, data map[string]any) {
	we, err := json.Marshal(websocketMessage{
		Type: string(eventType),
		Data: data,
	})
	logNonFatal(err)
	logNonFatal(p.WriteMessage(websocket.TextMessage, we))
}

func (p player) writeCloseMessageError(err error) {
	logNonFatal(p.WriteControl(
		websocket.CloseMessage,
		websocket.FormatCloseMessage(websocket.CloseInternalServerErr, err.Error()),
		time.Now().Add(time.Second),
	))
}
