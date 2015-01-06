// Copyright 2015 Liu Dong <ddliuhb@gmail.com>.
// Licensed under the MIT license.

package pipes

import (
    "github.com/ddliu/go-spider"
    "testing"
)

func TestIfUri(t *testing.T) {
    s := spider.NewSpider().Pipe(IfUri("http://*.google.com/*", nil, Ignore))
    shouldProduceUriList(t, s, []string {
        "http://www.google.com/a.html",
        "http://drive.google.com/b.html",
        "http://google.com/c.html",
        "http://www.github.com/",
    }, []string {
        "http://www.google.com/a.html",
        "http://drive.google.com/b.html",
    }, "")
}

func TestIfUriRegexp(t *testing.T) {
    s := spider.NewSpider().Pipe(IfUriRegexp(`.*.google.com/[ab].*`, nil, Ignore))
    shouldProduceUriList(t, s, []string {
        "http://www.google.com/a.html",
        "http://drive.google.com/b.html",
        "http://docs.google.com/c.html",
        "http://google.com/c.html",
        "http://www.github.com/",
    }, []string {
        "http://www.google.com/a.html",
        "http://drive.google.com/b.html",
    }, "")
}