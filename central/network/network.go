package network

import (
	. "../config"
	"./bcast"
	//"./network_helpfunctions"
	"./peers"
	 // "fmt"
	// "../ble/gatt"
)

func Sync(outgoing_message <-chan string, incoming_message chan string, online_status chan PEER_STATUS_UPDATE, localIP string) {

	// Channals for network communication, used in broadcast and transmit functions
	peer_update := make(chan peers.PeerUpdate)
	peer_enable := make(chan bool)
	//broadcast_elevator := make(chan ELEVATOR)
	//recive_elevator := make(chan ELEVATOR)
	//broadcast_message := make(chan []gatt.UUID)
	recive_message := make(chan string)
	//broadcast_acknowledge := make(chan ACKNOWLEDGE_MESSAGE)
	//recive_acknowledge := make(chan ACKNOWLEDGE_MESSAGE)

	go peers.Transmitter(20004, localIP, peer_enable)
	go peers.Receiver(20004, peer_update)
	//go bcast.Transmitter(15647, /*broadcast_elevator,*/ broadcast_message, broadcast_acknowledge)
	go bcast.Receiver(15647, /*recive_elevator,*/ recive_message/*, recive_acknowledge*/)

	for {
		select {
		case peer := <-peer_update:
			if peer.New != "" {
				peer_struct := PEER_STATUS_UPDATE{IP: peer.New, ONLINE: true}
				online_status <- peer_struct
			}
			if len(peer.Lost) > 0 {
				for i := 0; i < len(peer.Lost); i++ {
					peer_struct := PEER_STATUS_UPDATE{IP: peer.Lost[i], ONLINE: false}
					online_status <- peer_struct
				}
			}

		case in_message := <-recive_message:
			incoming_message <- in_message
		}
	}
}
