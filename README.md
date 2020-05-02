# Bluetooth Alarm System Framework (golang)

 - This repository contains firmware for a bluetooth alarm system. The relayer detects advertisement events from BLE devices and formward
   messages over UDP to the central not in a sharing ethernet/wifi network.
 - The central code continously listens to messages from the relayer node. When an incoming BLE advertisement message containing
   a matching UUID of know relayer codes an alarm is sounded.

The code is designed for linux based systems
