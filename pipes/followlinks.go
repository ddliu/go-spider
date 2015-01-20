// Copyright 2015 Liu Dong <ddliuhb@gmail.com>.
// Licensed under the MIT license.

package pipes

import (
    "regexp"
    "github.com/ddliu/go-spider"
    "github.com/ddliu/go-requery"
)

var regexpLink *regexp.Regexp
var regexpBase *regexp.Regexp
func FollowLinks(s *spider.Spider, t *spider.Task) {
    content := t.Data.MustGetBytes("content")

    base := requery.Find(content, regexpBase).Sub(1).String()

    // parse baseurl
    requery.FindAll(content, regexpLink).Each(func(c *requery.Context) {
        url, err := joinUrl(base, c.Sub(1).String())
        if err == nil {
            t.ForkUri(url)
        }
    })
}

func init() {
    regexpLink = regexp.MustCompile(`(?i)<a\s+[^<]*href="([^"]+)"`)
    regexpBase = regexp.MustCompile(`(?i)<base\s+href="([^"]+)"`)
}