package rconfig

type ChartList map[string]ChartConfig

// Populated indicates whether the chart list has any charts in it
func (c ChartList) Populated() bool {
	return len(c) > 0
}

// type ProviderVersionConfig struct {
// 	// Charts defines overrides for global chart values. It is used to deviate from global config for individual releases
// 	Charts ChartList `json:"charts,omitempty" yaml:"charts,omitempty"`
// 	// Images defines other images used in the release
// 	Images Images `json:"images,omitempty" yaml:"images,omitempty"`
// }

// ChartConfig is an individual chart configuration
type ChartConfig struct {
	Version  string `json:"version,omitempty" yaml:"version,omitempty"`
	Source   string `json:"source,omitempty" yaml:"source,omitempty"`
	External bool   `json:"external,omitempty" yaml:"external,omitempty"`
	Values   []byte `json:"values,omitempty" yaml:"values,omitempty"`
	// Pin indicates that we should not automatically update this dependency, manual updates only.
	Pin bool `json:"pin,omitempty" yaml:"pin,omitempty"`
}
