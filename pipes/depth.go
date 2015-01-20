// Copyright 2015 Liu Dong <ddliuhb@gmail.com>.
// Licensed under the MIT license.

package pipes

import (
    "fmt"
    "github.com/ddliu/go-spider"
)

func Depth(n uint) spider.Pipe {
    return func(s *spider.Spider, t *spider.Task) {
        if t.Depth >= n {
            t.Ignore(fmt.Sprintf("Maximum depth %d exceeded", n))
        }
    }
}