package client

import (
	dtv "github.com/ihoegen/DirecTV-Controller/pkg/client/directv"
)

// Actions is the list of supported actions
type Actions interface {
	ChangeChannel(string) error
	TurnOn() error
	TurnOff() error
	SkipCommercials() error
	PreviousChannel() error
}

//NewController returns a new controller object
func NewController(ipAddress, port string) Actions {
	return dtv.Controller{
		ReceiverEndpoint: "http://" + ipAddress + ":" + port,
	}
}
