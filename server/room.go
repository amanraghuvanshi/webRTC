package server

import (
	"fmt"
	"math/rand"
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

// a participant has a value that identifies if they are a host and holds a pointer to the WebSocket connection
type Participant struct {
	Host bool
	Conn *websocket.Conn
}

// the RoomMap struct has a Read and Write Mutex value; map that from a string to an array of participants
type RoomMap struct {
	Mutex *sync.RWMutex
	Map   map[string][]Participant
}

// purpose of RoomMap: to attach the name of a room to who is in it.

// this will return a Participant
func (r *RoomMap) Init() {
	r.Map = make(map[string][]Participant)
}

// create the name of this room, we use the library “math/rand” to create a random RoomID
func (r *RoomMap) Get(roomID string) []Participant {
	r.Mutex.RLock()
	defer r.Mutex.RUnlock()

	return r.Map[roomID]
}

// this function creates the room for the particular session
// locking and improvisioning of resources and allocation of url
func (r *RoomMap) CreateRoom() string {
	r.Mutex.Lock()
	defer r.Mutex.Unlock()

	rand.Seed(time.Now().UnixNano())
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")
	b := make([]rune, 8)

	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}

	roomID := string(b)

	r.Map[roomID] = []Participant{}

	return roomID
}

// this function adds the user to the room
// the mutex alongside, will lock the resource when the room are being created
func (r *RoomMap) InsertIntoRoom(roomID string, host bool, conn *websocket.Conn) {
	r.Mutex.Lock()
	defer r.Mutex.Unlock()

	p := Participant{host, conn}

	fmt.Println("Inserting into the RoomID: ", roomID)
	r.Map[roomID] = append(r.Map[roomID], p)
}

// the function for deleting the particular room

func (r *RoomMap) DeleteRoom(roomID string) {
	r.Mutex.Lock()
	defer r.Mutex.Unlock()
	delete(r.Map, roomID)
}
