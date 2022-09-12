package server

import (
	"encoding/json"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"hearts/data"
	"time"
)

type player struct {
	*websocket.Conn
	id     uuid.UUID
	hand   []data.Card
	points int
}

func (p player) writeClientViolation(msg string) {
	p.writeOutboundEvent(data.OutboundEventClientViolation, map[string]any{"msg": msg})
}

func (p player) writeOutboundEvent(eventType data.OutboundEventType, eventData map[string]any) {
	we, err := json.Marshal(data.WebsocketMessage{
		Type: string(eventType),
		Data: eventData,
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
