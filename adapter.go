package main

import "fmt"

type Printer interface {
	Print()
}

type OldPrinter struct{}

func (OldPrinter) PrintOld() {
	fmt.Println("Printing from Old Printer")
}

type Adapter struct {
	oldPrinter OldPrinter
}

func (a Adapter) Print() {
	a.oldPrinter.PrintOld()
}

func main() {

	var printer Printer

	printer = Adapter{
		oldPrinter: OldPrinter{},
	}

	printer.Print()
}

