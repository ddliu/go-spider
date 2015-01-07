package pipes

import (
    "testing"
    "github.com/ddliu/go-spider"
)

func TestRequest(t *testing.T) {
    s := spider.NewSpider().
        Pipe(Request).
        Pipe(func(s *spider.Spider, t *spider.Task) {
            t.Data.MustGetBytes("content")
        })

    shouldProduceUriList(t, s, []string{
        "http://httpbin.org",
        "http://httpbin.org/status/200",
        "http://httpbin.org/status/404",
    }, []string {
        "http://httpbin.org",
        "http://httpbin.org/status/200",
    }, "")
}