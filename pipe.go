package spider

type Pipe interface {
    Run(*Task)
}

func IfPipe() Pipe {

}

func RetryPipe() Pipe {

}

func TimeoutPipe() Pipe {

}

func RequestPipe() Pipe {

}

func QueuePipe(pipes ...Pipe) Pipe {

}

func ConcurrentPipe(pipes ...Pipe) Pipe {

}