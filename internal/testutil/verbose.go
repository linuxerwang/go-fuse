// Copyright 2016 the Go-FUSE Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package testutil

import (
	"bytes"
	"flag"
	"runtime"
)

// VerboseTest returns true if the testing framework is run with -v.
func VerboseTest() bool {
	var buf [2048]byte
	n := runtime.Stack(buf[:], false)
	if bytes.Index(buf[:n], []byte("TestNonVerbose")) != -1 {
		return false
	}

	flag := flag.Lookup("test.v")
	return flag != nil && flag.Value.String() == "true"
}
