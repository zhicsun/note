package structural

import (
	"fmt"
	"testing"
)

func TestAdapter(t *testing.T) {
	client := &Client{}
	mac := &MacA{}

	client.InsertLightningConnectorIntoComputer(mac)

	windowsMachine := &WindowsA{}
	windowsMachineAdapter := &WindowsAdapter{
		windowMachine: windowsMachine,
	}

	client.InsertLightningConnectorIntoComputer(windowsMachineAdapter)
}

type Client struct{}

func (c *Client) InsertLightningConnectorIntoComputer(com IComputer) {
	fmt.Println("Client inserts Lightning connector into computer.")
	com.InsertIntoLightningPort()
}

type IComputer interface {
	InsertIntoLightningPort()
}

type MacA struct{}

func (m *MacA) InsertIntoLightningPort() {
	fmt.Println("Lightning connector is plugged into mac machine.")
}

type WindowsA struct{}

func (w *WindowsA) insertIntoUSBPort() {
	fmt.Println("USB connector is plugged into windows machine.")
}

type WindowsAdapter struct {
	windowMachine *WindowsA
}

func (w *WindowsAdapter) InsertIntoLightningPort() {
	fmt.Println("Adapter converts Lightning signal to USB.")
	w.windowMachine.insertIntoUSBPort()
}
