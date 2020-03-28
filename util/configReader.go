package util

import (
	"fmt"
	"os"

	"github.com/kelseyhightower/envconfig"
	"gopkg.in/yaml.v2"
)

type Config struct {
	Database struct {
		Name     string `yaml:"name", envconfig:"DB_NAME"`
		Username string `yaml:"user", envconfig:"DB_USERNAME"`
		Password string `yaml:"password", envconfig:"DB_PASSWORD"`
	} `yaml:"database"`
}

func processError(err error) {
	fmt.Println(err)
	os.Exit(2)
}

func (c *Config) readFile() {

	// pwd, _ := os.Getwd()
	//fmt.Println(pwd)

	var f *os.File
	var err error
	// if strings.Contains(pwd, "MapApp") || strings.Contains(pwd, "applications") {
	// 	f, err = os.Open("config.yml")
	// } else {
	// 	f, err = os.Open("MapApp/config.yml")
	// }
	f, err = os.Open("config.yml")

	if err != nil {
		processError(err)
	}
	defer f.Close()

	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(c)
	if err != nil {
		processError(err)
	}
}

func (c *Config) readEnv() {
	err := envconfig.Process("", c)
	if err != nil {
		processError(err)
	}
}

func (c *Config) GetConfigData() {
	c.readFile()
	c.readEnv()
}
