package websocket

import (
	"fmt"
	"hotel/auth"
	"hotel/model"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/websocket"
)

// var clients = make(map[*websocket.Conn]bool)

// tao phong client
var clientRooms = make(map[string]*websocket.Conn)

var broadcast = make(chan MessageForEachPeople)

// Configure the upgrader
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type MessageForEachPeople struct {
	UserID  uint   `json:"userid"`
	Message string `json:"message"`
}

// Define our message object
type Message struct {
	Token   string `json:"token"`
	UserID2 string `json:"userid2"`
	Message string `json:"message"`
}

//tao trc 1 room

func ChatAPI() {
	go handleMessages()
}

var roomid string
var roomid1 string

func HandleConnections(w http.ResponseWriter, r *http.Request) {
	// tao websocket
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer ws.Close()
	// Register our new client
	for {
		var msg Message
		// Read in a new message as JSON and map it to a Message object
		err := ws.ReadJSON(&msg)
		DecodeToken := auth.DecodeToken(msg.Token)
		if DecodeToken == nil {
			fmt.Fprint(w, "Can't not decode token")
		}
		resultID := DecodeToken["user_id"]
		UserID1 := fmt.Sprintf("%v", resultID)
		if UserID1 != "" && msg.UserID2 != "" {
			roomid = UserID1 + "+" + msg.UserID2
			roomid1 = msg.UserID2 + "+" + UserID1
			fmt.Println(UserID1)
			fmt.Println("-------------")
			fmt.Println(msg.UserID2)
			fmt.Println("-------------")
			clientRooms[roomid] = ws
			if err != nil {
				log.Printf("error: %v", err)
				delete(clientRooms, roomid)
				continue
			}
			// tao conversation
			conversationID := model.CheckConvsersation(UserID1, msg.UserID2)
			// tao mess trong conversation
			arrmess := model.CreateMessenger(UserID1, msg.Message, conversationID)
			//hien thi toan bo tin nhan sau chuyen vao broadcast
			// hien thi toan bo tin nhan cu
			if msg.Message == "" {
				for i := 0; i < len(arrmess); i++ {
					a := MessageForEachPeople{
						UserID:  arrmess[i].UserID,
						Message: arrmess[i].Messenger,
					}
					fmt.Println("-------11--")
					broadcast <- a
				}
				continue
			} else {
				id1, _ := strconv.ParseUint(UserID1, 10, 64)
				newMess := MessageForEachPeople{
					UserID:  uint(id1),
					Message: msg.Message,
				}
				broadcast <- newMess
			}
		}
	}
}
func handleMessages() {
	for {
		// Grab the next message from the broadcast channel
		msg := <-broadcast
		// Send it out to every client that is currently connected
		for key := range clientRooms {
			if clientRooms[key] == clientRooms[roomid] {
				err := clientRooms[roomid].WriteJSON(msg)
				if err != nil {
					log.Printf("error: %v", err)
					clientRooms[roomid].Close()
					delete(clientRooms, roomid)
				}
			}
			if clientRooms[key] == clientRooms[roomid1] {
				err := clientRooms[roomid1].WriteJSON(msg)
				if err != nil {
					log.Printf("error: %v", err)
					clientRooms[roomid1].Close()
					delete(clientRooms, roomid1)
				}
			}
		}
	}
}
