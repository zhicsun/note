package structural

import (
	"fmt"
	"testing"
)

func TestBridge(t *testing.T) {
	hpPrinter := &Hp{}
	epsonPrinter := &Epson{}

	macComputer := &MacB{}
	macComputer.SetPrinter(hpPrinter)
	macComputer.Print()
	fmt.Println()
	macComputer.SetPrinter(epsonPrinter)
	macComputer.Print()
	fmt.Println()

	winComputer := &WindowsB{}
	winComputer.SetPrinter(hpPrinter)
	winComputer.Print()
	fmt.Println()

	winComputer.SetPrinter(epsonPrinter)
	winComputer.Print()
	fmt.Println()
}

type IComputerB interface {
	Print()
	SetPrinter(Printer)
}

type MacB struct {
	printer Printer
}

func (m *MacB) Print() {
	fmt.Println("Print request for mac")
	m.printer.PrintFile()
}

func (m *MacB) SetPrinter(p Printer) {
	m.printer = p
}

type WindowsB struct {
	printer Printer
}

func (w *WindowsB) Print() {
	fmt.Println("Print request for windows")
	w.printer.PrintFile()
}

func (w *WindowsB) SetPrinter(p Printer) {
	w.printer = p
}

type Printer interface {
	PrintFile()
}

type Epson struct{}

func (p *Epson) PrintFile() {
	fmt.Println("Printing by a EPSON Printer")
}

type Hp struct{}

func (p *Hp) PrintFile() {
	fmt.Println("Printing by a HP Printer")
}
