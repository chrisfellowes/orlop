// Copyright (c) 2021 Ketch Kloud, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

package errors

import (
	"go.ketch.com/lib/orlop/v2/errors/internal"
	"net/http"
)

type notFoundError struct {
	error
}

func (te notFoundError) Unwrap() error {
	return te.error
}

func (te notFoundError) NotFound() bool {
	return true
}

// NotFound returns a not found error
func NotFound(err error) error {
	return &notFoundError{err}
}

// IsNotFound returns true if the error is a NotFound error
func IsNotFound(err error) bool {
	var missing internal.NotFound
	var sc internal.StatusCode
	var ec internal.ErrorCode

	if As(err, &missing) && missing.NotFound() {
		return true
	}
	if As(err, &sc) && sc.StatusCode() == http.StatusNotFound {
		return true
	}
	if As(err, &ec) && ec.ErrorCode() == ENOTFOUND {
		return true
	}

	return false
}