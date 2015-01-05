package pipe

import (
    "github.com/ddliu/go-spider"
)

type Pipe func (*spider.Spider, *spider.Task)

func Join(pipes Pipe...) Pipe {

}

func Parallel(pipes Pipe...) Pipe {

}

func Series(pipes Pipe...) Pipe {

}
