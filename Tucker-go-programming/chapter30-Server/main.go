package main

import (
	"flag"
	"fmt"
	"log"

	gnet "github.com/panjf2000/gnet/v2"
)

type echoServer struct {
	gnet.BuiltinEventEngine

	eng       gnet.Engine
	addr      string
	multicore bool
}

func (es *echoServer) OnOpen(c gnet.Conn) (out []byte, action gnet.Action) {
	log.Printf("Client Connected: %s\n", c.RemoteAddr().String())
	return nil, gnet.None
}

func (es *echoServer) OnClose(c gnet.Conn, err error) (action gnet.Action) {
	log.Printf("Client Closed: %s\n", c.RemoteAddr().String())
	return gnet.None
}

func (es *echoServer) onBoot(eng gnet.Engine) (action gnet.Action) {
	es.eng = eng
	log.Printf("echo Server with multi=core=%t is listening on %s", es.multicore, es.addr)
	return gnet.None
}

func (es *echoServer) OnTraffic(c gnet.Conn) gnet.Action {
	log.Printf("Received from client: %s\n", c.RemoteAddr().String())
	buf, _ := c.Next(-1)
	log.Printf("%s", buf)
	c.Write(buf)
	return gnet.None
}

func main() {
	var port int
	var multicore bool

	flag.IntVar(&port, "port", 9000, "--port 9000")
	flag.BoolVar(&multicore, "multicore", false, "multicore true")
	flag.Parse()

	echo := &echoServer{
		addr:      fmt.Sprintf("tcp://:%d", port),
		multicore: multicore,
	}

	log.Fatal(gnet.Run(echo, echo.addr, gnet.WithMulticore(multicore)))
}
