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
	"context"

	"github.com/spf13/cobra"
)

func main() {
	ctx := context.Background()
	// errors are being printed by CLI handelers
	err := RunCLI(ctx)
	if err != nil {
		panic(err)
	}

}

// RunCLI returns user CLI
func RunCLI(ctx context.Context) error {
	cmd := &cobra.Command{
		Short: "Wayfinder Chart Update tool",
		Use:   "updatecharts update-charts",
	}

	cmd.AddCommand(NewUpdateChartsCmd())

	return cmd.ExecuteContext(ctx)
}
