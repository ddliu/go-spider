// Copyright 2015 Liu Dong <ddliuhb@gmail.com>.
// Licensed under the MIT license.

package pipes

import (
    "strings"
    pkgurl "net/url"
)

func joinUrl(base, url string) (string, error) {
    if base == "" {
        return url, nil
    }

    baseUrl, err := pkgurl.Parse(base)
    if err != nil {
        return "", err
    }

    // TODO: not strict
    if strings.HasPrefix(url, "//") {
        url = baseUrl.Scheme + ":" + url
    } else if strings.HasPrefix(url, "/") {
        url = baseUrl.Scheme + "://" + baseUrl.Host + url
    } else if !absRegex.MatchString(url) {
        urlInfo, err := pkgurl.Parse(url)
        if err != nil {
            return "", err
        }
        url = baseUrl.ResolveReference(urlInfo).String()
    }

    return url, nil
}