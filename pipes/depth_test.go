// Copyright 2015 Liu Dong <ddliuhb@gmail.com>.
// Licensed under the MIT license.

package pipes

import (
    "testing"
    "github.com/ddliu/go-spider"
)

func TestDepth(t *testing.T) {
    s := spider.NewSpider().
        Pipe(Depth(3)).
        Pipe(func(s *spider.Spider, t *spider.Task) {
            t.ForkUri("child1")
            t.ForkUri("child2")
            t.ForkUri("child3")
        })

    for i := 0; i < 10; i++ {
        s.AddUri("test")
    }

    s.Run()

    // 0    1   2   3   4
    // 10 * 3 * 3 * 3 * 3
    // 10  30   90  270 810
    // 10  40   130 400 1210

    if s.Stats[spider.DONE] != 130 || s.Stats[spider.IGNORED] != 400 - 130 {
        t.Error("Depth not working: ", s.Stats)
    }
}