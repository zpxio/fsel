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
	"github.com/zpxio/fsel/pkg/core"
	"github.com/zpxio/fsel/pkg/session"
	"os"
	"path/filepath"
	"sync"
)

const (
	DefaultMaxDepth = 1
)

type Reader struct {
	dir           string
	maxDepth      int
	activeSession *session.Session
	filterChannel chan core.Item
}

func CreateReader(dir string, s *session.Session) *Reader {
	r := Reader{
		dir:           dir,
		maxDepth:      DefaultMaxDepth,
		activeSession: s,
	}

	return &r
}

func (r *Reader) SetDepth(depth int) {
	if depth < 1 {
		depth = DefaultMaxDepth
	}
	r.maxDepth = depth
}

func (r *Reader) Read() error {
	// Ensure the directory exists
	if !FileExists(r.dir) {
		return errors.New("directory does not exist")
	}

	log.Debugf("Reading directory: %s", r.dir)
	// Create the channel
	workers := 3
	r.filterChannel = make(chan core.Item, 20)
	wg := sync.WaitGroup{}
	wg.Add(workers)
	for i := 0; i < workers; i++ {
		go func() {
			for fi := range r.filterChannel {
				log.Debugf("Filtering item: %s", fi.Path)

				// Do filter

				r.activeSession.Add(fi)
			}

			wg.Done()
		}()
	}

	walkErr := filepath.Walk(r.dir, r.visit)
	close(r.filterChannel)
	if walkErr != nil {
		return walkErr
	}
	log.Debugf("Done reading directory: %s", r.dir)

	wg.Wait()
	log.Debugf("All filtering is complete.")

	return nil
}

func (r *Reader) visit(path string, f os.FileInfo, err error) error {
	if err != nil {
		return err
	}

	// Check the depth
	depth := DepthFrom(r.dir, path)
	if f.IsDir() && depth == r.maxDepth {
		log.Debugf("Max depth reached: %s", path)
		return filepath.SkipDir
	}

	if f.IsDir() {
		return nil
	}

	item := core.Item{
		Path: path,
		Info: f,
	}

	r.filterChannel <- item

	return nil
}
