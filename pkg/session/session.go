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
	"github.com/apex/log"
	"github.com/zpxio/fsel/pkg/actions"
	"github.com/zpxio/fsel/pkg/core"
	"sort"
)

const MaxBuffer = 2048

var DefaultActions []Action

func init() {
	DefaultActions = make([]Action, 0)

	DefaultActions = append(DefaultActions, &actions.Print{})
}

type Session struct {
	Files   map[string]core.Item
	Actions []Action
}

func NewSession() *Session {
	s := Session{
		Files:   make(map[string]core.Item),
		Actions: make([]Action, 0),
	}

	return &s
}

func (s *Session) Add(item core.Item) {
	s.Files[item.Path] = item
}

func (s *Session) AddAction(action Action) {
	s.Actions = append(s.Actions, action)
}

func (s *Session) effectiveActions() []Action {
	if len(s.Actions) < 1 {
		return DefaultActions
	} else {
		return s.Actions
	}
}

func (s *Session) RunActions() {
	bufferSize := MaxBuffer

	for _, a := range s.effectiveActions() {
		if a.BatchSize() < bufferSize {
			bufferSize = a.BatchSize()
		}
	}

	// Get a sorted list of files
	items := make([]*core.Item, len(s.Files))
	i := 0
	for path := range s.Files {
		item := s.Files[path]
		items[i] = &item
		i++
	}

	// Sort the items
	sort.Slice(items, func(a int, b int) bool {
		return items[a].Path < items[b].Path
	})

	// Run all the actions
	for start := 0; start < len(s.Files); start += bufferSize {
		for _, a := range s.effectiveActions() {
			bufferEnd := start + bufferSize
			if bufferEnd > len(items) {
				bufferEnd = len(items)
			}
			err := a.Run(items[start:bufferEnd])
			if err != nil {
				log.Errorf("Action [%s] failed: %s", a.Name(), err)
				break
			}
		}
	}
}
