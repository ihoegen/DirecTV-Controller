package main

import (
	"fmt"

	controller "github.com/ihoegen/DirecTV-Controller/pkg/client"
	"github.com/integrii/flaggy"
)

func main() {
	flaggy.SetName("DTV CLI")
	flaggy.SetDescription("A CLI for interacting with DirecTV")
	var ip string
	port := "8080"
	var channel string

	flaggy.String(&ip, "a", "address", "(Required) IP Address of Receiver")
	flaggy.String(&port, "p", "port", "Port of Receiver (default: 8080)")

	changeChannelSubcommand := flaggy.NewSubcommand("change")
	changeChannelSubcommand.AddPositionalValue(&channel, "channel", 1, true, "Channel to change to")

	skipCommercialsSubcommand := flaggy.NewSubcommand("skip")

	prevSubcommand := flaggy.NewSubcommand("previous")

	powerSubcommand := flaggy.NewSubcommand("power")

	channelSubcommand := flaggy.NewSubcommand("channel")

	onSubcommand := flaggy.NewSubcommand("on")
	offSubcommand := flaggy.NewSubcommand("off")

	powerSubcommand.AttachSubcommand(onSubcommand, 1)
	powerSubcommand.AttachSubcommand(offSubcommand, 1)

	channelSubcommand.AttachSubcommand(prevSubcommand, 1)
	channelSubcommand.AttachSubcommand(changeChannelSubcommand, 1)

	flaggy.AttachSubcommand(powerSubcommand, 1)
	flaggy.AttachSubcommand(channelSubcommand, 1)
	flaggy.AttachSubcommand(skipCommercialsSubcommand, 1)
	flaggy.Parse()

	if ip == "" {
		flaggy.ShowHelpAndExit("Must specify port")
	}

	client := controller.NewController(ip, port)

	if onSubcommand.Used {
		printErr(client.TurnOn())
	} else if offSubcommand.Used {
		printErr(client.TurnOff())
	} else if changeChannelSubcommand.Used {
		printErr(client.ChangeChannel(channel))
	} else if skipCommercialsSubcommand.Used {
		printErr(client.SkipCommercials())
	} else if prevSubcommand.Used {
		printErr(client.PreviousChannel())
	} else {
		flaggy.ShowHelpAndExit("")
	}

}

func printErr(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
