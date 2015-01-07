package pipes

import (
    "fmt"
    "io/ioutil"
    "os"
    "github.com/ddliu/go-spider"
    "github.com/ddliu/go-httpclient"
)

func Download(mapper func(string) string, perm os.FileMode) spider.Pipe {
    return func(s *spider.Spider, t *spider.Task) {
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

        content, err := res.ReadAll()
        if err != nil {
            panic(err)
        }
        
        if err := ioutil.WriteFile(mapper(t.Uri), content, perm); err != nil {
            panic(err)
        }
    }
}