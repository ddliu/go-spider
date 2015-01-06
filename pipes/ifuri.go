// Copyright 2015 Liu Dong <ddliuhb@gmail.com>.
// Licensed under the MIT license.

package pipes

import (
    "strings"
    "regexp"
    "github.com/ddliu/go-spider"
)

func IfUri(pattern string, ifPipe, elsePipe spider.Pipe) spider.Pipe {
    // Expand * with regexp
    if strings.Contains(pattern, "*") {
        pattern = regexp.QuoteMeta(pattern)
        pattern = "^" + strings.Replace(pattern, `\*`, ".*", -1) + "$"
        return IfUriRegexp(pattern, ifPipe, elsePipe)
    }

    // match exactly
    return func(s *spider.Spider, t *spider.Task) {
        if t.Uri == pattern {
            if ifPipe != nil {
                ifPipe(s, t)
            }
        } else {
            if elsePipe != nil {
                elsePipe(s, t)
            }
        }
    }
}

func IfUriRegexp(regexpString string, ifPipe, elsePipe spider.Pipe) spider.Pipe {
    r := regexp.MustCompile(regexpString)
    return func(s *spider.Spider, t *spider.Task) {
        if r.MatchString(t.Uri) {
            if ifPipe != nil {
                ifPipe(s, t)
            }
        } else {
            if elsePipe != nil {
                elsePipe(s, t)
            }
        }
    }
}