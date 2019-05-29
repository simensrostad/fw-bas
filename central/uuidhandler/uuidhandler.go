package uuidhandler

import (
	"fmt"
    //"time"
	. "../config"
//	"../elevio"
	// "../ble/gatt"
	"../readwrite"
)

func UUIDHandler(outgoing_message chan string, incoming_message <-chan string, online_status <-chan PEER_STATUS_UPDATE, local_IP string, localUUID <-chan string) {

	onlinestatusMatrix := [N_NODES]PEER_STATUS_UPDATE{}
	onlinestatusMatrix[0].IP = local_IP

	for {

		select {
			case online_status := <-online_status:
				for node := 1; node < N_NODES; node++ {
					if onlinestatusMatrix[node].IP == online_status.IP	{
							onlinestatusMatrix[node].ONLINE = online_status.ONLINE
					}

				}

			case localUUID := <-localUUID:
				ja := readwrite.Read_from_file(localUUID)
				if ja.UUID != "" {
					fmt.Println(ja)
				}

			case incoming_message := <-incoming_message:
				ja := readwrite.Read_from_file(incoming_message)
				if ja.UUID != "" {
					fmt.Println(ja)
				}
		}
	}
}
