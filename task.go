// Copyright 2015 Liu Dong <ddliuhb@gmail.com>.
// Licensed under the MIT license.

package spider

type Status int

const (
    PENDING Status = iota
    WORKING
    FAILED
    IGNORED
    DONE
)

type Data map[string]interface{}

type Task struct {
    Uri string
    Status Status
    Data Data
    Spider *Spider
    Parent *Task
}

func (this *Task) IsEnded() bool {
    return this.Status == FAILED || this.Status == IGNORED || this.Status == DONE
}

func (this *Task) Fail(reason interface{}) {
    this.Spider.FailTask(this, reason)
}

func (this *Task) Ignore(reason interface{}) {
    this.Spider.IgnoreTask(this, reason)
}

func (this *Task) Done() {
    this.Spider.DoneTask(this)
}

func (this *Task) Start() {
    this.Spider.StartTask(this)
}

// Create a new task from it
func (this *Task) Fork(uri string, data Data) {
    task := &Task {
        Uri: uri,
        Status: PENDING,
        Data: data,
        Parent: this,
    }

    this.Spider.AddTask(task)
}