package network

import (
	. "../config"
	"./bcast"
	"./network_helpfunctions"
	"./peers"
	"fmt"
)

func Sync(outgoing_message <-chan MESSAGE, incoming_message chan MESSAGE, online_status chan PEER_STATUS_UPDATE, localIP string) {

	// Channals for network communication, used in broadcast and transmit functions
	peer_update := make(chan peers.PeerUpdate)
	peer_enable := make(chan bool)
	//broadcast_elevator := make(chan ELEVATOR)
	//recive_elevator := make(chan ELEVATOR)
	broadcast_message := make(chan MESSAGE)
	recive_message := make(chan MESSAGE)
	broadcast_acknowledge := make(chan ACKNOWLEDGE_MESSAGE)
	recive_acknowledge := make(chan ACKNOWLEDGE_MESSAGE)
	check_acknowledge := make(chan MESSAGE) // Used for checking if a broadcasted order has been acknowledged

	go peers.Transmitter(20004, localIP, peer_enable)
	go peers.Receiver(20004, peer_update)
	go bcast.Transmitter(15647, /*broadcast_elevator,*/ broadcast_message, broadcast_acknowledge)
	go bcast.Receiver(15647, /*recive_elevator,*/ recive_message, recive_acknowledge)

	//var broadcasted_orders [N_NODES]string // Array that holds all orders currently being broadcasted

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

		case out_message := <-outgoing_message:
				go network_helpfunctions.Broadcast_message(out_message, broadcast_message, check_acknowledge)

		case in_message := <-recive_message:
			if in_message.IP == localIP {
				go network_helpfunctions.Broadcast_acknowledge(broadcast_acknowledge, in_message)
				incoming_message <- in_message
			}

		case in_acknowledge := <-recive_acknowledge:
			fmt.Println(in_acknowledge)

		case broadcasted_order := <-check_acknowledge:
			fmt.Println(broadcasted_order)
		}
	}
}
