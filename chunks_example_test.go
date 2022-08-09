// Copyright 2022 go-deeper. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package chunks_test

import (
	"errors"
	"fmt"

	"github.com/go-deeper/chunks"
)

func ExampleSplit() {
	slice := []int64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	arrChunks := chunks.Split(slice, 4)
	fmt.Println("chunks:", arrChunks)

	// Output:
	// chunks: [[1 2 3] [4 5 6] [7 8 9 10]]
}

func ExampleSplitFunc() {
	slice := []int64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	err := chunks.SplitFunc(slice, 4, func(chunk []int64) error {
		fmt.Println("chunk:", chunk)
		return nil
	})
	fmt.Println("error:", err)

	// Output:
	// chunk: [1 2 3]
	// chunk: [4 5 6]
	// chunk: [7 8 9 10]
	// error: <nil>
}

func ExampleSplitFunc_withBreak() {
	slice := []int64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	var i int
	err := chunks.SplitFunc(slice, 4, func(chunk []int64) error {
		if i == 2 {
			return chunks.ErrBreak
		}
		i++

		fmt.Println("chunk:", chunk)
		return nil
	})
	fmt.Println("error:", err)

	// Output:
	// chunk: [1 2 3]
	// chunk: [4 5 6]
	// error: <nil>
}

func ExampleSplitFunc_withError() {
	slice := []int64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	var i int
	err := chunks.SplitFunc(slice, 4, func(chunk []int64) error {
		if i == 2 {
			return errors.New("some error")
		}
		i++

		fmt.Println("chunk:", chunk)
		return nil
	})
	fmt.Println("error:", err)

	// Output:
	// chunk: [1 2 3]
	// chunk: [4 5 6]
	// error: some error
}
