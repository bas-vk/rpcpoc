package main

import (
	"fmt"
	"math/big"
	"net"

	"time"

	"github.com/bas-vk/rpcpoc/rpc"
	"github.com/bas-vk/rpcpoc/rpc/jsonrpc"
)

// representation of core.ChainManager
type ChainManager struct {
}

func (cm *ChainManager) SomeNonExposedMethod() uint64 {
	return 123456
}

func (cm *ChainManager) LatestBlockNumber() BlockNumber {
	return NewBlockNumber(392398)
}

func (cm *ChainManager) GetBlock(number BlockNumber) (*Block, error) {
	return &Block{
		Number:   number,
		Hash:     "0x9392392323232",
		GasLimit: big.NewInt(39293932),
		GasUsed:  big.NewInt(39823923932),
		Nonce:    3023992332,
	}, nil
}

// Proxy objects will wrap the core types and exposes the RPC methods.
type ChainManagerProxy struct {
	cm *ChainManager // the proxy would call methods on this instance
}

func (p *ChainManagerProxy) LatestBlockNumber() (BlockNumber, error) {
	return p.cm.LatestBlockNumber(), nil
}

func (p *ChainManagerProxy) GetBlockByNumber(number BlockNumber) (*Block, error) {
	return p.cm.GetBlock(number)
}

type unexportedType struct{}

func (p *ChainManagerProxy) SkipMethod(val unexportedType) {
}

type EchoResponse struct {
	A    int
	B    *int
	Str1 string
	Str2 *string
}

func (p *ChainManagerProxy) Echo(a int, b *int, str1 string, str2 *string) EchoResponse {
	return EchoResponse{a, b, str1, str2}
}

func (p *ChainManagerProxy) EchoWithError(a int, b *int, str1 string, str2 *string) (int, *int, string, *string, error) {
	return a, b, str1, str2, fmt.Errorf("error %d/%d/%s/%s", a, *b, str1, *str2)
}

func (p *ChainManagerProxy) GasPrice() GasValue {
	return big.NewInt(39293)
}

func startUnixServer() {
	l, e := net.Listen("unix", "/tmp/test.sock")
	if e != nil {
		panic(e)
	}

	svr := rpc.NewServer()
	cm := new(ChainManagerProxy)
	svr.RegisterName("chainmanager", cm)

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
