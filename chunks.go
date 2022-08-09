// Copyright 2022 go-deeper. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package chunks allows to split a slice into chunks.
package chunks

import (
	"errors"
)

// Split splits a slice of any type into balanced chunks.
// Split returns nil if applied slice is empty or maxSize < 1.
// The size of the last chunk can be distinct from others.
// The last chunk can have distinct size from others.
func Split[S ~[]E, E any](slice S, maxSize int) []S {
	arrLen := len(slice)
	if arrLen == 0 || maxSize < 1 {
		return nil
	}
	if arrLen <= maxSize {
		return []S{slice}
	}

	chunkLen := (arrLen + maxSize - 1) / maxSize
	chunkSize := int(float64(arrLen)/float64(chunkLen) + 0.5)

	chunks := make([]S, chunkLen)
	chunkLen--
	var low, high int
	for i := 0; i < chunkLen; i++ {
		low = i * chunkSize
		high = low + chunkSize
		chunks[i] = slice[low:high]
	}
	chunks[chunkLen] = slice[high:]
	return chunks
}

// ErrBreak stops the iteration.
var ErrBreak = errors.New("break")

// SplitFunc works like Split, but calls cb for each chunk.
// SplitFunc stops the iteration if cb returns non-nil error.
// SplitFunc returns nil if cb returns ErrBreak, otherwise error will be returned.
func SplitFunc[S ~[]E, E any](slice S, maxSize int, cb func(chunk S) error) error {
	chunks := Split(slice, maxSize)
	if len(chunks) == 0 {
		return nil
	}

	for _, chunk := range chunks {
		if err := cb(chunk); err != nil {
			if errors.Is(err, ErrBreak) {
				return nil
			}
			return err
		}
	}
	return nil
}
