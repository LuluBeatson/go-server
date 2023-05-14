package utils

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"

	"gopkg.in/yaml.v2"
)

// ParseBody parses the body of a request into a struct
func ParseBody(r *http.Request, x interface{}) {
	if body, err := ioutil.ReadAll(r.Body); err == nil {
		if err := json.Unmarshal([]byte(body), x); err != nil {
			return
		}
	}
}

type Config struct {
	MySQL struct {
		User     string `yaml:"usr"`
		Password string `yaml:"pwd"`
		Host     string `yaml:"host"`
		Port     string `yaml:"port"`
	} `yaml:"mySql"`
}

func GetConfig() (*Config, error) {
	file, err := os.Open("config.yaml")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	data, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}

	var config Config
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return nil, err
	}
	return &config, nil
}
