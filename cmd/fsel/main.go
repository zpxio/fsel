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

package main

import (
	"github.com/apex/log"
	"github.com/zpxio/fsel/pkg/dir"
	"github.com/zpxio/fsel/pkg/session"
	"os"
)

func main() {
	log.Infof("Starting up.")

	log.Debugf("Creating Session")
	s := session.NewSession()

	cwd, _ := os.Getwd()
	log.Infof("Reading directory: %s", cwd)
	r := dir.CreateReader(cwd, s)
	err := r.Read()
	if err != nil {
		log.Errorf("Error adding files: %s", err)
	}

	s.RunActions()
}
