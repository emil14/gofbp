package testing

import (
	"testing"

	"github.com/jpaulm/gofbp/components/websocket"
	"github.com/jpaulm/gofbp/core"
	"github.com/jpaulm/gofbp/testing/components"
)

func TestWebSocket(t *testing.T) {
	params, err := core.LoadXMLParams("../params.xml")
	if err != nil {
		panic(err)
	}
	net := core.NewNetwork("TestWebSocket")
	net.SetParams(params)
	path := "localhost:8080"
	proc1 := net.NewProc("WSRequest", &websocket.WSRequest{})
	proc2 := net.NewProc("AnsReq", &components.WSAnsReq{})
	proc3 := net.NewProc("WSRespond", &websocket.WSRespond{})
	net.Initialize(path, proc1, "ADDR")
	net.Connect(proc1, "OUT", proc2, "IN", 6)
	net.Connect(proc2, "OUT", proc3, "IN", 6)
	net.Run()
}
