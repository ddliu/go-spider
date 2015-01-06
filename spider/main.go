// Copyright 2015 Liu Dong <ddliuhb@gmail.com>.
// Licensed under the MIT license.

package main

import (
    "flag"
    "time"
    "github.com/ddliu/go-spider"
    "github.com/ddliu/go-spider/pipes"
)

func main() {
    // enable logging...
    flag.Parse()

    s := spider.NewSpider()

    s.Concurrency = 3

    s.
        Pipe(func(s *spider.Spider, t *spider.Task) {
            time.Sleep(1000 * time.Millisecond)
        }).
        Pipe(pipes.Print)

    for i := 0; i < 1000; i++ {
        s.AddUri("http://google.com/").AddUri("http://github.com/")
    }
        
    s.Run()
}