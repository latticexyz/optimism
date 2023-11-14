package client

import (
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

type CLIConfig struct {
	DAServer string
}

func (c *CLIConfig) Config() Config {
	return Config{
		Enabled: c.DAServer != "",
		URL:     c.DAServer,
	}
}

func ReadCLIConfig(c *cli.Context) CLIConfig {
	return CLIConfig{
		DAServer: c.String(DARpcAddressFlagName),
	}
}
