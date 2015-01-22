// Copyright 2015 Liu Dong <ddliuhb@gmail.com>.
// Licensed under the MIT license.

package spider

import (
    "time"
    // "fmt"
    "sync"
    "log"
)

const defaultConcurrency = 3

const (
    ON_START = iota
    ON_STOP = iota
)

type Listener func(*Spider, *Task)

// Create a spider.
func NewSpider() *Spider {
    spider := &Spider {
    }

    spider.prepare()
    return spider
}

type Spider struct {
    Concurrency int
    pipes []Pipe
    tasks []*Task
    events map[int][]Listener
    Stats map[Status]uint64
    IsPaused bool
    IsStopped bool
    IsDebug bool
    m sync.Mutex
    statsLock sync.Mutex
}

// Chain a pipe.
func (this *Spider) Pipe(pipe Pipe) *Spider {
    this.pipes = append(this.pipes, pipe)

    return this
} 

// Initialize the spider objects.
func (this *Spider) prepare() {
    this.events = make(map[int][]Listener)
    this.Stats = make(map[Status]uint64)
    this.Stats[PENDING] = 0
    this.Stats[WORKING] = 0
    this.Stats[FAILED] = 0
    this.Stats[IGNORED] = 0
    this.Stats[DONE] = 0
}

// Run spider forever, and accept a quit channel to close it.
// 
// Loop through the task list and run each of them with the help of a buffered channel.
func (this *Spider) RunForever(quit chan bool) {
    this.IsStopped = false
    this.IsPaused = false
    log.Println("[INFO] Spider started")

    if this.Concurrency <= 0 {
        this.Concurrency = defaultConcurrency
    }

    chs := make(chan bool, this.Concurrency)

    for {
        select {
        case <- quit:
            break
        default:
            // do nothing
        }

        if this.IsStopped {
            break
        }

        if this.IsPaused {
            time.Sleep(10 * time.Millisecond)
            continue
        }

        this.m.Lock()
        if len(this.tasks) > 0 {
            task := this.tasks[len(this.tasks) - 1]
            this.tasks = this.tasks[:len(this.tasks) - 1]

            this.m.Unlock()
            chs <- true
            go func() {
                this.do(task)
                <-chs
            }()
        } else {
            this.m.Unlock()

            // there is nothing to do, sleep for 10 ms
            time.Sleep(10 * time.Millisecond)
        }
    }
}

// Run spider and stop when complete.
func (this *Spider) Run() {
    quit := make(chan bool)
    go this.RunForever(quit)

    // check finish
    for {
        // if all tasks are finished, we can go out of the loop
        if this.IsFinished() {
            quit <- true
            log.Println("[INFO] Spider finished")
            break
        } else {
            time.Sleep(10 * time.Millisecond)
        }
    }
}

// Run spider and start a RPC server
func (this *Spider) RunAndServe(listen string) error {
    quit := make(chan bool)
    go this.RunForever(quit)

    return StartRPCServer(this, listen)
}

// Run a single task (should never panic)
func (this *Spider) do(task *Task) {
    defer func() {
        // error occured
        if r := recover(); r != nil {
            this.FailTask(task, r)
        } else if !task.IsEnded() {
            this.DoneTask(task)
        }

        // cleanup
        task.Parent = nil
        task.Data = nil
    }()

    this.StartTask(task)

    Series(this.pipes...)(this, task)
}

func (this *Spider) Pause() {
    this.IsPaused = true
}

func (this *Spider) Resume() {
    this.IsPaused = false
}

func (this *Spider) Stop() {
    this.IsStopped = true
}

// Check if all tasks have been processed.
func (this *Spider) IsFinished() bool {
    return this.Stats[PENDING] == 0 && this.Stats[WORKING] == 0
}

// Add tasks from uri.
func (this *Spider) AddUri(uris ...string) *Spider {
    for _, uri := range uris {
        this.AddTask(NewTask(uri))
    }

    return this
}

// Add a task to queue
func (this *Spider) AddTask(task *Task) *Spider {
    task.Spider = this
    this.m.Lock()
    this.tasks = append(this.tasks, task)
    this.m.Unlock()

    this.statsLock.Lock()
    this.Stats[PENDING]++
    this.statsLock.Unlock()

    return this
}

// Mark a task as failed.
func (this *Spider) FailTask(task *Task, reason interface{}) {
    task.Status = FAILED

    this.statsLock.Lock()
    this.Stats[FAILED]++
    this.Stats[WORKING]--
    this.statsLock.Unlock()

    log.Println("[WARN] Task failed: ", task.Uri, "\t", reason)
}

// Mark a task as done.
func (this *Spider) DoneTask(task *Task) {
    task.Status = DONE
    this.statsLock.Lock()
    this.Stats[DONE]++
    this.Stats[WORKING]--
    this.statsLock.Unlock()

    if this.IsDebug {
        log.Println("[DEBUG] Task done: ", task.Uri)
    }
}

// Mark a task as ignored.
func (this *Spider) IgnoreTask(task *Task, reason interface{}) {
    task.Status = IGNORED

    this.statsLock.Lock()
    this.Stats[IGNORED]++
    this.Stats[WORKING]--
    this.statsLock.Unlock()

    if this.IsDebug {
        log.Println("[DEBUG] Task ignored: ", task.Uri, "\t", reason)
    }
}

// Mark a task as started.
func (this *Spider) StartTask(task *Task) {
    task.Status = WORKING

    this.statsLock.Lock()
    this.Stats[WORKING]++
    this.Stats[PENDING]--
    this.statsLock.Unlock()

    if this.IsDebug {
        log.Println("[DEBUG] Task started: ", task.Uri)
    }
}

// Register events
func (this *Spider) On(e int, f Listener) *Spider {
    this.events[e] = append(this.events[e], f)

    return this
}

// Trigger an event
func (this *Spider) Trigger(e int, t *Task) {
    if events, ok := this.events[e]; ok {
        for _, e := range events {
            e(this, t)
        }
    }
}