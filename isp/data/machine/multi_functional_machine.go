package machine

import (
	"github.com/q10357/solid-principles-go/solid/isp/data"
)

type MultiFunctionMachine struct {
	printer data.Printer
	scanner data.Scanner
	faxer   data.Faxer
}

func (m MultiFunctionMachine) Print(d *data.Document) {
	m.printer.Print(d)
}

func (m MultiFunctionMachine) Fax(d *data.Document) {
	m.scanner.Scan(d)
}

func (m MultiFunctionMachine) Scan(d *data.Document) {
	m.scanner.Scan(d)
}
