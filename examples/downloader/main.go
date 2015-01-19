// Copyright 2015 Liu Dong <ddliuhb@gmail.com>.
// Licensed under the MIT license.

package main

import (
    "os"
    "regexp"
    "strings"

    "github.com/ddliu/go-spider"
    "github.com/ddliu/go-spider/pipes"
    "github.com/codegangsta/cli"
)

func main() {
    app := cli.NewApp()
    app.Name = "spider"
    app.Usage = "download links"

    app.Flags = []cli.Flag {
        cli.StringFlag {
            Name: "download, d",
            Value: "*",
            Usage: "which url to download",
        },
        cli.StringFlag {
            Name: "follow",
            Value: "*",
            Usage: "which url to follow",
        },
        cli.StringFlag {
            Name: "target, t",
            Value: ".",
            Usage: "target folder to save downloaded files",
        },
        cli.IntFlag {
            Name: "concurrency, c",
            Value: 5,
            Usage: "spider concurrency",
        },
        cli.IntFlag {
            Name: "max, m",
            Value: 0,
            Usage: "maximum urls to follow",
        },
        cli.IntFlag {
            Name: "depth, i",
            Value: 0,
            Usage: "maximum depth to follow",
        },
    }

    app.Action = func(c *cli.Context) {
        var downloadPath = c.String("target")
        var normPathRe = regexp.MustCompile(`[:/\\'"_\?\[\]\&_]+`)
        const SEPERATOR = "-"
        s := spider.NewSpider().
            Pipe(pipes.NormalizeUrl).
            Pipe(pipes.Unique()).
            Pipe(pipes.IfUri(c.String("follow"), nil, pipes.Ignore)).
            Pipe(pipes.IfUri(c.String("download"), pipes.Download(func(uri string) string {
                return strings.Trim(downloadPath + "/" + normPathRe.ReplaceAllLiteralString(uri, SEPERATOR), SEPERATOR)
            }, 0777), nil)).
            Pipe(pipes.Request).
            Pipe(pipes.FollowLinks)

        for _, u := range c.Args() {
            s.AddUri(u)
        }

        if len(c.Args()) == 0 {
            cli.ShowAppHelp(c)
            return
        }

        s.Run()
    }

    app.Run(os.Args)
}