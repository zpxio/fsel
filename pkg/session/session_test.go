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
	"errors"
	"fmt"
	"github.com/stretchr/testify/suite"
	"github.com/zpxio/fsel/pkg/core"
	"os"
	"testing"
)

type SessionSuite struct {
	suite.Suite
	Results []string
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

	i1 := core.Item{
		Path: "/path/a",
		Info: info,
	}

	x.Add(i1)
	s.Len(x.Files, 1)

	i2 := core.Item{
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

	i2 := core.Item{
		Path: "/path/b",
		Info: info,
	}

	x.Add(i2)
	s.Require().Len(x.Files, 1)

	i3 := core.Item{
		Path: "/path/b",
		Info: info,
	}

	x.Add(i3)
	s.Len(x.Files, 1)
}

type TestAction struct {
	Batch int
	Suite *SessionSuite
}

func (a *TestAction) Name() string {
	return "TestAction"
}

func (a *TestAction) BatchSize() int {
	return a.Batch
}

func (a *TestAction) Run(items []*core.Item) error {

	for _, i := range items {
		if i.Path == "" {
			return errors.New("empty path")
		}
		fmt.Printf("TEST ITEM %s\n", i.Path)
		a.Suite.Results = append(a.Suite.Results, i.Path)
	}

	return nil
}

func (s *SessionSuite) TestAddAction() {
	x := NewSession()

	s.Len(x.Actions, 0)

	x.AddAction(&TestAction{Batch: 4})

	s.Len(x.Actions, 1)

	x.AddAction(&TestAction{Batch: 12})

	s.Len(x.Actions, 2)
}

func (s *SessionSuite) TestEffectiveActions() {
	x := NewSession()

	s.Equal(DefaultActions, x.effectiveActions())

	ta := []Action{&TestAction{
		Batch: 10,
	}}

	x.AddAction(ta[0])

	s.Equal(ta, x.effectiveActions())
}

func (s *SessionSuite) TestRunActions() {
	x := NewSession()

	ta := TestAction{
		Batch: 1000,
		Suite: s,
	}

	defaultInfo, _ := os.Stat("/")
	filenames := []string{"/a/file001.txt", "/b/file002.txt", "/b/file003.txt"}

	x.AddAction(&ta)

	for _, fn := range filenames {
		x.Add(core.Item{
			Path: fn,
			Info: defaultInfo,
		})
	}
	s.Len(x.Files, len(filenames))

	s.Results = make([]string, 0)
	x.RunActions()

	s.Len(s.Results, len(filenames))
	s.Equal(filenames, s.Results)
}

func (s *SessionSuite) TestRunActions_Failure() {
	x := NewSession()

	var testResults []string
	ta := TestAction{
		Batch: 1000,
		Suite: s,
	}

	defaultInfo, _ := os.Stat("/")
	filenames := []string{"/a/file001.txt", "/b/file002.txt", "/b/file003.txt", ""}

	x.AddAction(&ta)

	for _, fn := range filenames {
		x.Add(core.Item{
			Path: fn,
			Info: defaultInfo,
		})
	}
	s.Len(x.Files, len(filenames))

	x.RunActions()

	s.Len(testResults, 0)
}
