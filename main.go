package main

import (
	. "./config"
	"./elevio"
	"./fsm"
	"./fsmfunctions"
	"./network"
	"./network/localip"
	"./orderhandler"
	"fmt"
	"os"
	"os/exec"
	"os/signal"
	"time"
)

func main() {

	new_order := make(chan elevio.ButtonEvent) //Order from OrderHandler to FSM



	localUUID := make(chan string)

	incoming_message := make(chan MESSAGE)
	outgoing_message := make(chan MESSAGE)

	online_status := make(chan PEER_STATUS_UPDATE) // The online status of peers to OrderHandler

	my_ip, err := localip.LocalIP()
	if err != nil {
		//errorhandling
	}

	go ble.Scanner(localUUID)
	go uuidhandler.UUIDHandler(incoming_message, online_status, my_ip, localUUID)
	go network.Sync(incoming_message, online_status, my_ip)

}
