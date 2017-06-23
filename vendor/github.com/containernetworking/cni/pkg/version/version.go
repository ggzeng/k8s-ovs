// Copyright 2016 CNI authors
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

package version

// Current reports the version of the CNI spec implemented by this library
func Current() string {
	return "0.2.0"
}

// Legacy PluginInfo describes a plugin that is backwards compatible with the
// CNI spec version 0.1.0.  In particular, a runtime compiled against the 0.1.0
// library ought to work correctly with a plugin that reports support for
// Legacy versions.
//
// Any future CNI spec versions which meet this definition should be added to
// this list.
var Legacy = PluginSupports("0.1.0", "0.2.0")
