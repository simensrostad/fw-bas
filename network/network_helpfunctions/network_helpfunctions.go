package network_helpfunctions

import (
	. "../../config"
	"time"
)

func Broadcast_message(message MESSAGE, broadcast_message chan<- MESSAGE, check_acknowledge chan<- MESSAGE) {
	broadcast_message <- message
	no_acknowledge_timer:= time.NewTimer(150 * time.Millisecond)
	for {
		select {
		case <-no_acknowledge_timer.C:
			check_acknowledge <- message
			return
		}
	}
}

func Broadcast_acknowledge(broadcast_ack chan<- ACKNOWLEDGE_MESSAGE , message MESSAGE) {
	out_acknowledge_message := ACKNOWLEDGE_MESSAGE{IP: message.IP, Order: message.Order, NOT_ACKNOWLEDGED: false}
	broadcast_ack <- out_acknowledge_message
}

func Broadcast_struct(local_struct ELEVATOR, broadcast_channel chan<- ELEVATOR) {
	broadcast_ticker := time.NewTicker(30 * time.Millisecond)
	broadcast_done_timer := time.NewTimer(150 * time.Millisecond)
	for {
		select {
		case <-broadcast_ticker.C:
			broadcast_channel <- local_struct 
		case <-broadcast_done_timer.C:
			return
		}
	}
}
