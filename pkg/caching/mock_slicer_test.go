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

package caching_test

import (
	"github.com/NVIDIA/vdisc/pkg/caching"
	"github.com/NVIDIA/vdisc/pkg/storage"
	"github.com/stretchr/testify/mock"
)

// mockSlicer is an autogenerated mock type for the Slicer type
type mockSlicer struct {
	mock.Mock
}

// Bsize provides a mock function with given fields:
func (_m *mockSlicer) Bsize() int64 {
	ret := _m.Called()

	var r0 int64
	if rf, ok := ret.Get(0).(func() int64); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int64)
	}

	return r0
}

// Slice provides a mock function with given fields: obj, offset
func (_m *mockSlicer) Slice(obj storage.Object, offset int64) caching.Slice {
	ret := _m.Called(obj, offset)

	var r0 caching.Slice
	if rf, ok := ret.Get(0).(func(storage.Object, int64) caching.Slice); ok {
		r0 = rf(obj, offset)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(caching.Slice)
		}
	}

	return r0
}
