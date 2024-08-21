package service

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"wallet/pkg"

	"gopkg.in/yaml.v2"
)

type SymbolConfig struct {
	Network     string   `yaml:"network"`
	PrivateKeys []string `yaml:"privatekeys"`
	Index       int      `yaml:"index"`
}

func NewSymbolConfig() SymbolConfig {
	return SymbolConfig{
		Network:     "",
		Index:       0,
		PrivateKeys: []string{},
	}
}

type Config struct {
	Mnemonic string                          `yaml:"mnemonic"`
	Password string                          `yaml:"password"`
	Symbols  map[pkg.CoinSymbol]SymbolConfig `yaml:"symbols"`
}

func NewConfig() Config {
	return Config{
		Mnemonic: "",
		Password: "",
		Symbols:  make(map[pkg.CoinSymbol]SymbolConfig),
	}
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

func GetConfig() Config {
	config, err := ReadConfig()
	if err != nil {
		fmt.Printf("Unable to read config: %s \n", err)
		fmt.Println("Setting up default config")
		config = NewConfig()
		err := WriteConfig(config)
		if err != nil {
			fmt.Println(err)
		}
	}
	return config
}
