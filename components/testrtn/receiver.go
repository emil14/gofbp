package testrtn

import (
	"fmt"

	"github.com/jpaulm/gofbp/core"
)

type Receiver struct {
	ipt core.InputConn
}

func (receiver *Receiver) Setup(p *core.Process) {
	receiver.ipt = p.OpenInPort("IN")
	p.MustRun = true
}

func (receiver *Receiver) Execute(p *core.Process) {
	fmt.Println(p.Name + " started")

	for {
		var pkt = p.Receive(receiver.ipt)
		if pkt == nil {
			break
		}
		fmt.Println("Output: ", pkt.Contents)
		p.Discard(pkt)
	}

	fmt.Println(p.Name + " ended")
}
