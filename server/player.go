package server

import (
	"encoding/json"
	"github.com/aliics/hearts/data"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"time"
)

type player struct {
	*websocket.Conn
	isClosed bool
	id       uuid.UUID
	hand     []data.Card
	points   int
}

func newPlayer(conn *websocket.Conn, id uuid.UUID) *player {
	p := &player{Conn: conn, id: id}
	conn.SetCloseHandler(func(_ int, _ string) error {
		p.isClosed = true
		return nil
	})
	return p
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
