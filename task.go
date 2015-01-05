package spider

type Status int

import (
    "github.com/golang/glog"
)

const (
    PENDING Status = iota
    WORKING
    FAILED
    IGNORED
    DONE
)

type Task struct {
    Uri string
    Status Status
    Data map [string]interface{}
}

func (this *Task) Fail(reason string) {
    this.Status = FAILED
    glog.Warning("Task failed: ", this.Uri, "\t", reason)
}

func (this *Task) Ignore(reason string) {
    this.Status = IGNORED
    glog.Info("Task ignored: ", this.Uri, "\t", reason)
}

func (this *Task) Done() {
    this.Status = DONE
    glog.Info("Task done: ", this.Uri)
}

func (this *Task) Start() {
    this.Status = WORKING
    glog.Info("Task started: ", this.Uri)
}