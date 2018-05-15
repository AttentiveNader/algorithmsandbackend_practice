package main

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
	"github.com/satori/go.uuid"
)

func Home(w http.ResponseWriter, r *http.Request) {
	//	fmt.Println(r.Body)
	//fmt.Println(control.Online,cookie)
	http.ServeFile(w, r, "./home.html")
}

type user struct {
	name string
	conn *websocket.Conn
}

func (c *Controller) ReadMessage() {
	for v, k := range c.Online {
		fmt.Println(v, "then", k)
	}
}
func (u *user) MessageListener( /*cha chan []byte*/ ) {
	defer u.conn.Close()
	u.conn.SetReadLimit(maxMessageSize)
	//u.conn.SetReadDeadline(time.Now().Add(pongWait))
	u.conn.SetPongHandler(func(string) error { u.conn.SetReadDeadline(time.Now().Add(pongWait)); return nil })
	for {
		som, message, err := u.conn.ReadMessage()
		fmt.Println(som)
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway) {
				log.Printf("error: %v", err)
			}
			log.Println(err)
			break
		}
		message = bytes.TrimSpace(message)
		fmt.Println(message)
		u.MessageWriter(message)
		//cha <- message
	}
	fmt.Println(control.Online)
}

type Controller struct {
	Online map[uuid.UUID]*user
}

const (
	writeWait = 10 * time.Second
	pongWait  = 60 * time.Second
	//pingPeriod = (pongWait * 9) / 10
	maxMessageSize = 512
)

var (
	control Controller = Controller{Online: make(map[uuid.UUID]*user)}
)

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	//Home(w,r)
	username, ok := r.URL.Query()["name"]
	fmt.Println("hello", username)
	if !ok || len(username) < 1 {
		log.Fatal(ok)
		return
	}
	fmt.Println(username)
	//fmt.Println(r)
	cookie, err := r.Cookie("uuid")
	var id uuid.UUID
	if err != nil {
		id, _ := uuid.NewV4()
		cookie = &http.Cookie{
			Name:     "uuid",
			Value:    id.String(),
			Path:     "/",
			HttpOnly: true,
		}
		//control.Online[id] = &user{username[0],nil}
		http.SetCookie(w, cookie)
	}
	id = uuid.Must(uuid.FromString(cookie.Value))
	if control.Online[id] != nil {
		control.Online[id].name = username[0]
	} else {
		control.Online[id] = &user{name: username[0]}
	}
	Home(w, r)
}
func (u *user) MessageWriter(message []byte) {
	u.conn.SetWriteDeadline(time.Now().Add(writeWait))
	w, err := u.conn.NextWriter(websocket.TextMessage)
	if err != nil {
		return
	}
	fmt.Println(message)
	w.Write(message)
	w.Close()
	//	w.Write([]byte{'z','z','z'})
}
func WebSocketHandler(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("uuid")
	if err != nil {
		fmt.Println(err)
		return
	}
	upgrader := &websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err, "hello")
		return
	}
	id := uuid.Must(uuid.FromString(cookie.Value))
	if control.Online[id] != nil {
		control.Online[id].conn = conn
	} else {
		control.Online[id] = &user{conn: conn}
	}
	fmt.Println(control.Online)
	//	cha := make(chan []byte)
	go control.Online[id].MessageListener()
	//control.Online[id].MessageWriter([]byte{'h','h','a'})
	//go MessageListener(control.Online[uuid.Must(uuid.FromString(cookie.Value))])
}
func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/ws", WebSocketHandler)
	http.HandleFunc("/register", RegisterHandler)
	http.ListenAndServe(":8080", nil)
	//	control.ReadMessage()
}
