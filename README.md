gotrimet
========

[TriMet API](http://developer.trimet.org/) consumer in Go

## Usage

### Install

	$ go get github.com/juniorrobot/gotrimet

### Getting Started

You'll need to [register for an AppID](http://developer.trimet.org/appid/registration/).

Simply construct a `Client` with the AppID and use the available services:

```go
package main

import (
    "flag"
    "log"

    "github.com/juniorrobot/gotrimet"
)

var appID = flag.String("appID", "", "TriMet API app ID")

func main() {
    flag.Parse()
    tm := trimet.NewClient(*appID, nil)
    request := &trimet.ArrivalsRequest{
        LocationIDs: []int{10775, 8989},
        Streetcar: true,
    }
    if response, err := tm.Arrivals.Get(request); nil != err {
        log.Panic(err)
    } else {
        for index, arrival := range response.Arrivals {
            log.Printf("Arrival %v:\n\t%+v\n", index, arrival)
        }
    }
}
```

The tests provide more examples.

### Service support
BETA web services are not yet supported.

Until the Trip Planner supports a JSON response, I don't anticipate adding support
for it.

## About

Created and maintained by [John Robinson](http://github.com/juniorrobot).

Distributed under the MIT License.
