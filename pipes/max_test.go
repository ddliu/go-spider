package pipes

import (
    "testing"
    "github.com/ddliu/go-spider"
)

func TestMax(t *testing.T) {
    s := spider.NewSpider().
        Pipe(Max(10))

    for i := 0; i < 100; i++ {
        s.AddUri("test")
    }

    s.Run()

    if s.Stats[spider.DONE] != 10 || s.Stats[spider.IGNORED] != 90 {
        t.Error("Max not working: ", s.Stats)
    }
}