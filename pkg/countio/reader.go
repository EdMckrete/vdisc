// Copyright © 2019 NVIDIA Corporation
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package countio

import (
	"io"
	"sync/atomic"
)

type Reader struct {
	r  io.Reader
	pn *int64
	n  int64
}

func NewReader(r io.Reader, counter ...*int64) *Reader {
	return &Reader{
		r: r,
	}
}

func NewReaderWithAtomicCounter(r io.Reader, counter *int64) *Reader {
	return &Reader{
		r:  r,
		pn: counter,
	}
}

func (cr *Reader) Read(p []byte) (n int, err error) {
	n, err = cr.r.Read(p)
	cr.n += int64(n)
	if cr.pn != nil {
		atomic.AddInt64(cr.pn, int64(n))
	}
	return
}

func (cr *Reader) BytesRead() int64 {
	return cr.n
}
