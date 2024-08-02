package config

import (
	"path/filepath"

	"github.com/spf13/pflag"
)

// GRPCAuthorizerConfig holds the configuration settings for the gRPC authorizer.
type GRPCAuthorizerConfig struct {
	// EnableMock indicates whether the gRPC authorizer mock should be enabled.
	EnableMock bool `json:"enable_grpc_authorizer_mock"`

	// Type specifies the type of gRPC authorizer, e.g., "kube".
	Type string `json:"type"`

	// AuthorizerConfig specifies the path to the authorizer configuration file.
	AuthorizerConfig string `json:"authorizer_config"`
}

// NewGRPCAuthorizerConfig creates and returns a new instance of GRPCAuthorizerConfig with default values.
func NewGRPCAuthorizerConfig() *GRPCAuthorizerConfig {
	return &GRPCAuthorizerConfig{
		EnableMock:       false,
		Type:             "kube",
		AuthorizerConfig: filepath.Join(GetProjectRootDir(), "secrets/kube.config"),
	}
}

// AddFlags adds the gRPC authorizer configuration flags to the given FlagSet.
func (a *GRPCAuthorizerConfig) AddFlags(fs *pflag.FlagSet) {
	fs.BoolVar(&a.EnableMock, "enable_grpc_authorizer_mock", a.EnableMock, "Enable or disable gRPC authorizer mock.")
	fs.StringVar(&a.Type, "grpc_authorizer_type", a.Type, "Specify the gRPC authorizer type (e.g., kube).")
	fs.StringVar(&a.AuthorizerConfig, "authorizer_config", a.AuthorizerConfig, "Path to the gRPC authorizer configuration file.")
}
