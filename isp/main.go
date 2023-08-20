package main

import (
	"github.com/q10357/solid-principles-go/solid/isp/data"
	"github.com/q10357/solid-principles-go/solid/isp/data/machine"
)

/*
Interface Integration Principle -- dont put too much n an interface
You want to make an interface that allow people to build different
machines, constructs for working on a document
ex. printing, sending, scanning etc
*/

/*
What if some printer dont have fax method? Lets split up this interface
type Machine interface {
	Print(d Document)
	Fax(d Document)
	Scan(d Document)
}*/

func main() {
	//No functionality here, just to see it in action
	pc := machine.PhotoCopier{}
	pc.Print(&data.Document{})
}
