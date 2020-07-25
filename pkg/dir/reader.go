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
	"errors"
	"github.com/apex/log"
	"github.com/zpxio/fsel/pkg/session"
	"os"
	"path/filepath"
)

type Reader struct {
	dir           string
	maxDepth      int
	activeSession *session.Session
}

func CreateReader(dir string, s *session.Session) *Reader {
	r := Reader{
		dir:           dir,
		maxDepth:      1,
		activeSession: s,
	}

	return &r
}

func (r *Reader) Read() error {
	// Ensure the directory exists
	if !FileExists(r.dir) {
		return errors.New("directory does not exist")
	}

	filepath.Walk(r.dir, r.visit)

	return nil
}

func (r *Reader) visit(path string, f os.FileInfo, err error) error {

	// Check the depth
	depth := DepthFrom(r.dir, path)
	if f.IsDir() && depth == r.maxDepth {
		log.Debugf("Max depth reached: %s", path)
		return filepath.SkipDir
	}

	if f.IsDir() {
		return nil
	}

	item := session.Item{
		Path: path,
		Info: f,
	}

	r.activeSession.Add(item)

	return nil
}
