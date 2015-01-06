package pipes

import (
    "github.com/ddliu/go-spider"
)

func Ignore(s *spider.Spider, t *spider.Task) {
    t.Ignore(nil)
}