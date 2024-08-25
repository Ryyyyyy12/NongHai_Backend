package config

import (
	"backend/internal/util/text"
	"os"

	"github.com/davecgh/go-spew/spew"
	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"
)

var Conf = new(config)

func init() {

	//load config
	yml, err := os.ReadFile("./config.yaml")
	if err != nil {
		logrus.Fatalf("Error reading config file: %v", err)
	}
	if err := yaml.Unmarshal(yml, Conf); err != nil {
		logrus.Fatalf("Error parsing config file: %v", err)
	}

	//validate config
	if err := text.Validator.Struct(Conf); err != nil {
		logrus.Fatalf("Error validating config file: %v", err)
	}

	//apply log level config
	logrus.SetLevel(logrus.Level(Conf.LogLevel))
	spew.Config = spew.ConfigState{Indent: "  "}

}
