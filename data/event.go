package data

type WebsocketMessage struct {
	Type string         `json:"type"`
	Data map[string]any `json:"data"`
}

type InboundEventType string

const (
	InboundEventPlayCard InboundEventType = "playCard"
)

type OutboundEventType string

const (
	OutboundEventClientViolation OutboundEventType = "clientViolation"
	OutboundEventGameUpdate      OutboundEventType = "gameUpdate"
)
