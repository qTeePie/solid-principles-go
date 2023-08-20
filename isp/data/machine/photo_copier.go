package machine

import (
	"fmt"

	"github.com/q10357/solid-principles-go/solid/isp/data"
)

type PhotoCopier struct{}

func (pc PhotoCopier) Print(d *data.Document) {
	fmt.Println("Photocopier printing document...")
}

func (pc PhotoCopier) Scan(d *data.Document) {
	fmt.Println("Photocopier scanning document...")
}
