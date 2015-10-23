package main

import (
	"runtime"
	"time"

	"github.com/gschat/rpctest"
	"github.com/gsdocker/gslogger"
	"github.com/gsrpc/gorpc"
	"github.com/gsrpc/gorpc/tcp"
)

type _MockServer struct{}

func (mock *_MockServer) Push(content []byte) error {
	return nil
}

func main() {

	gslogger.NewFlags(gslogger.ERROR | gslogger.INFO)

	var eventLoop = gorpc.NewEventLoop(uint32(runtime.NumCPU()), 2048, 500*time.Millisecond)

	tcp.NewServer(
		gorpc.BuildPipeline(eventLoop),
	).EvtNewPipeline(tcp.EvtNewPipeline(func(pipeline gorpc.Pipeline) {
		pipeline.AddService(rpctest.MakeServer(0, &_MockServer{}))
	})).Listen(":13512")
}
