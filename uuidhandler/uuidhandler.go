package orderhandler

import (
//	"fmt"
//	"time"
//	. "../config"
//	"../elevio"
)

func UUIDHandler(incoming_message <-chan MESSAGE, online_status <-chan PEER_STATUS_UPDATE, local_IP string) {

	onlinestatusMatrix := [N_NODES]PEER_STATUS_UPDATE{}
	onlinestatusMatrix[0].IP := local_IP

	for {

		select {

			case online_status := <-online_status:
//
//				for node := 1; elevator < N_NODES; node++ {
//
//						if onlinestatusMatrix.[node].IP == online_status.IP	{
//								onlinestatusMatrix.[node].IP = online_status.ONLINE
//						}
//
//					}

			case localUUID := <-localUUID:
				outgoing_message <- localUUID

			case incoming_message := <-incoming_message:
				fmt.Println(incoming_message)

		}
	}
}
