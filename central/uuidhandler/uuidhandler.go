package uuidhandler

import (
	"fmt"
	// "time"
	. "../config"
//	"../elevio"
	// "../ble/gatt"
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
				if localUUID == "6e400001b5a3e393e0e9a50e24dcca9f"	{
					fmt.Println("i found a package myself,", localUUID)
				}


			case incoming_message := <-incoming_message:
				if incoming_message == "6e400001b5a3e393e0e9a50e24dcca9f"	{
					fmt.Println("The Bacon found the package", incoming_message)
				}
		}
	}
}
