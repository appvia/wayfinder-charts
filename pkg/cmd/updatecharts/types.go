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

package updatecharts

import "time"

type HelmIndex struct {
	APIVersion string             `yaml:"apiVersion"`
	Generated  time.Time          `yaml:"generated"`
	Entries    map[string][]Entry `yaml:"entries"`
}
type Maintainers struct {
	Email string `yaml:"email"`
	Name  string `yaml:"name"`
}
type Entry struct {
	APIVersion  string        `yaml:"apiVersion"`
	AppVersion  string        `yaml:"appVersion"`
	Created     time.Time     `yaml:"created"`
	Description string        `yaml:"description"`
	Digest      string        `yaml:"digest"`
	Home        string        `yaml:"home"`
	Maintainers []Maintainers `yaml:"maintainers"`
	Name        string        `yaml:"name"`
	Sources     []string      `yaml:"sources"`
	Urls        []string      `yaml:"urls"`
	Version     string        `yaml:"version"`
}
