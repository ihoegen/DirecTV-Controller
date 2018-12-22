package directv

import (
	"net/http"
)

// Controller implements Actions
type Controller struct {
	ReceiverEndpoint string
}

// ChangeChannel changes the channel on a DirecTV receiver
func (c Controller) ChangeChannel(channel string) error {
	_, err := http.Get(c.ReceiverEndpoint + "/tv/tune?major=" + channel + "&minor=65535")
	return err
}

// TurnOn turns on a DirecTV receiver
func (c Controller) TurnOn() error {
	return c.togglePower()
}

// TurnOff turns off a DirecTV receiver
func (c Controller) TurnOff() error {
	return c.togglePower()
}

func (c Controller) togglePower() error {
	_, err := http.Get(c.ReceiverEndpoint + "/remote/processKey?key=power&hold=keyPress")
	return err
}

// PreviousChannel changes the channel to previous on a DirecTV receiver
func (c Controller) PreviousChannel() error {
	_, err := http.Get(c.ReceiverEndpoint + "/remote/processKey?key=prev&hold=keyPress")
	return err
}

// SkipCommercials skips 3 minutes on a DirecTV receiver
func (c Controller) SkipCommercials() error {
	for i := 0; i < 6; i++ {
		_, err := http.Get(c.ReceiverEndpoint + "/remote/processKey?key=advance&hold=keyPress")
		if err != nil {
			return err
		}
	}
	return nil
}
