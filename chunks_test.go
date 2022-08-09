// Copyright 2022 go-deeper. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package chunks_test

import (
	"errors"
	"reflect"
	"testing"

	"github.com/go-deeper/chunks"
)

var tests = []struct {
	name  string
	slice []int
	max   int
	want  [][]int
}{
	{
		name:  "empty slice",
		slice: []int{},
		max:   1,
		want:  nil,
	},
	{
		name:  "max size is equal or less than zero",
		slice: []int{1},
		max:   0,
		want:  nil,
	},
	{
		name:  "max size is equal or more than slice size",
		slice: []int{1, 2},
		max:   2,
		want:  [][]int{{1, 2}},
	},
	{
		name:  "balanced chunks",
		slice: []int{1, 2, 3, 4},
		max:   3,
		want:  [][]int{{1, 2}, {3, 4}},
	},
	{
		name:  "smaller last chunk",
		slice: []int{1, 2, 3, 4, 5},
		max:   3,
		want:  [][]int{{1, 2, 3}, {4, 5}},
	},
	{
		name:  "bigger last chunk",
		slice: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		max:   4,
		want:  [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9, 10}},
	},
}

func TestSplit(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Act
			got := chunks.Split(tt.slice, tt.max)

			// Assert
			if !reflect.DeepEqual(tt.want, got) {
				t.Errorf("Split() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSplitFunc(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Arrange
			var got [][]int
			fn := func(chunk []int) error {
				got = append(got, chunk)
				return nil
			}

			// Act
			err := chunks.SplitFunc(tt.slice, tt.max, fn)

			// Assert
			if err != nil {
				t.Errorf("SplitFunc() err = %v, want nil", err)
			}
			if !reflect.DeepEqual(tt.want, got) {
				t.Errorf("SplitFunc() got = %v, want %v", got, tt.want)
			}
		})
	}

	t.Run("with break", func(t *testing.T) {
		// Arrange
		slice := []int{1, 2, 3}
		want := [][]int{{1}, {2}}
		var got [][]int
		fn := func(chunk []int) error {
			if chunk[0] == 3 {
				return chunks.ErrBreak
			}
			got = append(got, chunk)
			return nil
		}

		// Act
		err := chunks.SplitFunc(slice, 1, fn)

		// Assert
		if err != nil {
			t.Errorf("SplitFunc() err = %v, want nil", err)
		}
		if !reflect.DeepEqual(want, got) {
			t.Errorf("SplitFunc() got = %v, want %v", got, want)
		}
	})

	t.Run("with error", func(t *testing.T) {
		// Arrange
		slice := []int{1, 2, 3}
		want := [][]int{{1}, {2}}
		wantErr := errors.New("error")
		var got [][]int
		fn := func(chunk []int) error {
			if chunk[0] == 3 {
				return wantErr
			}
			got = append(got, chunk)
			return nil
		}

		// Act
		err := chunks.SplitFunc(slice, 1, fn)

		// Assert
		if err != wantErr {
			t.Errorf("SplitFunc() err = %v, want %v", err, wantErr)
		}
		if !reflect.DeepEqual(want, got) {
			t.Errorf("SplitFunc() got = %v, want %v", got, want)
		}
	})
}
