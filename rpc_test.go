package rpctest

import (
	"runtime"
	"testing"
	"time"

	"github.com/gsdocker/gslogger"
	"github.com/gsrpc/gorpc"
	"github.com/gsrpc/gorpc/tcp"
)

var eventLoop = gorpc.NewEventLoop(uint32(runtime.NumCPU()), 2048, 500*time.Millisecond)

var client tcp.Client

var server Server

var content = []byte("hello world")

var content2 = []byte("hello worldhello worldhello worldhello worldhello worldhello worldhello worldhello worldhello worldhello world")

func init() {
	gslogger.NewFlags(gslogger.ERROR | gslogger.INFO)

	var err error
	client, err = tcp.BuildClient(gorpc.BuildPipeline(eventLoop)).Connect("test")

	if err != nil {
		panic(err)
	}

	server = BindServer(0, client.Pipeline())
}

func BenchmarkCall0(t *testing.B) {

	for i := 0; i < t.N; i++ {
		server.Push(content)
	}

}

func BenchmarkCall1(t *testing.B) {

	for i := 0; i < t.N; i++ {
		server.Push(content)
	}

}
