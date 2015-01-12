package pipes

import (
    "github.com/ddliu/go-spider"
    "github.com/ddliu/go-requery"
)

func FollowLinks(s *spider.Spider, t *spider.Task) {
    content := t.Data.MustGetBytes("content")
    requery.FindAll(content, `<a\s+[^<]*href="([^"]+)"`).Each(func(c *requery.Context) {
        t.ForkUri(c.Sub(1).String())
    })
}