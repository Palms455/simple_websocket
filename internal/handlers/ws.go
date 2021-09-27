package handlers

import (
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"sort"
)

//wsChan канал для передачи данных от пользователей
var wsChan = make(chan WsPayload)

// clients map с подключенными клиентами
var clients = make(map[WebSocketConnection]string)

var wsConnection = websocket.Upgrader{
	ReadBufferSize: 1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},

}

// WebSocketConnection содержит соединеие websocket
type WebSocketConnection struct {
	*websocket.Conn
}


// WsJson структура передаваемого сообщения
type WsJson struct {
	Action string `json:"action"`
	Message string `json:"message"`
	MessageType string `json:"message_type"`
	ConnectedUsers []string `json:"connected_users"`
}

// WsPayload структура для передачи сообщения в канал
type WsPayload struct {
	Action string `json:"action"`
	Username string `json:"username"`
	Message string `json:"message"`
	Conn WebSocketConnection `json:"-"`

}


// WsEndpoint views для ws
func WsEndpoint(w http.ResponseWriter, r *http.Request) {
	ws, err := wsConnection.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
	}
	responseJson := WsJson{
		Message: `<em>Connected to server</em>`,
	}
	conn := WebSocketConnection{Conn: ws}
	clients[conn] = ""

	if err := ws.WriteJSON(responseJson); err != nil {
		log.Println(err)
	}
	go ListenWs(&conn)
}

// ListenWs функция прослушиватель даннных из соединения ws
func ListenWs(conn *WebSocketConnection) {
	defer func () {
		if err := recover(); err != nil {
			log.Printf("Ошибка %s", err)
		}
	}()
	var payload WsPayload
	for {
		err := conn.ReadJSON(&payload)
		if err != nil {
			log.Printf("Ошибка чтения json: %s", err)
		} else {
			payload.Conn = *conn
			wsChan <- payload
		}
	}
}

// ListenToWsChan  прослушивание канала с сообщениями
func ListenToWsChan() {
	var response WsJson

	for {
		e := <- wsChan
		switch e.Action {
		case "username":
			clients[e.Conn] = e.Username
			users := getUsers()
			response.Action = "list_users"
			response.ConnectedUsers = users
			broadCastAll(response)

			//response.Action = "Получены данные"
			//response.Message = fmt.Sprintf("Сообщение %s", e.Action)
			//broadCastAll(response)
		}
	}
}

// broadCastAll передача полученного сообщения из канала всем соединениям clients
func broadCastAll(response WsJson)  {
	for client := range clients {
		err := client.WriteJSON(response)
		if err != nil {
			log.Printf("Ошибка websocket: %s", err)
			_ = client.Close()
			delete(clients, client)
		}
	}

}

// getUsers возвращает список пользователей на момент вызова функции
func getUsers() []string{
	var users []string
	for _, user := range clients {
		users = append(users, user)
	}
	sort.Strings(users)
	return users

}