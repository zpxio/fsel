/*
 * Copyright 2020 zpxio (Jeff Sharpe)
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

package dir

import (
	"os"
	"path/filepath"
)

func FileExists(path string) bool {
	_, err := os.Stat(path)

	if os.IsNotExist(err) {
		return false
	}

	return true
}

func DepthFrom(base string, target string) int {
	depth := 0

	base = filepath.Clean(base)
	target = filepath.Clean(target)

	for target != base {
		if target == "/" {
			return -1
		}
		target = filepath.Dir(target)
		depth += 1
	}

	return depth
}
