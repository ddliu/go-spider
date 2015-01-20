// Copyright 2015 Liu Dong <ddliuhb@gmail.com>.
// Licensed under the MIT license.

package pipes

import (
    "fmt"
    "github.com/ddliu/go-spider"
    "github.com/ddliu/go-httpclient"
)

func Request(s *spider.Spider, t *spider.Task) {
    httpclient.Begin()

    // auto referer
    if t.Parent != nil {
        httpclient.WithHeader("Referer", t.Parent.Uri)
    }

    res, err := httpclient.Get(t.Uri, nil)

    if err != nil {
        panic(err)
    }

    if res.StatusCode != 200 {
        panic(fmt.Sprintf("GET %s with status code %d", t.Uri, res.StatusCode))
    }

    // check for redirects
    if res.Request.URL.String() != t.Uri {
        t.Uri = res.Request.URL.String()
    }

    t.Data["content"], err = res.ReadAll()
    if err != nil {
        panic(err)
    }
}