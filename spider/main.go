// Copyright 2015 Liu Dong <ddliuhb@gmail.com>.
// Licensed under the MIT license.

package main

import (
    "flag"
    "time"
    "strings"
    "os"
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
        }
    }

    app.Action = func(c *cli.Context) {
        var downloadPath = c.String("target")
        s := spider.NewSpider().
            Pipe(pipes.NormalizeUrl).
            Pipe(pipes.IfUrl(c.String("download"), pipes.Download(func(uri string) string {
                return downloadPath + "/" + Path(uri)
            }, 0777), nil)).
            Pipe(pipes.IfUrl(c.String("follow"), nil, pipes.Ignore)).
            Pipe(pipes.Request).
            Pipe(pipes.FollowLinks)

        for u in range c.Args() {
            s.AddUri(u)
        }

        s.Run()
    }

    app.Run(os.Args)
}

// Makes a string safe to use as an url path, cleaned of .. and unsuitable characters
// see: https://github.com/kennygrant/sanitize/blob/master/sanitize.go
func Path(text string) string {
    // Start with lowercase string
    fileName := strings.ToLower(text)
    fileName = strings.Replace(fileName, "..", "", -1)
    fileName = path.Clean(fileName)
    fileName = strings.Trim(fileName, " ")

    // Replace certain joining characters with a dash
    seps, err := regexp.Compile(`[ &_=+:]`)
    if err == nil {
        fileName = seps.ReplaceAllString(fileName, "-")
    }

    // Flatten accents first
    fileName = Accents(fileName)

    // Remove all other unrecognised characters
    // we are very restrictive as this is intended for ascii url slugs
    legal, err := regexp.Compile(`[^\w\_\~\-\./]`)
    if err == nil {
        fileName = legal.ReplaceAllString(fileName, "")
    }

    // Remove any double dashes caused by existing - in name
    fileName = strings.Replace(fileName, "--", "-", -1)

    // NB this may be of length 0, caller must check
    return fileName
}