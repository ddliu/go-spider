// Copyright 2015 Liu Dong <ddliuhb@gmail.com>.
// Licensed under the MIT license.

package main

import (
    "os"
    "github.com/ddliu/go-spider"
    "github.com/ddliu/go-spider/pipes"
    "github.com/codegangsta/cli"

    "strings"
    "bytes"
    "path"
    "regexp"
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
        s := spider.NewSpider().
            Pipe(pipes.NormalizeUrl).
            Pipe(pipes.Unique()).
            Pipe(pipes.IfUri(c.String("download"), pipes.Download(func(uri string) string {
                uri = strings.Replace(uri, "/", "-", -1)
                return downloadPath + "/" + Path(uri)
            }, 0777), nil)).
            Pipe(pipes.IfUri(c.String("follow"), nil, pipes.Ignore)).
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



// Makes a string safe to use as an url path, cleaned of .. and unsuitable characters
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

// Replace a set of accented characters with ascii equivalents.
func Accents(text string) string {
    // Replace some common accent characters
    b := bytes.NewBufferString("")
    for _, c := range text {
        // Check transliterations first
        if val, ok := transliterations[c]; ok {
            b.WriteString(val)
        } else {
            b.WriteRune(c)
        }
    }
    return b.String()
}

// A very limited list of transliterations to catch common european names translated to urls.
// This set could be expanded with at least caps and many more characters.
var transliterations = map[rune]string{
    'À': "A",
    'Á': "A",
    'Â': "A",
    'Ã': "A",
    'Ä': "A",
    'Å': "AA",
    'Æ': "AE",
    'Ç': "C",
    'È': "E",
    'É': "E",
    'Ê': "E",
    'Ë': "E",
    'Ì': "I",
    'Í': "I",
    'Î': "I",
    'Ï': "I",
    'Ð': "D",
    'Ł': "L",
    'Ñ': "N",
    'Ò': "O",
    'Ó': "O",
    'Ô': "O",
    'Õ': "O",
    'Ö': "O",
    'Ø': "OE",
    'Ù': "U",
    'Ú': "U",
    'Ü': "U",
    'Û': "U",
    'Ý': "Y",
    'Þ': "Th",
    'ß': "ss",
    'à': "a",
    'á': "a",
    'â': "a",
    'ã': "a",
    'ä': "a",
    'å': "aa",
    'æ': "ae",
    'ç': "c",
    'è': "e",
    'é': "e",
    'ê': "e",
    'ë': "e",
    'ì': "i",
    'í': "i",
    'î': "i",
    'ï': "i",
    'ð': "d",
    'ł': "l",
    'ñ': "n",
    'ń': "n",
    'ò': "o",
    'ó': "o",
    'ô': "o",
    'õ': "o",
    'ō': "o",
    'ö': "o",
    'ø': "oe",
    'ś': "s",
    'ù': "u",
    'ú': "u",
    'û': "u",
    'ū': "u",
    'ü': "u",
    'ý': "y",
    'þ': "th",
    'ÿ': "y",
    'ż': "z",
    'Œ': "OE",
    'œ': "oe",
}