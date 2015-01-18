package spider

import (
    "net"
    "net/rpc"
    "log"
)

func StartRPCServer(spider *Spider, listen string) {
    rpc.Register(NewRPC(spider))

    l, e := net.Listen("tcp", listen)
    if e != nil {
        log.Fatal("Start RPC error:", e)
    }

    rpc.Accept(l)
}

func NewRPC(spider *Spider) *RPC {
    return &RPC{
        spider,
    }
}

type RPC struct {
    spider *Spider
}

func (rpc *RPC) Echo(message string, ack *string) error {
    *ack = "Hello: " + *ack

    return nil
}