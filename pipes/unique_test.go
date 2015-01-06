// Copyright 2015 Liu Dong <ddliuhb@gmail.com>.
// Licensed under the MIT license.

package pipes

import (
    "github.com/ddliu/go-spider"
    "testing"
)

func TestUnique(t *testing.T) {
    s := spider.NewSpider().Pipe(Unique())
    shouldProduceUriList(t, s, []string {
        "http://www.google.com/",
        "http://www.google.com/",
        "http://www.google.com/",
        "http://www.google.com/",
        "http://www.google.com/",
        "http://www.github.com/",
        "http://www.github.com/",
        "http://www.github.com/",
        "http://www.google.com/",
        "http://www.google.com/",
        "http://www.github.com/",
    }, []string {
        "http://www.google.com/", 
        "http://www.github.com/",
    }, "")
}