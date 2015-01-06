// Copyright 2015 Liu Dong <ddliuhb@gmail.com>.
// Licensed under the MIT license.

package pipes

import (
    "github.com/ddliu/go-spider"
    "testing"
)

func TestUnique(t *testing.T) {
    var result []string
    spider.NewSpider().
        Pipe(Unique()).
        Pipe(func(s *spider.Spider, t *spider.Task) {
            result = append(result, t.Uri)
        }).
        AddUri(
            "http://www.google.com/",
            "http://www.google.com/",
            "http://www.google.com/",
            "http://www.google.com/",
            "http://www.google.com/",
            "http://www.github.com/",
            "http://www.github.com/",
            "http://www.github.com/",
            "http://www.google.com/",
            "http://www.google.com/",
            "http://www.github.com/",
        ).
        Run()

    if !sliceEquals(result, []string{"http://www.google.com/", "http://www.github.com/"}) {
        t.Error(result)
    }
}

func sliceEquals(s1, s2 []string) bool {
    if len(s1) != len(s2) {
        return false
    }

    for _, v1 := range s1 {
        found := false
        for _, v2 := range s2 {
            if v1 == v2 {
                found = true
                break
            }
        }

        if !found {
            return false
        }
    }

    return true
}