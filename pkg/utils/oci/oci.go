package oci

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"os"

	"helm.sh/helm/v3/pkg/cli"
	"helm.sh/helm/v3/pkg/registry"
)

func Get(ctx context.Context, chartUrl string) (io.Reader, error) {
	settings := cli.New()

	client, err := registry.NewClient(
		registry.ClientOptCredentialsFile(settings.RegistryConfig),
		registry.ClientOptWriter(os.Stdout),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create registry client: %w", err)
	}

	// Parse chart URL to get name and version
	ref := chartUrl[6:] // Remove "oci://" prefix

	// Pull the chart
	pullResult, err := client.Pull(ref, registry.PullOptWithChart(true), registry.PullOptIgnoreMissingProv(true))
	if err != nil {
		return nil, fmt.Errorf("failed to download chart: %w", err)
	}

	return bytes.NewReader(pullResult.Chart.Data), nil
}
