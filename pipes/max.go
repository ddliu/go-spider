package pipes

import (
    "fmt"
    "github.com/ddliu/go-spider"
)

func Max(n uint64) spider.Pipe {
    var counter uint64
    return func(s *spider.Spider, t *spider.Task) {
        if counter >= n {
            t.Ignore(fmt.Sprintf("Maximum task number of %d exceeded", n))
        } else {
            counter++
        }
    }
}