# go-spider [![Build Status](https://travis-ci.org/ddliu/go-spider.png)](https://travis-ci.org/ddliu/go-spider) [![GoDoc](https://godoc.org/github.com/ddliu/go-spider?status.svg)](https://godoc.org/github.com/ddliu/go-spider)

A flexible spider as well as a general perposed task runner.

## Usage

```go
import (
    "github.com/ddliu/go-spider"
    "github.com/ddliu/go-spider/pipes"
)

func main() {
    s := spider.NewSpider()
    s.Concurrency = 3

    s.
        Pipe(pipes.PipeA).
        Pipe(pipes.PipeB).
        Run()
}
```

## Example

See spider folder:

```
cd spider
mkdir download
go run main.go -depth=3 -max=100 -follow=http://tooling.github.io/book-of-modern-frontend-tooling/* -target=./download/ http://tooling.github.io/book-of-modern-frontend-tooling/
```