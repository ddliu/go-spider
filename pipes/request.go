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

    t.Data["content"], err = res.ReadAll()
    if err != nil {
        panic(err)
    }
}