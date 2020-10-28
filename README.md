# RAWG Video Games Database Golang client

[![Build Status](https://travis-ci.org/dimuska139/rawg-sdk-go.svg?branch=master)](https://travis-ci.org/dimuska139/rawg-sdk-go)
[![codecov](https://codecov.io/gh/dimuska139/rawg-sdk-go/branch/master/graph/badge.svg)](https://codecov.io/gh/dimuska139/rawg-sdk-go)
[![Go Report Card](https://goreportcard.com/badge/github.com/dimuska139/rawg-sdk-go)](https://goreportcard.com/report/github.com/dimuska139/rawg-sdk-go)
[![License](https://img.shields.io/github/license/mashape/apistatus.svg)](https://github.com/dimuska139/rawg-sdk-go/blob/master/LICENSE)
[![Mentioned in Awesome Go](https://awesome.re/mentioned-badge.svg)](https://github.com/avelino/awesome-go) 

This is unofficial RAWG SDK GO. This library contains methods for interacting with [RAWG API](https://rawg.io/).

## Installation

```shell
go get github.com/dimuska139/rawg-sdk-go
```

## Usage

```go
package main

import (
    "fmt"
    "net/http"
    "strings"
    "github.com/dimuska139/rawg-sdk-go"
)

func main() {
    config := rawg.Config{
        ApiKey:  "YourAppName", // 'YourAppName' will be set as User-Agent header
        Language: "ru",
        Rps:      5,
    }
    client := rawg.NewClient(http.DefaultClient, &config)
    
    filter := rawg.NewGamesFilter().
        SetSearch("Gta5").
        SetPage(1).
        SetPageSize(10).
        ExcludeCollection(1).
        WithoutParents()
    
    data, total, err := client.GetGames(filter)

    ...
}
```

The tests should be considered a part of the documentation. Also you can read [official docs](https://rawg.io/apidocs).

## API limitations

Only 5 requests per second allowed from one IP. The "Rps" parameter in configuration performs this limitation. 
So you don't worry about it.

## License

RAWG SDK GO is released under the
[MIT License](http://www.opensource.org/licenses/MIT).