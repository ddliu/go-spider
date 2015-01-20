// Copyright 2015 Liu Dong <ddliuhb@gmail.com>.
// Licensed under the MIT license.

package pipes

import (
    "regexp"
    "github.com/ddliu/go-spider"
    "github.com/PuerkitoBio/purell"
)

var absRegex *regexp.Regexp
const flags = purell.FlagsSafe | purell.FlagRemoveDotSegments | 
              purell.FlagRemoveFragment | purell.FlagRemoveDuplicateSlashes | 
              purell.FlagRemoveUnnecessaryHostDots |
              purell.FlagRemoveEmptyPortSeparator

func DoNormalizeUrl(url, base string) (string, error) {
    url, err := joinUrl(base, url)
    if err != nil {
        return "", err
    }

    return purell.NormalizeURLString(url, flags)
}

func NormalizeUrl(s *spider.Spider, t *spider.Task) {
    var base = ""
    if t.Parent != nil {
        base = t.Parent.Uri 
    }

    u, err := DoNormalizeUrl(t.Uri, base)
    if err != nil {
        panic(err)
    }

    t.Uri = u
}

func init() {
    absRegex = regexp.MustCompile(`^[a-z0-9]+:`)
}