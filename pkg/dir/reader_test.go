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
	"github.com/apex/log"
	"github.com/zpxio/fsel/pkg/session"
	"os"
	"path/filepath"
)

func testDir(subdirs ...string) string {
	cwd, _ := os.Getwd()

	dir := filepath.Join(filepath.Dir(filepath.Dir(cwd)), "testdata")

	for _, sub := range subdirs {
		dir = filepath.Join(dir, sub)
	}

	return dir
}

func (s *DirSuite) TestRead_D1Full() {
	dir := testDir("unit", "pkg", "dir", "read-001")
	log.Infof("Testing with directory: %s", dir)

	x := session.NewSession()
	r := CreateReader(dir, x)

	err := r.Read()

	s.NoError(err)

	s.Len(x.Files, 3)
}

func (s *DirSuite) TestRead_D255Full() {
	dir := testDir("unit", "pkg", "dir", "read-001")
	log.Infof("Testing with directory: %s", dir)

	x := session.NewSession()
	r := CreateReader(dir, x)
	r.SetDepth(255)

	err := r.Read()

	s.NoError(err)

	s.Len(x.Files, 13)
}

func (s *DirSuite) TestRead_NoDir() {
	dir := testDir("unit", "pkg", "dir", "fail-001")
	log.Infof("Testing with directory: %s", dir)

	x := session.NewSession()
	r := CreateReader(dir, x)

	err := r.Read()

	s.Error(err)
}

func (s *DirSuite) TestSetDepth_Simple() {
	dir := testDir("unit", "pkg", "dir", "test-001")

	x := session.NewSession()
	r := CreateReader(dir, x)

	testDepth := 4

	r.SetDepth(testDepth)
	s.Equal(testDepth, r.maxDepth)
}

func (s *DirSuite) TestSetDepth_Zero() {
	dir := testDir("unit", "pkg", "dir", "test-001")

	x := session.NewSession()
	r := CreateReader(dir, x)
	r.SetDepth(4)

	testDepth := 0

	r.SetDepth(testDepth)
	s.Equal(DefaultMaxDepth, r.maxDepth)
}

func (s *DirSuite) TestSetDepth_SubZero() {
	dir := testDir("unit", "pkg", "dir", "test-001")

	x := session.NewSession()
	r := CreateReader(dir, x)
	r.SetDepth(4)

	testDepth := -2

	r.SetDepth(testDepth)
	s.Equal(DefaultMaxDepth, r.maxDepth)
}
