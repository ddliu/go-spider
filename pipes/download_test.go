package pipes

import (
    "testing"
    "os"
    "fmt"
    "crypto/md5"
    "github.com/ddliu/go-spider"
)

func mapFilename(uri string) string {
    return "/tmp/" + fmt.Sprintf("%x", md5.Sum([]byte(uri)))
}

func TestDownload(t *testing.T) {
    s := spider.NewSpider()
    s.
        Pipe(Download(mapFilename, 0777)).
        AddUri("http://httpbin.org").
        Run()

    path := mapFilename("http://httpbin.org")

    _, err := os.Stat(path)
    if err != nil {
        t.Error(err)
    }

    if s.Stats[spider.DONE] != 1 {
        t.Error("DONE != 1")
    }
}