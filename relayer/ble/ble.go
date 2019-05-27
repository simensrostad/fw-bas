package ble

import (
//	"../elevio"
//	"../fsmfunctions"
	"time"
    "fmt"
    "log"
    "./gatt"
	"strings"
	// "reflect"
)

//var temp gatt.UUID
var test string
var temp []gatt.UUID
var temp2 []gatt.UUID
var str string
var send string

var DefaultClientOptions = []gatt.Option{
	gatt.LnxMaxConnections(1),
	gatt.LnxDeviceID(-1, true),
}

func onStateChanged(d gatt.Device, s gatt.State) {
	fmt.Println("State:", s)
	switch s {
	case gatt.StatePoweredOn:
		fmt.Println("scanning...")
		d.Scan(nil, true)
		return
	default:
		d.StopScanning()
	}
}

func onPeriphDiscovered(p gatt.Peripheral, a *gatt.Advertisement, rssi int) {
		//This section of code is ugly, i dont care
        temp = a.Services
		str = fmt.Sprint(temp)
		t := strings.Replace(str, "[", "", -1)
		send = strings.Replace(t, "]", "", -1)
		//i dont care, really


}

func Scanner(localUUID chan string) {

    d, err := gatt.NewDevice(DefaultClientOptions...)
	if err != nil {
		log.Fatalf("Failed to open device, err: %s\n", err)
		return
	}

	// Register handlers.
	d.Handle(gatt.PeripheralDiscovered(onPeriphDiscovered))
	d.Init(onStateChanged)
    for {
        localUUID <- send
        time.Sleep(10*time.Millisecond)
		//fmt.Println(temp.Services)
    }
}
