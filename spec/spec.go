/*
 Copyright 2020 Qiniu Cloud (qiniu.com)

 Licensed under the Apache License, Version 2.0 (the "License");
 you may not use this file except in compliance with the License.
 You may obtain a copy of the License at

     http://www.apache.org/licenses/LICENSE-2.0

 Unless required by applicable law or agreed to in writing, software
 distributed under the License is distributed on an "AS IS" BASIS,
 WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 See the License for the specific language governing permissions and
 limitations under the License.
*/

package spec

import (
	exec "github.com/qiniu/goplus/exec/bytecode"
)

// -----------------------------------------------------------------------------

// A Context represents the context of an executor.
type Context = exec.Context

// -----------------------------------------------------------------------------

// GoPackage represents a Go package.
type GoPackage = exec.GoPackage

// NewGoPackage creates a new builtin Go Package.
func NewGoPackage(pkgPath string) *GoPackage {
	return exec.NewGoPackage(pkgPath)
}

// ToStrings converts []interface{} into []string.
func ToStrings(args []interface{}) []string {
	ret := make([]string, len(args))
	for i, arg := range args {
		ret[i] = arg.(string)
	}
	return ret
}

// -----------------------------------------------------------------------------
