// Copyright 2015 Liu Dong <ddliuhb@gmail.com>.
// Licensed under the MIT license.

package spider

import (
    "net"
    "net/rpc"
    "time"
    "runtime"
)

type SpiderInfo struct {
    StartTime time.Time
    MemoryUsage uint64
    Stats map[Status]uint64
    IsStopped bool
    IsPaused bool
}

func StartRPCServer(spider *Spider, listen string) error {
    rpc.Register(NewRPC(spider))

    l, e := net.Listen("tcp", listen)
    if e != nil {
        return e
    }

    rpc.Accept(l)

    return nil
}

func NewRPC(spider *Spider) *RPC {
    return &RPC{
        spider,
        time.Now(),
    }
}

type RPC struct {
    spider *Spider
    startTime time.Time
}

func (rpc *RPC) Pong(skip bool, message *string) error {
    *message = "pong"

    return nil
}

func (rpc *RPC) Pause(skip bool, ack *bool) error {
    rpc.spider.Pause()
    return nil
}

func (rpc *RPC) Resume(skip bool, ack *bool) error {
    rpc.spider.Resume()
    return nil
}

func (rpc *RPC) Stop(skip bool, ack *bool) error {
    rpc.spider.Stop()
    return nil
}

func (rpc *RPC) Add(uriList []string, ack *bool) error {
    rpc.spider.AddUri(uriList...)
    return nil
}

func (rpc *RPC) Info(skip bool, info *SpiderInfo) error {
    var memStats runtime.MemStats

    runtime.ReadMemStats(&memStats)

    *info = SpiderInfo {
        StartTime: rpc.startTime,
        MemoryUsage: memStats.Alloc,
        Stats: rpc.spider.Stats,
        IsStopped: rpc.spider.IsStopped,
        IsPaused: rpc.spider.IsPaused,
    }

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

func (client *RPCClient) Ping() error {
    var message string
    err := client.connection.Call("RPC.Pong", true, &message)
    if err == nil {
        println(message)
    }

    return err
} 

func (client *RPCClient) Pause() error {
    var ack bool
    return client.connection.Call("RPC.Pause", true, &ack)
}

func (client *RPCClient) Resume() error {
    var ack bool
    return client.connection.Call("RPC.Resume", true, &ack)
}

func (client *RPCClient) Stop() error {
    var ack bool
    return client.connection.Call("RPC.Stop", true, &ack)
}

func (client *RPCClient) Add(uriList ...string) error {
    var ack bool
    return client.connection.Call("RPC.Add", uriList, &ack)
}

func (client *RPCClient) Info() (SpiderInfo, error) {
    var info SpiderInfo
    err := client.connection.Call("RPC.Info", true, &info)
    return info, err
}