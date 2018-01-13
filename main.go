package main

import (
	"net/http"
	"fmt"
	"log"
	"github.com/gorilla/websocket"
	"github.com/satori/go.uuid"
)

func Home(w http.ResponseWriter,r *http.Request){
	fmt.Println(r.Body)
	http.ServeFile(w, r, "home.html")
}
type user struct{
	name string
	conn *websocket.Conn
}
func (c *Controller) ReadMessage(){

}
type Controller struct {
	Online map[uuid.UUID]*user
}
const (
	writeWait = 10 * time.Second
	pongWait = 60 * time.Second
	pingPeriod = (pongWait * 9) / 10
	maxMessageSize = 512
)
var (
	online  = make(map[uuid.UUID]*user)
	Control Controller = Controller{online}
)

func WebSocketHandler(w http.ResponseWriter,r *http.Request){
	cookie,err := r.Cookie("uuid")
	if err != nil {
		http.setCookie()
	}
	username,ok := r.URL.Query()["name"]
	fmt.Println(r)
	if !ok || len(username) < 1{
		log.Fatal(ok)
		return
	}
	fmt.Println(username)
	upgrader := &websocket.Upgrader{}
	conn,err := upgrader.Upgrade(w,r,nil)
	if err != nil {
		log.Println(err)
		return
	}
	id := uuid.Must(uuid.NewV4())
	online[id] = &user{username[0],conn}
	fmt.Println(online)
	go Control.ReadMessage()
}
func main() {
	http.HandleFunc("/",Home)
	http.HandleFunc("/ws",WebSocketHandler)
	http.ListenAndServe(":8080",nil)
}