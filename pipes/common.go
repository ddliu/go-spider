package pipes

import (
    "github.com/ddliu/go-spider"
)

func Join(pipes Pipe...) Pipe {

}

func Parallel(pipes Pipe...) Pipe {
    return func(s *spider.Spider, t *spider.Task) {
        chs := make()
    }
}

func Series(pipes Pipe...) Pipe {
    return func(s *spider.Spider, t *spider.Task) {
        for i := 0; i < len(pipes); i++ {
            pipes[i](s, t)    
        }
    }
}
