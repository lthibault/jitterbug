# jitterbug

Tickers with random jitter

[![Godoc Reference](https://img.shields.io/badge/godoc-reference-blue.svg?style=flat-square)](https://godoc.org/github.com/lthibault/jitterbug)
[![Go Report Card](https://goreportcard.com/badge/github.com/SentimensRG/ctx?style=flat-square)](https://goreportcard.com/report/github.com/lthibault/jitterbug)

## Installation

```bash
go get -u github.com/lthibault/jitterbug
```

## Usage

Jitterbug is used by instantiating a `jitterbug.Ticker` with an interval and a
`jitterbug.Jitter`.  The former specifies a baseline interval for the ticker,
to which a jitter is added by the latter.

```go
package main

import (
    "log"

    "github.com/lthibault/jitterbug"
)

func main() {
    t := jitterbug.Ticker{
        Interval: time.Millisecond * 300,
        Jitter: &jitterbug.Norm{ Stdev: time.Millisecond * 100},
    }

    // NOTE:  jitterbug.Ticker.C is a method that returns a channel, not a
    //        concrete channel (unlike time.Ticker)
    for tick := <- range t.C() {
        log.Println(tick)
    }
}

```

## RFC

If you find this useful please let me know:  <l.thibault@sentimens.com>

Seriously, even if you just used it in your weekend project, I'd like to hear
about it :)
