# go-spider

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