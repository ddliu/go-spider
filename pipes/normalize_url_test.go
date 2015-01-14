package pipes

import (
    "testing"
    "github.com/ddliu/go-spider"
)

func TestNormalizeUrl(t *testing.T) {
    s := spider.NewSpider().
        Pipe(NormalizeUrl).
        Pipe(func(s *spider.Spider, t *spider.Task) {
            if t.Uri == "http://www.google.com/main/" {
                t.ForkUri(
                    "a.html",
                    "http://drive.google.com/",
                    "../b.html",
                    "/c/d.html",
                    "mailto:test@test.com",
                )
            }
        })

    shouldProduceUriList(t, s, []string{"http://www.google.com/main/"}, []string {
        "http://www.google.com/main/",
        "http://www.google.com/main/a.html",
        "http://drive.google.com/",
        "http://www.google.com/b.html",
        "http://www.google.com/c/d.html",
    }, "")
}