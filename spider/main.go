package main

import (
    "github.com/ddliu/go-spider"
    "github.com/ddliu/go-spider/pipe"
)

func main() {
    println("aa")

    new(spider.Spider)
        .Pipe(pipe.PrintPipe)
        .AddUri("http://google.com/")
        .Run()
}