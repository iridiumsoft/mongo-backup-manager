package config

import (
	"encoding/json"
	"io/ioutil"
)

// Main include info about API settings.
type Main struct {
	Port                       string        `json:"port"`
	Debug                      Debug         `json:"debug"`
	DB                         DB            `json:"database"`
	AWS                        AWSConfig     `json:"aws"`
	ExcludeCollectionsInBackup []interface{} `json:"exclude_collections_in_backup"`
}

type AWSConfig struct {
	Region     string `json:"region"`
	AccessKey  string `json:"access_key"`
	SecreteKey string `json:"secrete_key"`
	Bucket     string `json:"bucket"`
}

// DB include info about database connection.
type DB struct {
	Name string `json:"name"`
	Host string `json:"host"`
}

// Debug include all debugging flags
type Debug struct {
}

// FromFile read config from path and create configuration struct.
func FromFile(path string) Main {
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		panic("error reading config " + path + ": " + err.Error())
	}
	var conf Main
	if err := json.Unmarshal(bytes, &conf); err != nil {
		panic("error parsing config " + path + ": " + err.Error())
	}
	return conf
}
