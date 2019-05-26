package ble

import (
//	"../elevio"
//	"../fsmfunctions"
	"time"
    "fmt"
    "log"
    "./gatt"
)

var temp gatt.Advertisement

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

        temp = *a

}

func Scanner(localUUID chan gatt.Advertisement) {

    d, err := gatt.NewDevice(DefaultClientOptions...)
	if err != nil {
		log.Fatalf("Failed to open device, err: %s\n", err)
		return
	}

	// Register handlers.
	d.Handle(gatt.PeripheralDiscovered(onPeriphDiscovered))
	d.Init(onStateChanged)
    for {
        localUUID <- temp
        time.Sleep(30*time.Millisecond)
		fmt.Println(temp.Services)
    }
}
