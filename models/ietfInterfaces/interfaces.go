package ietfInterface

import (
	"time"
)

type Interfaces struct {
	NwInterface     []Interface
	InterfacesState []InterfaceState
}

type Interface struct {
	Name                 string
	Description          string
	IntfType             string
	Enabled              bool
	LinkUpDownTrapEnable int
}

type InterfaceState struct {
	Name        string
	IntfType    string
	AdminStatus int
	OperStatus  int
	LastChange  time.Time
	IfIndex     int
	PhysAddress string
	Speed       int64
}
