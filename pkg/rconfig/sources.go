/**
 * Copyright 2021 Appvia Ltd <info@appvia.io>
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

package rconfig

import (
	_ "embed"
	"os"

	"gopkg.in/yaml.v2"
)

//go:embed sources.yaml
var sources []byte

// GetSources returns the chart source list
func GetSources() ChartList {
	var conf ChartList

	if err := yaml.Unmarshal(sources, &conf); err != nil {
		panic(err)
	}

	return conf
}

func PersistSources(src ChartList) error {
	data, err := yaml.Marshal(src)
	if err != nil {
		return err
	}

	return os.WriteFile("pkg/rconfig/sources.yaml", data, 0644)
}
