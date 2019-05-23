package ble

import (
//	. "../config"
//	"../elevio"
//	"../fsmfunctions"
	//"time"
    "fmt"
    "log"
    "github.com/paypal/gatt"
    "github.com/paypal/gatt/examples/option"
)

func onStateChanged(d gatt.Device, s gatt.State) {
	fmt.Println("State:", s)
	switch s {
	case gatt.StatePoweredOn:
		fmt.Println("scanning...")
		d.Scan([]gatt.UUID{}, true)
		return
	default:
		d.StopScanning()
	}
}

func onPeriphDiscovered(p gatt.Peripheral, a *gatt.Advertisement, rssi int) {
    // tmp := []gatt.UUID{}
    //if a.Services == tmp {
        fmt.Println("  UUID   =", a.Services)
    //}
}

func Scanner(localUUID chan string) {

    d, err := gatt.NewDevice(option.DefaultClientOptions...)
	if err != nil {
		log.Fatalf("Failed to open device, err: %s\n", err)
		return
	}

    //fmt.Println(d)

	// Register handlers.
	d.Handle(gatt.PeripheralDiscovered(onPeriphDiscovered))
	d.Init(onStateChanged)
    select {}

}
