package helpers

import (
	"fmt"

	"github.com/pelletier/go-toml"
)

/* LoadEnvFile - Loading TOML file */
func LoadEnvFile() (*toml.Tree, error) {
	config, err := toml.LoadFile("config.toml")
	if err != nil {
		return nil, fmt.Errorf("[ERROR] loading toml file: %+v\n", err)
	}

	return config, nil
}
