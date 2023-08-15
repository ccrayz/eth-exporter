package config

import (
	"fmt"
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

type YamlConfig struct {
	EthAccounts []EthAccount `yaml:"eth_accounts"`
}

type EthAccount struct {
	Purpose string `yaml:"purpose"`
	Address string `yaml:"address"`
}

func LoadConfig(configFileName string) []EthAccount {
	yamlFile, err := os.ReadFile(configFileName)
	if err != nil {
		log.Fatalf("Error reading YAML file: %v\n", err)
	}

	var config YamlConfig
	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		log.Fatalf("Error parsing YAML file: %v\n", err)

	}

	for _, acc := range config.EthAccounts {
		fmt.Printf("Purpose: %s, Address: %s\n", acc.Purpose, acc.Address)
	}

	return config.EthAccounts
}
