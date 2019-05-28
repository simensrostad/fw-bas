package ble

import (
//	"../elevio"
//	"../fsmfunctions"
	"time"
    "fmt"
    "log"
    "./gatt"
    "strings"
    //"time"
)

var temp []gatt.UUID
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

        //temp.Services = a.Services
        temp = a.Services
        str := fmt.Sprint(temp)
        t := strings.Replace(str, "[", "", -1)
        send = strings.Replace(t, "]", "", -1)

}

func Scanner(localUUID chan string) {
	//
    d, err := gatt.NewDevice(DefaultClientOptions...)
	if err != nil {
		log.Fatalf("Failed to open device, err: %s\n", err)
		return
	}

	// Register handlers.
	d.Handle(gatt.PeripheralDiscovered(onPeriphDiscovered))
	d.Init(onStateChanged)
    for {
        localUUID <- send   //change back to send after network testing works
        time.Sleep(10*time.Millisecond)
    }
}
