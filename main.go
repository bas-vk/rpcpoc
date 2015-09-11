package main

import (
	"net"

	"time"

	"github.com/bas-vk/rpcpoc/rpc"
	"github.com/bas-vk/rpcpoc/rpc/jsonrpc"
	"golang.org/x/net/context"
)

type DummyRpcService struct {
}

func (s *DummyRpcService) Hello() string {
	return "hi"
}

type EchoArgs struct {
	Input string `json:"input"`
	IntInput int `json:"inputi"`
}

func (s *DummyRpcService) Echo(args EchoArgs) string {
	return args.Input
}

func (s *DummyRpcService) EchoWithContext(ctx context.Context, args EchoArgs) string {
	return args.Input
}

func startUnixServer() {
	l, e := net.Listen("unix", "/tmp/test.sock")
	if e != nil {
		panic(e)
	}

	svr := rpc.NewServer()
	cm := new(DummyRpcService)
	svr.RegisterName("service", cm)

	go func() {
		for {
			c, err := l.Accept()
			if err != nil {
				panic(e)
			}

			codec := jsonrpc.NewServerCodec(c)
			go svr.ServeCodec(codec)
		}
	}()
}

func main() {
	startUnixServer()

	for {
		<-time.After(time.Second)
	}
}
