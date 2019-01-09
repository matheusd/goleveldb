// Copyright (c) 2012, Suryandaru Triandana <syndtr@gmail.com>
// All rights reserved.
//
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

// +build wasm

package storage

import (
	"fmt"
	"os"
)

func newFileLock(path string, readOnly bool) (fl fileLock, err error) {
	return nil, fmt.Errorf("unsupported platform")
}

func setFileLock(f *os.File, readOnly, lock bool) error {
	return fmt.Errorf("unsupported platform")
}

func rename(oldpath, newpath string) error {
	return fmt.Errorf("unsupported platform")
}

func isErrInvalid(err error) bool {
	return false
}

func syncDir(name string) error {
	return fmt.Errorf("unsupported platform")
}
