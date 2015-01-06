// Copyright 2015 Liu Dong <ddliuhb@gmail.com>.
// Licensed under the MIT license.

package pipes

import (
    "github.com/ddliu/go-spider"
)

func Print(s *spider.Spider, t *spider.Task) {
    println(t.Uri)
}