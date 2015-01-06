// Copyright 2015 Liu Dong <ddliuhb@gmail.com>.
// Licensed under the MIT license.

package pipes

import (
    "sync"
    "github.com/ddliu/go-spider"
)

func Unique() spider.Pipe {
    m := make(map[string]interface{})
    var l sync.Mutex
    
    var insert = func(uri string) bool {
        l.Lock()
        defer l.Unlock()
        if _, ok := m[uri]; !ok {
            m[uri] = nil

            return true
        } else {
            return false
        }
    }
    
    return func(s *spider.Spider, t *spider.Task) {
        if !insert(t.Uri) {
            t.Ignore("Duplicated task")
        }
    }

}