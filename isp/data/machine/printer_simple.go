package machine

import (
	"fmt"

	"github.com/q10357/solid-principles-go/solid/isp/data"
)

type MyPrinter struct {
}

func (mp MyPrinter) Print(d *data.Document) {
	fmt.Println("Simple printer printing document...")
}
