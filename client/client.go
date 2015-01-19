package main

import (
    "github.com/ddliu/go-spider"
    "time"
)

func main() {
    go TestServer()

    time.Sleep(1 * time.Second)

    client, err := spider.NewRPCClient(":1234", 10 * time.Second)
    if err != nil {
        panic(err)
    }
    client.Echo("...")
}

func TestServer() {
    spider.StartRPCServer(nil, ":1234")
}