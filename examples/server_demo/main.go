// Copyright 2015 Liu Dong <ddliuhb@gmail.com>.
// Licensed under the MIT license.

package main

import (
    "time"
    "math/rand"
    "github.com/ddliu/go-spider"
    "strconv"
)

func main() {
    rand.Seed(time.Now().UnixNano())
    s := spider.NewSpider()
    s.Pipe(func(s *spider.Spider, t *spider.Task) {
        time.Sleep((time.Duration(rand.Intn(5)) + 3) * time.Second)
        if rand.Intn(10) == 0 {
            panic("fail it")
        }
        if rand.Intn(10) == 1 {
            panic("ignore it")
        }
    })

    for i := 1; i <= 100; i++ {
        u := strconv.Itoa(i)
        s.AddUri(u)
    }

    s.RunAndServe(":1234")
}