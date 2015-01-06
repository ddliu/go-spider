// Copyright 2015 Liu Dong <ddliuhb@gmail.com>.
// Licensed under the MIT license.

package pipes

import (
    "sync"
    "github.com/ddliu/go-spider"
)

func Unique() {
    m := make(map[string]interface{})
    var l sync.Mutex
    return func(s *spider.Spider, t *spider.Task) {
        if check() {
            t.Ignore("Duplicated task")
        }
    }

    var check = func(uri string) bool {
        l.Lock()
        defer l.Unlock()
        if _, ok := m[uri]; !ok {
            m[uri] = nil

            return false
        } else {
            return true
        }
    }
}