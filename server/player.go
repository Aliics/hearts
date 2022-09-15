package server

import (
	"encoding/json"
	"github.com/aliics/hearts/data"
	"github.com/aliics/hearts/util"
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

func (p player) writeClientViolation(message string) {
	p.writeOutboundPayload(data.OutboundPayload{
		Type: data.OutboundClientViolation,
		Data: data.ClientViolationEvent{Message: message},
	})
}

func (p player) writeOutboundPayload(payload data.OutboundPayload) {
	we, err := json.Marshal(payload)
	util.LogNonFatal(err)
	util.LogNonFatal(p.WriteMessage(websocket.TextMessage, we))
}

func (p player) WriteCloseMessageError(err error) {
	util.LogNonFatal(p.WriteControl(
		websocket.CloseMessage,
		websocket.FormatCloseMessage(websocket.CloseInternalServerErr, err.Error()),
		time.Now().Add(time.Second),
	))
}
