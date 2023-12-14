package client

import (
	"fmt"
	"net/url"

	opservice "github.com/ethereum-optimism/optimism/op-service"
	"github.com/urfave/cli/v2"
)

const DARpcAddressFlagName = "da-server"

func CLIFlags(envPrefix string) []cli.Flag {
	return []cli.Flag{
		&cli.StringFlag{
			Name:    DARpcAddressFlagName,
			Usage:   "HTTP RPC address of a DAServer",
			EnvVars: opservice.PrefixEnvVar(envPrefix, "DA_SERVER"),
		},
	}
}

type Config struct {
	Enabled bool
	URL     string
}

func (c Config) Check() error {
	if c.Enabled {
		if c.URL == "" {
			return fmt.Errorf("DA server URL is required when AltDA is enabled")
		}
		if _, err := url.Parse(c.URL); err != nil {
			return fmt.Errorf("DA server URL is invalid: %w", err)
		}
	}
	return nil
}

type CLIConfig struct {
	DAServer string
}

func (c *CLIConfig) Config(enabled bool) Config {
	return Config{
		Enabled: enabled,
		URL:     c.DAServer,
	}
}

func ReadCLIConfig(c *cli.Context) CLIConfig {
	return CLIConfig{
		DAServer: c.String(DARpcAddressFlagName),
	}
}
