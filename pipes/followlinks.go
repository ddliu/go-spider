package pipes

import (
    "bytes"
    "github.com/ddliu/go-spider"
    "golang.org/x/net/html"
)

func FollowLinks(s *spider.Spider, t *spider.Task) {
    content := t.MustGetBytes("content")
    doc, err := html.Parse(bytes.NewReader(content))
    if err != nil {
        panic(err)
    }

    var f func(*html.Node)
    f = func(n *html.Node) {
        if n.Type == html.ElementNode && n.Data == "a" {
            for _, a := range n.Attr {
                if a.Key == "href" {
                    t.ForkUri(a.Value)
                    break
                }
            }
        }
        for c := n.FirstChild; c != nil; c = c.NextSibling {
            f(c)
        }
    }
    f(doc)
}