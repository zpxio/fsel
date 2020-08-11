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

package actions

import (
	"fmt"
	"github.com/apex/log"
	"github.com/zpxio/fsel/pkg/core"
)

type Print struct {
}

func (a *Print) Name() string {
	return "Print"
}

func (a *Print) BatchSize() int {
	return 2048
}

func (a *Print) Run(items []*core.Item) error {

	log.Infof("Printing batch of %d", len(items))
	for _, i := range items {
		fmt.Printf("%s\n", i.Path)
	}

	return nil
}
