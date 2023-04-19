package main

import (
	"fmt"
	"log"
	"net/http"
	"webRTC_VC/server"
)

func main() {
	server.AllRooms.Init()

	http.HandleFunc("/create", server.CreateRoomRequest)
	http.HandleFunc("/join", server.JoiningtheRoom)

	log.Println("Starting RTC servers\nEstablishing Server Connection with PORT:5555")
	fmt.Println("")
	err := http.ListenAndServe(":9999", nil)
	if err != nil {
		log.Fatal(err)
	}
}
