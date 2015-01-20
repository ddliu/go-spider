// Copyright 2015 Liu Dong <ddliuhb@gmail.com>.
// Licensed under the MIT license.

package main

import (
    "fmt"
    "time"
    "log"
    "os"
    "github.com/ddliu/go-spider"
    "github.com/codegangsta/cli"
    "github.com/dustin/go-humanize"
    tm "github.com/buger/goterm"
)

const (
    VERSION = "0.1.0"
)

type FormatedInfo struct {
    DisplayStatus string
    DisplayTime string
    DisplayMemory string

    Pending uint64
    Working uint64
    Failed uint64
    Ignored uint64
    Done uint64
    Total uint64
}

func formateInfo(info spider.SpiderInfo) FormatedInfo {

    pending, _ := info.Stats[spider.PENDING]
    working, _ := info.Stats[spider.WORKING]
    failed, _ := info.Stats[spider.FAILED]
    ignored, _ := info.Stats[spider.IGNORED]
    done, _ := info.Stats[spider.DONE]
    total := pending + working + failed + ignored + done
    
    var status string

    if info.IsStopped {
        status = "Stopped"
    } else if info.IsPaused {
        status = "Paused"
    } else if pending == 0 && working == 0 {
        status = "Finished"
    } else {
        status = "Running"
    }

    displayMemory := humanize.Bytes(info.MemoryUsage)
    displayTime := (time.Duration(time.Since(info.StartTime).Seconds()) * time.Second).String()

    return FormatedInfo {
        DisplayStatus: status,
        DisplayTime: displayTime,
        DisplayMemory: displayMemory,
        Pending: pending,
        Working: working,
        Failed: failed,
        Ignored: ignored,
        Done: done,
        Total: total,
    }
}

func main() {
    app := cli.NewApp()
    app.Name = "gospider"
    app.Usage = "The go-spider client (https://github.com/ddliu/go-spider)"
    app.Version = VERSION
    app.Author = "Liu Dong"
    app.Email = "ddliuhb@gmail.com"

    app.Flags = []cli.Flag {
        cli.StringFlag{
            Name: "server, s",
            Usage: "Server IP and port to connect(127.0.0.1:1234)",
            EnvVar: "GO_SPIDER_SERVER",
        },
    }

    app.Commands = []cli.Command{
        {
            Name: "watch",
            Usage: "Keep watching the spider",
            Action: doWatch,
        },
        {
            Name: "info",
            Usage: "Show spider info",
            Action: doInfo,
        },
        {
            Name: "add",
            Usage: "Add tasks",
            Action: doAdd,
        },
        {
            Name: "pause",
            Usage: "Pause the spider",
            Action: doPause,
        },
        {
            Name: "resume",
            Usage: "Resume the spider",
            Action: doResume,
        },
        {
            Name: "stop",
            Usage: "Stop the spider",
            Action: doStop,
        },
        {
            Name: "ping",
            Usage: "Ping the RPC service",
            Action: doPing,
        },
    }

    app.Run(os.Args)
}

func getClient(c *cli.Context) *spider.RPCClient {
    s := c.GlobalString("server")
    if s == "" {
        log.Fatal("--server not specified")
    }

    client, err := spider.NewRPCClient(s, 10 * time.Second)
    if err != nil {
        log.Fatal(err)
    }

    return client
}

func loopGetClient(c *cli.Context) *spider.RPCClient {
    s := c.GlobalString("server")
    if s == "" {
        log.Fatal("--server not specified")
    }

    for {
        client, err := spider.NewRPCClient(s, 10 * time.Second)
        if err == nil {
            return client
        }

        time.Sleep(3 * time.Second)
    }

    return nil
}

func doInfo(c *cli.Context) {
    client := getClient(c)
    info, err := client.Info()
    if err != nil {
        log.Fatal(err)
    }

    formatedInfo := formateInfo(info)

    println("Start Time:", humanize.Time(info.StartTime))
    println("Status:", formatedInfo.DisplayStatus)
    println("Memory:", formatedInfo.DisplayMemory)
    println("Total:", formatedInfo.Total)
    println("Done:", formatedInfo.Done)
    println("Pending:", formatedInfo.Pending)
    println("Working:", formatedInfo.Working)
    println("Failed:", formatedInfo.Failed)
    println("ignored:", formatedInfo.Ignored)
}

func doAdd(c *cli.Context) {
    tasks := c.Args()
    if len(tasks) == 0 {
        log.Fatal("No task added")
    }

    client := getClient(c)

    err := client.Add(tasks...)
    if err != nil {
        log.Fatal(err)
    }

    println("Added", len(tasks), "tasks")
}

func doWatch(c *cli.Context) {
    client := getClient(c)

    tm.Clear()

    for {
        info, err := client.Info()
        if err == nil {
            formatedInfo := formateInfo(info)
            // By moving cursor to top-left position we ensure that console output
            // will be overwritten each time, instead of adding new.
            tm.MoveCursor(1,1)

            tm.Println("Status:", formatedInfo.DisplayStatus, ", time:", formatedInfo.DisplayTime, "memory:", formatedInfo.DisplayMemory, "     ")

            if formatedInfo.Total > 0 {
                width := tm.Width()
                if width < 10 {
                    width = 10
                }

                if width > 80 {
                    width = 80
                }

                finished := formatedInfo.Done + formatedInfo.Failed + formatedInfo.Ignored
                var finishedProgressNumber int
                if finished > 0 {
                    finishedProgressNumber = int(uint64(width) * finished / formatedInfo.Total)
                }

                progressString := ""
                for i := 0; i < width; i++ {
                    if i < finishedProgressNumber {
                        progressString += ">"
                    } else {
                        progressString += "."
                    }
                }

                tm.Println(progressString)
            }

            tm.Print(fmt.Sprintf("Total: %d, pending: %d, working: %d     \nDone: %d, failed: %d, ignored: %d     \n", 
                formatedInfo.Total,
                formatedInfo.Pending,
                formatedInfo.Working,
                formatedInfo.Done,
                formatedInfo.Failed,
                formatedInfo.Ignored,
            ))

            tm.Flush() // Call it every time at the end of rendering
        } else {
            client = loopGetClient(c)
        }
        time.Sleep(time.Second)
    }
}

func doPause(c *cli.Context) {
    client := getClient(c)
    err := client.Pause()
    if err != nil {
        log.Fatal(err)
    }

    println("Paused")
}

func doResume(c *cli.Context) {
    client := getClient(c)
    err := client.Resume()
    if err != nil {
        log.Fatal(err)
    }

    println("Resumed")
}

func doStop(c *cli.Context) {
    println("stop")
}

func doPing(c *cli.Context) {
    client := getClient(c)
    err := client.Ping()
    if err != nil {
        log.Fatal(err)
    }
}