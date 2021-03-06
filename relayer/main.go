package main

import (
	. "./config"
	"./uuidhandler"
	"./network"
	"./network/localip"
	"./ble"
	// "./ble/gatt"
 	//"fmt"
	// "os"
	// "os/exec"
	// "os/signal"
	//"time"
)

func main() {

	localUUID := make(chan string)

	//incoming_message := make(chan gatt.Advertisement)
	outgoing_message := make(chan string)

	online_status := make(chan PEER_STATUS_UPDATE) // The online status of peers to OrderHandler

	my_ip, err := localip.LocalIP()
	if err != nil {
		//errorhandling
	}

	my_ip = my_ip

	go ble.Scanner(localUUID)
	go uuidhandler.UUIDHandler(outgoing_message, online_status, my_ip, localUUID)
	go network.Sync(outgoing_message, online_status, my_ip)

	select{}

}
