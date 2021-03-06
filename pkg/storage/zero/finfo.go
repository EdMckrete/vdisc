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

package zerodriver

import (
	"fmt"
	"os"
	"time"
)

type finfo struct {
	size int64
}

// base name of the file
func (fi *finfo) Name() string {
	return fmt.Sprintf("%d", fi.size)
}

// length in bytes for regular files; system-dependent for others
func (fi *finfo) Size() int64 {
	return fi.size
}

// file mode bits
func (fi *finfo) Mode() os.FileMode {
	return 0444
}

// modification time
func (fi *finfo) ModTime() time.Time {
	return time.Unix(0, 0).UTC()
}

// abbreviation for Mode().IsDir()
func (fi *finfo) IsDir() bool {
	return false
}

// underlying data source (can return nil)
func (fi *finfo) Sys() interface{} {
	return nil
}
