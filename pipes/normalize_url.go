package pipes

import (
    "regexp"
    "strings"
    pkgurl "net/url"
    "github.com/ddliu/go-spider"
    "github.com/PuerkitoBio/purell"
)

var absRegex *regexp.Regexp
const flags = purell.FlagsSafe | purell.FlagRemoveDotSegments | 
              purell.FlagRemoveFragment | purell.FlagRemoveDuplicateSlashes | 
              purell.FlagRemoveUnnecessaryHostDots |
              purell.FlagRemoveEmptyPortSeparator

func DoNormalizeUrl(url, base string) (string, error) {
    if base == "" {
        return purell.NormalizeURLString(url, flags)
    }

    baseUrl, err := pkgurl.Parse(base)
    if err != nil {
        return "", err
    }

    // TODO: not strict
    if strings.HasPrefix(url, "//") {
        url = baseUrl.Scheme + ":" + url
    } else if strings.HasPrefix(url, "/") {
        url = baseUrl.Scheme + "://" + baseUrl.Host + url
    } else if !absRegex.MatchString(url) {
        url = base + url
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
    absRegex = regexp.MustCompile(`^[a-z0-9]+://`)
}