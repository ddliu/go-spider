package spider

import (
    "sync"
)

type Spider struct {
    Concurrency int
    pipes []Pipe
    tasks []*Task
    m sync.Mutex
}

func (this *Spider) Pipe(pipe) *Spider {
    if this.pipes == nil {
        this.pipes = make([]Pipe)
    }

    return this
} 

// Run spider
func (this *Spider) Run() {
    for {

    }
}

// Run tasks in a pool
func (this *Spider) pool() {
    for {
        this.m.Lock()
        if len(this.tasks) {

        }
        this.m.Unlock()
    }
}

// Run a single task
func (this *Spider) do(task *Task) {
    for i := 0; i < len(this.pipes); i++ {
        this.pipes[i].
    }
}

func (this *Spider) RunForever() {

}

func (this *Spider) AddUri(uri string) *Spider {
    return this.AddTask(&Task{Uri: uri})
}

func (this *Spider) AddTask(task *Task) *Spider {
    this.m.Lock()
    this.tasks = append(this.tasks, task)
    this.m.Unlock()

    return this
}