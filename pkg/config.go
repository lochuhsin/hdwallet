package pkg

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

type SymbolConfig struct {
	Network     string   `yaml:"network"`
	PrivateKeys []string `yaml:"privatekeys"`
}

type Config struct {
	Mnemonic string                  `yaml:"mnemonic"`
	Symbols  map[string]SymbolConfig `yaml:"symbols"`
}

func NewConfig() Config {
	return Config{}
}

func ValidateConfig(config Config) error {
	for s := range config.Symbols {
		_, err := CoinSelector(s)
		if err != nil {
			return err
		}
	}
	return nil
}

func WriteConfig(config Config) error {
	d, err := yaml.Marshal(&config)
	if err != nil {
		return err
	}
	return WriteFile(d)
}

func ReadConfig() (Config, error) {
	b, err := os.ReadFile(GetConfigPath())
	if err != nil {
		return *new(Config), err
	}
	c := Config{}
	err = yaml.Unmarshal(b, &c)
	if err != nil {
		return *new(Config), err
	}
	return c, nil
}

func WriteFile(bytes []byte) error {
	dir := filepath.Dir(GetConfigPath())
	fmt.Printf("Writing config file to %s \n", dir)
	err := os.MkdirAll(dir, 0755)
	if err != nil {
		log.Fatalf("failed to create directories: %s", err)
		return err
	}
	return os.WriteFile(GetConfigPath(), bytes, FILE_PERMISSIONS)
}

func GetConfigPath() string {
	dir, err := os.UserHomeDir()
	if err != nil {
		panic("Unable to get user home directory")
	}
	return dir + "/.config/wallet/wallet.yaml"
}
