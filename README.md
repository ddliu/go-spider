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

## RPC Server

There is a builtin RPC server which makes it easy to integrite with other systems.

```
s := spier.NewSpider()

// setup
// ...

s.RunAndServe("127.0.0.1:1234")
```

## Client

After starting the spider and the RPC server, you can take control easily with the cli client.

### Install

```
go get github.com/ddliu/go-spider/gospider
```

### Usage

```
$ gospider

NAME:
   gospider - The go-spider client (https://github.com/ddliu/go-spider)

USAGE:
   gospider [global options] command [command options] [arguments...]

VERSION:
   0.1.0

AUTHOR:
  Liu Dong - <ddliuhb@gmail.com>

COMMANDS:
   watch    Keep watching the spider
   info     Show spider info
   add      Add tasks
   pause    Pause the spider
   resume   Resume the spider
   stop     Stop the spider
   ping     Ping the RPC service
   help, h  Shows a list of commands or help for one command
   
GLOBAL OPTIONS:
   --server, -s     Server IP and port to connect(127.0.0.1:1234) [$GO_SPIDER_SERVER]
   --help, -h       show help
   --version, -v    print the version
```

### The Watch Example

```
$ export GO_SPIDER_SERVER=127.0.0.1:1234
$ gospider watch
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