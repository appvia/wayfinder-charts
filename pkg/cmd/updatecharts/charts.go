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

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/Masterminds/semver/v3"
	"gopkg.in/yaml.v2"

	"github.com/appvia/wayfinder-charts/pkg/rconfig"
	"github.com/appvia/wayfinder-charts/pkg/utils/compression"
	httputils "github.com/appvia/wayfinder-charts/pkg/utils/http"
)

// Flow:
// 1. check versions with latest in the index.yaml
// 2. If we have latest version than one in the versions file
// 2.1. Download and extract new version into ./charts
// 2.2. Update versions file with new version
func UpdateCharts(ctx context.Context, only string) error {

	var changed bool
	src := rconfig.GetSources()
	for name, chart := range src {
		if only != "" && name != only { // skip others if running for one chart
			continue
		}
		if !chart.External || chart.Pin { // skip non external and pinned charts
			continue
		}
		fmt.Println("Checking chart:", name)
		chartIndexUrl := chart.Source
		if chartIndexUrl == "" {
			return fmt.Errorf("no chart source available for %s - add a source to the release config file", name)
		}

		index, err := getChartIndex(ctx, chartIndexUrl)
		if err != nil {
			return fmt.Errorf("error getting chart index: %s", err.Error())
		}

		// check if index.yaml/chart repository contains what we want
		exist := checkIfExist(ctx, name, *index)
		if !exist {
			return fmt.Errorf("chart %s not found in index", name)
		}

		namedChart := index.Entries[name]

		currentVersion, err := semver.NewVersion(chart.Version)
		if err != nil {
			return err
		}

		// check current version vs whats upstream
		latest, err := getLatest(ctx, namedChart)
		if err != nil {
			return err
		}

		if currentVersion.LessThan(latest) {
			fmt.Printf("Updating chart: %s from %s to %s\n", name, chart.Version, latest)
			// update chart files by downloading and extracting files
			err := updateLocalChart(ctx, name, chartIndexUrl, namedChart, latest)
			if err != nil {
				return err
			}
			// update chart version
			d := src[name]
			d.Version = latest.String()
			src[name] = d
			changed = true
		}
	}
	if changed {
		return rconfig.PersistSources(src)
	}

	return nil
}

func getChartIndex(ctx context.Context, url string) (*HelmIndex, error) {
	body, err := httputils.Get(ctx, url)
	if err != nil {
		return nil, err
	}
	defer body.Body.Close()

	charts := HelmIndex{}
	err = yaml.NewDecoder(body.Body).Decode(&charts)
	if err != nil {
		return nil, err
	}

	return &charts, nil
}

func checkIfExist(_ context.Context, name string, index HelmIndex) bool {
	for iname, idx := range index.Entries {
		if iname == name && len(idx) > 0 {
			return true
		}
	}
	return false
}

func getLatest(_ context.Context, entries []Entry) (*semver.Version, error) {
	latestV, err := semver.NewVersion("0.0.0")
	if err != nil {
		return nil, err
	}
	for _, entry := range entries {
		v, err := semver.NewVersion(entry.Version)
		if err != nil {
			fmt.Printf("Error parsing version %s: %s\n", entry.Version, err.Error())
			continue
		}
		if latestV.LessThan(v) && v.Prerelease() == "" {
			latestV = v
		}
	}

	return latestV, nil
}

func updateLocalChart(ctx context.Context, name, url string, entries []Entry, v *semver.Version) error {

	for _, entry := range entries {
		ev, err := semver.NewVersion(entry.Version)
		if err != nil {
			fmt.Printf("Error parsing version %s: %s\n", entry.Version, err.Error())
			continue
		}
		if v.Equal(ev) {
			// download and extract
			if len(entry.Urls) == 0 {
				return fmt.Errorf("no urls found for %s", entry.Version)
			}
			chartSourceUrl := entry.Urls[0]

			valid := isUrl(chartSourceUrl)
			if !valid {
				// some helm charts have urls that is local. We need to prefix it with original source url
				base := strings.ReplaceAll(url, "/index.yaml", "")
				chartSourceUrl = fmt.Sprintf("%s/%s", base, chartSourceUrl)
			}

			resp, err := httputils.Get(ctx, chartSourceUrl)
			if err != nil {
				return err
			}
			defer resp.Body.Close()

			if resp.StatusCode != http.StatusOK {
				return fmt.Errorf("bad status: %s", resp.Status)
			}

			chartTemp, err := os.CreateTemp("/tmp", "chart-")
			if err != nil {
				return err
			}
			defer os.Remove(chartTemp.Name())

			_, err = io.Copy(chartTemp, resp.Body)
			if err != nil {
				return err
			}
			err = resp.Body.Close()
			if err != nil {
				return err
			}
			err = chartTemp.Close()
			if err != nil {
				return err
			}

			f, err := os.Open(chartTemp.Name())
			if err != nil {
				return err
			}
			defer f.Close()

			output := fmt.Sprintf("./charts/%s/%s", name, ev.String())
			err = compression.ExtractTarGz(f, name, output)
			if err != nil {
				return err
			}

		}
	}

	return nil
}

func isUrl(str string) bool {
	u, err := url.Parse(str)
	return err == nil && u.Scheme != "" && u.Host != ""
}
