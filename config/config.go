package config

import (
	"flag"
	logger "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

func init(){
	var fileConf string
	flag.StringVar(&fileConf, "config", `./config/app.config.yaml`, "Absolute path of configuration file")
	flag.Parse()
	ReadConfig(fileConf)
}

type AppConfiguration struct {
	ServerHost				string `yaml:"server.host"`
	ServerPort				string `yaml:"server.port"`

	DynamoCredentialsPath 	string `yaml:"dynamo.credentials.path"`
	Region 					string `yaml:"region"`
}

// Configuration ... The configuration of system
var AppConf AppConfiguration

// ReadConfig:
func ReadConfig(fileConfig string){
	loadConfiguration(fileConfig, &AppConf)
}


func loadConfiguration(fileConf string, conf *AppConfiguration)  {
	yamlFile, err := ioutil.ReadFile(fileConf)
	if err != nil {
		logger.Error("Can not load configuration file ", err)
	}

	err = yaml.Unmarshal(yamlFile, &conf)
	if err != nil {
		logger.Error("Can not load App Configuration ", err)
	}

	return
}