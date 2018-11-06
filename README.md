# EasyCron [![GoDoc](https://godoc.org/github.com/tj/go-dropbox?status.svg)](https://godoc.org/github.com/anastalaz/easycron) ![](https://img.shields.io/badge/license-MIT-blue.svg) <a href='https://github.com/jpoles1/gopherbadger' target='_blank'>![gopherbadger-tag-do-not-edit](https://img.shields.io/badge/coverage-77%25-brightgreen.svg?longCache=true&style=flat)</a> [![Go Report Card](https://goreportcard.com/badge/github.com/anastalaz/easycron)](https://goreportcard.com/report/github.com/anastalaz/easycron)


 Simple [EasyCron](https://www.easycron.com) client for Go.

## About

 Modelled more or less 1:1 with the API for consistency and parity with the [official documentation](https://www.easycron.com/document).

## Installation

```console
$ go get github.com/anastalaz/easycron
```

## Example

```
package main

import (
    "fmt"
	"net/url"

    "github.com/anastalaz/easycron"
)

func main() {
	client := easycron.New(easycron.NewConfig("EASYCRON_ACCESS_TOKEN"))

	// Get list of all cronjobs
	v := url.Values{}
	v.Set("size", "1")
	myCronJobs, err := client.List(v) // If no optional parameters pass nil
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(myCronJobs)
}
```
