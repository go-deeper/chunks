# chunks

[![Go Reference](https://pkg.go.dev/badge/github.com/go-deeper/chunks.svg)](https://pkg.go.dev/github.com/go-deeper/chunks)
[![Go Report Card](https://goreportcard.com/badge/github.com/go-deeper/chunks)](https://goreportcard.com/report/github.com/go-deeper/chunks)
[![Test](https://github.com/go-deeper/chunks/actions/workflows/test.yaml/badge.svg)](https://github.com/go-deeper/chunks/actions/workflows/test.yaml)
[![CodeQL](https://github.com/go-deeper/chunks/actions/workflows/codeql-analysis.yml/badge.svg)](https://github.com/go-deeper/chunks/actions/workflows/codeql-analysis.yml)

Package chunks allows to split a slice into chunks.

## Install

```shell
go get github.com/go-deeper/chunks
```

## Usage

```go
package main

import (
	"fmt"

	"github.com/go-deeper/chunks"
)

func main() {
	slice := []int64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	sliceChunks := chunks.Split(slice, 4)
	fmt.Println(sliceChunks) // Output: [[1 2 3] [4 5 6] [7 8 9 10]]
}
```
