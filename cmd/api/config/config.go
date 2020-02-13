package config

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"path/filepath"
	"runtime"
)

// Load returns Configuration struct
func Load(env string) (*Configuration, error) {
	_, filePath, _, _ := runtime.Caller(0)
	configFile := filePath[:len(filePath)-9]
	bytes, err := ioutil.ReadFile(configFile +
		"env" + string(filepath.Separator) + "env." + env + ".yaml")
	if err != nil {
		return nil, fmt.Errorf("error reading config file, %s", err)
	}
	var cfg = new(Configuration)
	if err := yaml.Unmarshal(bytes, cfg); err != nil {
		return nil, fmt.Errorf("unable to decode into struct, %v", err)
	}
	return cfg, nil
}

// Configuration holds data necessery for configuring application
type Configuration struct {
	Server *Server
	DB     *Database
}

// Database holds data necessery for database configuration
type Database struct {
	Path     string
	Database string
}

// Server holds data necessery for server configuration
type Server struct {
	Port int
}
