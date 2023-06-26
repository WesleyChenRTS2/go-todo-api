//go:build tools

package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"

	"gopkg.in/yaml.v2"
)

type DBConfig struct {
	Open string `yaml:"open"`
}

func main() {
	configPath := "db/dbconf.yml"
	yamlFile, err := ioutil.ReadFile(configPath)
	if err != nil {
		log.Fatalf("Failed to read the config file: %v", err)
	}

	var config DBConfig
	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		log.Fatalf("Failed to unmarshal the config file: %v", err)
	}

	fmt.Print(strings.TrimSpace(config.Open))
}
