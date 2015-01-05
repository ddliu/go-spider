package spider

type Pipe func (*Spider, *Task) {

}



type Pipe interface {
    Run(*Task)
}

type FuncPipe struct {

}

func (this *FuncPipe) Run()

func FuncPipe()

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