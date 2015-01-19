package spider

import (
    "net"
    "net/rpc"
    "log"
    "time"
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

// test
func (rpc *RPC) Echo(message string, ack *string) error {
    *ack = "Hello: " + message

    return nil
}

func (rpc *RPC) Pause(skip bool, ack *bool) error {
    return nil
}

func (rpc *RPC) Resume(skip bool, ack *bool) error {
    return nil
}


func NewRPCClient(dsn string, timeout time.Duration) (*RPCClient, error) {
    connection, err := net.DialTimeout("tcp", dsn, timeout)
    if err != nil {
        return nil, err
    }

    return &RPCClient{connection: rpc.NewClient(connection)}, nil
}

type RPCClient struct {
    connection *rpc.Client
}

func (client *RPCClient) Echo(message string) error {
    err := client.connection.Call("RPC.Echo", message, &message)
    if err == nil {
        println(message)
    }

    return err
} 
