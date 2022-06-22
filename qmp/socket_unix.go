// Copyright 2016 The go-qemu Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// +build !windows

package qmp

import (
	"fmt"
	"os"
	"time"

	"golang.org/x/sys/unix"
)

// NewPipeMonitor configures a connection to the provided QEMU monitor pipe.
// An error is returned if the pipe cannot be successfully dialed, or the
// dial attempt times out.
func NewPipeMonitor(addr string, timeout time.Duration) (*SocketMonitor, error) {
	return nil, fmt.Errorf("unimplemented")
}

func getUnixRights(file *os.File) []byte {
	return unix.UnixRights(int(file.Fd()))
}