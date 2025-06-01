# chunks

[![Go Reference](https://pkg.go.dev/badge/github.com/go-deeper/chunks.svg)](https://pkg.go.dev/github.com/go-deeper/chunks)
[![Go Report Card](https://goreportcard.com/badge/github.com/go-deeper/chunks)](https://goreportcard.com/report/github.com/go-deeper/chunks)
[![Go](https://github.com/go-deeper/chunks/actions/workflows/go.yaml/badge.svg)](https://github.com/go-deeper/chunks/actions/workflows/go.yaml)
[![CodeQL](https://github.com/go-deeper/chunks/actions/workflows/codeql.yml/badge.svg)](https://github.com/go-deeper/chunks/actions/workflows/codeql.yml)

Package chunks allows to split a slice of any type into chunks with approximately equals sum of values.

## Install

```shell
go get github.com/go-deeper/chunks
```

## Usage

### Split slice into chunks

```go
package main

import (
	"fmt"

	"github.com/go-deeper/chunks"
)

func main() {
	slice := []int64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	sliceChunks := chunks.Split(slice, 9)
	fmt.Println(sliceChunks)
}
```

Output:

```text
[[1 2 3 4 5] [6 7 8 9 10]]
```

### Iterate over chunks using a callback function

```go
package main

import (
	"fmt"

	"github.com/go-deeper/chunks"
)

func main() {
	slice := []int64{1, 2, 3, 4, 5, 6}
	err := chunks.SplitFunc(slice, 5, func(chunk []int64) error {
		fmt.Println(chunk)
		return nil
	})
	fmt.Println(err)
}
```

Output:

```text
[1 2 3]
[4 5 6]
<nil>
```

### A callback function returns error

```go
package main

import (
	"errors"
	"fmt"

	"github.com/go-deeper/chunks"
)

func main() {
	slice := make([]int64, 10)
	err := chunks.SplitFunc(slice, 1, func(_ []int64) error {
		return errors.New("some error")
	})
	fmt.Println(err)
}
```

Output:

```text
some error
```

### A callback function stops the iterations without error

```go
package main

import (
	"fmt"

	"github.com/go-deeper/chunks"
)

func main() {
	slice := make([]string, 10)
	err := chunks.SplitFunc(slice, 1, func(_ []string) error {
		return chunks.ErrBreak
	})
	fmt.Println(err)
}
```

Output:

```text
<nil>
```
