package network_helpfunctions

import (
	// . "../../config"
	//"time"
	"../../ble/gatt"
)

func Broadcast_message(message []gatt.UUID, broadcast_message chan<- []gatt.UUID) {
	// broadcast_ticker := time.NewTicker(30 * time.Millisecond)
	// broadcast_done_timer := time.NewTimer(100 * time.Millisecond)
	// for {
	// 	select {
	// 	case <-broadcast_ticker.C:
			broadcast_message <- message
	// 	case <-broadcast_done_timer.C:
	// 		return
	// 	}
	// }
}

// func Broadcast_struct(local_struct ELEVATOR, broadcast_channel chan<- ELEVATOR) {
// 	broadcast_ticker := time.NewTicker(30 * time.Millisecond)
// 	broadcast_done_timer := time.NewTimer(150 * time.Millisecond)
// 	for {
// 		select {
// 		case <-broadcast_ticker.C:
// 			broadcast_channel <- local_struct
// 		case <-broadcast_done_timer.C:
// 			return
// 		}
// 	}
// }
