package config


import (
   //. "../elevio"
)

const (
  N_NODES = 3
  master = 1
)

type MESSAGE struct  {
  IP string
  UUID string
}

type ACKNOWLEDGE_MESSAGE struct  {
    IP string
    UUID string
    NOT_ACKNOWLEDGED bool
}

type PEER_STATUS_UPDATE struct  {
  IP string
  ONLINE bool
}

type PATIENT struct {
    NAME string
    ID string
    UUID string
    LOCATION string
    TIMESTAMP string
}

type Users struct {
    Users []PATIENT
}
