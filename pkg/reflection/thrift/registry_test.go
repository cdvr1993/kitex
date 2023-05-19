/*
 * Copyright 2023 CloudWeGo Authors
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package thriftreflection

import (
	"testing"

	"github.com/cloudwego/thriftgo/reflection"

	"github.com/cdvr1993/kitex/internal/test"
)

func TestRegistry(t *testing.T) {
	GlobalFiles.Register(&reflection.FileDescriptor{
		Filename: "testa",
	})
	test.Assert(t, GlobalFiles.GetFileDescriptors()["testa"] != nil)

	f := NewFiles()
	f.Register(&reflection.FileDescriptor{
		Filename: "testb",
	})
	test.Assert(t, f.GetFileDescriptors()["testb"] != nil)
}
