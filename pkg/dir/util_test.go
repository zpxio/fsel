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
	"github.com/stretchr/testify/suite"
	"path/filepath"
	"testing"
)

type DirSuite struct {
	suite.Suite
}

func TestDirSuite(t *testing.T) {
	suite.Run(t, new(DirSuite))
}

func (s *DirSuite) TestDepthFrom_Simple() {
	base := "/opt/test"

	s.Equal(1, DepthFrom(base, filepath.Join(base, "a")))
	s.Equal(2, DepthFrom(base, filepath.Join(base, "a", "b")))
	s.Equal(4, DepthFrom(base, filepath.Join(base, "a", "b", "c", "d")))
}

func (s *DirSuite) TestDepthFrom_Zero() {
	base := "/opt/test"

	s.Zero(DepthFrom(base, base))
}

func (s *DirSuite) TestDepthFrom_NotChild() {
	base := "/opt/test"

	s.Equal(-1, DepthFrom(base, "/tmp"))
}

func (s *DirSuite) TestDepthFrom_BaseIsRoot() {
	base := "/"

	s.Equal(1, DepthFrom(base, "/tmp"))
	s.Equal(0, DepthFrom(base, "/"))
}

func (s *DirSuite) TestDepthFrom_Dirty() {
	base := "/tmp"

	s.Equal(1, DepthFrom(base, "/tmp/stuff/"))
	s.Equal(1, DepthFrom(base, "/tmp/stuff"))
	s.Equal(2, DepthFrom(base, "/tmp/stuff//things"))
}
