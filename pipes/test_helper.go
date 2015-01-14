package pipes

import (
    "testing"
    "github.com/ddliu/go-spider"
)

func shouldProduceUriList(t *testing.T, s *spider.Spider, input []string, expected []string, message string) {
    var result []string
    s.Pipe(func(s *spider.Spider, t *spider.Task) {
        result = append(result, t.Uri)
    }).AddUri(input...).Run()


    if !sliceEquals(expected, result) {
        t.Error(message, " ", result, " != ", expected)
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