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

package main

import (
	"github.com/spf13/cobra"

	"github.com/appvia/wayfinder-charts/pkg/cmd/updatecharts"
)

var only string

// NewUpdateChartsCmd returns the cobra command for "update-charts".
func NewUpdateChartsCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update-charts",
		Short: "Updates all charts",
		Long: `Updates all charts based on the sources configured in /pkg/rconfig/sources.yaml. 
Use the optional --only flag to update just one chart.

Example - update all:
$ go run ./cmd/updatecharts/ update-charts

Example - update one:
$ go run ./cmd/updatecharts/ update-charts --only terranetes-controller
`,
		RunE: func(cmd *cobra.Command, args []string) error {
			return updatecharts.UpdateCharts(cmd.Context(), only)
		},
	}

	cmd.Flags().StringVar(&only, "only", "", "Pull new version for just one chart, e.g. --only cert-manager")

	return cmd
}
