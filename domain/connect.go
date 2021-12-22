package domain

import "github.com/gorilla/websocket"

type WsJsonResponse struct {
	Action string `json:"action"`
	Post   string `json:"post"`
}

type WebSocketConnection struct {
	*websocket.Conn
}

type WsPayload struct {
	Action   string              `json:"action"`
	Post     string              `json:"Post"`
	Username string              `json:"username"`
	Conn     WebSocketConnection `json:"-"`
}
