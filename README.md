# go-spider [![Build Status](https://travis-ci.org/ddliu/go-spider.png)](https://travis-ci.org/ddliu/go-spider) [![GoDoc](https://godoc.org/github.com/ddliu/go-spider?status.svg)](https://godoc.org/github.com/ddliu/go-spider)

A flexible spider as well as a general perposed task runner.

## Go Package Dependencies

See Godeps/Godeps.json

## Usage

### Workflow

```go
package main

import (
    // Import the main package
    "github.com/ddliu/go-spider"

    // Import a lot of useful pipes
    "github.com/ddliu/go-spider/pipes"
)

func main() {
    // Create a spider
    s := spider.NewSpider()

    // Config it
    s.Concurrency = 3

    // Combine pipes together
    s.
        Pipe(pipes.PipeA).
        Pipe(pipes.PipeB)

    // Let's go!
    s.Run()
}
```

## Take Control on the Fly

Start the server:

```
s := spier.NewSpider()
// setup
s.RunAndServe("127.0.0.1:1234")
```

Control it with the client:

```
export GO_SPIDER_SERVER=127.0.0.1:1234
go run client/client.go info
go run client/client.go add uri1 uri2 uri3
go run client/client.go pause
go run client/client.go resume
go run client/client.go watch
```

### A Watch Example

```
Status: Running , time: 15s memory: 264KB
>>>>>>>>........................................................................
Total: 100, pending: 86, working: 3
Done: 7, failed: 4, ignored: 0
```

## Examples

See `examples/downloader` folder:

```
cd examples/downloader
mkdir download
go run main.go -depth=3 -max=100 -follow=http://tooling.github.io/book-of-modern-frontend-tooling/* -target=./download/ http://tooling.github.io/book-of-modern-frontend-tooling/
```