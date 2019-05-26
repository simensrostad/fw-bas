Elevator project in Real-time programming
==========================================
Out solution uses the UDP comminication protocol. Every elevator contains the elevator struct of all the other elevators in the system.

type ELEVATOR struct {  
  State - Elevator state  
  Dir  - Elevator direction  
  Orders - Orders currently in the elevator  
  Floor  - Current floor the elevator  
  IP     - Elevators IP address  
  Online - Onlinestatus of the elevator  
}

Based on this information each elevator decides if an order should be taken by itself or given to one of the elevators on the network. (peer-to-peer)
If the order is to be handled by another elevator, the elevator broadcasts the order and waits for an acknowledge message from that particular elevator.
If an elevator goes offline the other elevators are notified, and will redistribute the offline elevators orders after a given amount of time.

Modules:
main - Initializes channels and elevator server/simulator, starts other modules and contains a backup function which restarts main upon error.
fsm - Elevator functionality, controls the behvaiour of the elevator with an state machine and detects malfunction in the actual elevator hardware.
orderhandler - Builds a large matrix containing the struct from the other elevator, calculates which elevator to take a particular order, sets lights and onlinestatus of
the elevators in the system, redistributes orders from offline elevators.
network - Handles network communications, broadcasts orders and elevator structs and checks which elevators are connected to the system.

The system is scalable by changing the parameters contained in the config file


The program utilizes a predefined network and elevator driver module from the TTK4145 github course page (www.github.com/TTK4145).
The costfunction in the orderhandler module and backup function in main module is based on www.github.com/PerKjelsvik/TTK4145-sanntid.
