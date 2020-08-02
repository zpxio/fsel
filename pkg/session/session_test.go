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

package session

import (
	"github.com/stretchr/testify/suite"
	"os"
	"testing"
)

type SessionSuite struct {
	suite.Suite
}

func TestSessionSuite(t *testing.T) {
	suite.Run(t, new(SessionSuite))
}

func (s *SessionSuite) TestNewSession() {
	x := NewSession()

	s.NotNil(x)
	s.Empty(x.Files)
}

func (s *SessionSuite) TestAdd() {
	x := NewSession()
	s.Require().Empty(x.Files)

	info, _ := os.Stat("/")

	i1 := Item{
		Path: "/path/a",
		Info: info,
	}

	x.Add(i1)
	s.Len(x.Files, 1)

	i2 := Item{
		Path: "/path/b",
		Info: info,
	}

	x.Add(i2)
	s.Len(x.Files, 2)
}

func (s *SessionSuite) TestAdd_Duplicate() {
	x := NewSession()
	s.Require().Empty(x.Files)
	info, _ := os.Stat("/")

	i2 := Item{
		Path: "/path/b",
		Info: info,
	}

	x.Add(i2)
	s.Require().Len(x.Files, 1)

	i3 := Item{
		Path: "/path/b",
		Info: info,
	}

	x.Add(i3)
	s.Len(x.Files, 1)
}
