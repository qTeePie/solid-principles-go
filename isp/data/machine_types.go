package data

type Document struct {
}

type Printer interface {
	Print(d *Document)
}

type Scanner interface {
	Scan(d *Document)
}

type Faxer interface {
	Fax(d *Document)
}

type MultiFunctionDevice interface {
	Printer
	Scanner
	Faxer
}
