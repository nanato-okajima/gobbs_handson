package handlers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/CloudyKit/jet/v6"
	"github.com/gorilla/websocket"

	"gobbs_handson/bulletin-board/domain"
)

var (
	wsChan = make(chan domain.WsPayload)

	clients = make(map[domain.WebSocketConnection]string)

	views = jet.NewSet(
		jet.NewOSFileSystemLoader("./html"),
		jet.InDevelopmentMode(),
	)

	upgradeConnection = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin:     func(r *http.Request) bool { return true },
	}
)

func WsEndpoint(w http.ResponseWriter, r *http.Request) {
	wsConn, err := upgradeConnection.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
	}
	log.Println("WS Connectiong....")

	conn := domain.WebSocketConnection{Conn: wsConn}
	clients[conn] = ""

	go ListenForWs(&conn)
}

func Home(w http.ResponseWriter, r *http.Request) {
	err := renderPage(w, "home.jet", nil)
	if err != nil {
		log.Println(err)
	}
}

func renderPage(w http.ResponseWriter, tmpl string, data jet.VarMap) error {
	view, err := views.GetTemplate(tmpl)
	if err != nil {
		log.Println(err)
		return err
	}

	err = view.Execute(w, data, nil)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func ListenForWs(wsConn *domain.WebSocketConnection) {
	defer func() {
		if r := recover(); r != nil {
			log.Println("Error", fmt.Sprintf("%v", r))
		}
	}()

	var payload domain.WsPayload

	for {
		if err := wsConn.ReadJSON(&payload); err != nil {
			//noaction
		} else {
			payload.Conn = *wsConn
			wsChan <- payload
		}
	}
}

func broadcastToAllUser(response domain.WsJsonResponse) {
	for client := range clients {
		if err := client.WriteJSON(response); err != nil {
			log.Println("WebSockets Error")
			_ = client.Close()
			delete(clients, client)
		}
	}
}

func ListenToWsChannel() {
	var response domain.WsJsonResponse

	for {
		payload := <-wsChan

		switch payload.Action {
		case "username":
			clients[payload.Conn] = payload.username
		case "left":
			delete(clients, payload.Conn)
		}
		// response.Action = "Sample Action"
		// response.Post = fmt.Sprintf(`
		// 	<div class="message">
		// 		<p>%s</p>
		// 		<small class="name">%s</small>
		// 	</div>`, payload.Username, payload.Post)

		// broadcastToAllUser(response)
	}
}
