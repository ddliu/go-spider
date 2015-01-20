// Copyright 2015 Liu Dong <ddliuhb@gmail.com>.
// Licensed under the MIT license.

package pipes

import (
    "github.com/ddliu/go-spider"
)

func Ignore(s *spider.Spider, t *spider.Task) {
    t.Ignore(nil)
}