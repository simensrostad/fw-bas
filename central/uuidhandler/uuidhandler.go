package uuidhandler

import (
	"fmt"
//	"time"
	. "../config"
//	"../elevio"
	"../ble/gatt"
)

func UUIDHandler(outgoing_message chan []gatt.UUID, incoming_message <-chan []gatt.UUID, online_status <-chan PEER_STATUS_UPDATE, local_IP string, localUUID <-chan []gatt.UUID) {

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
				fmt.Println("i found a package myself,", localUUID)

			case incoming_message := <-incoming_message:
				fmt.Println("This is incoming message from IP", incoming_message)

		}
	}
}
