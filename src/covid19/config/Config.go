package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"path/filepath"
)

//MH12Config mh12 app config
type MH12Config struct {
	Database    DatabaseConfig `yaml:"database"`
	Log         LogConfig      `yaml:"log"`
	Application Application    `yaml:"application"`
}

//Application application config
type Application struct {
	Port           string      `yaml:"http-port"`
	Name           string      `yaml:"name"`
	Token          string      `yaml:"token"`
	TokenExpiryHrs int         `yaml:"tokenExpiryHrs"`
	User           UserAccount `yaml:"user"`
}

//UserAccount user account config
type UserAccount struct {
	Name     string `yaml:"name"`
	Phone    string `yaml:"phone"`
	Email    string `yaml:"email"`
	Password string `yaml:"password"`
}

//DatabaseConfig database config
type DatabaseConfig struct {
	DBName   string `yaml:"dbname"`
	Username string `yaml:"user"`
	Password string `yaml:"password"`
	Location string `yaml:"location"`
	Port     string `yaml:"port"`
}

//LogConfig log config
type LogConfig struct {
	Location   string `yaml:"location"`
	Level      string `yaml:"level"`
	MaxBackups int    `yaml:"maxbackups"`
	MaxAge     int    `yaml:"maxage"`
}

//Config config
type Config struct {
	MH12Config MH12Config `yaml:"mh12config"`
}

var configFile Config

//GetConfig get config
func GetConfig() Config {
	return configFile
}

//ReadYamlConfigFile Initial Function
func ReadYamlConfigFile() error {
	var configMH12 = MH12Config{}
	// mh12 Config
	yamlFile, err := ioutil.ReadFile(getMH12ConfigPath())
	if err != nil {
		return err
	}
	err = yaml.Unmarshal(yamlFile, &configMH12)
	if err != nil {
		return err
	}
	configFile = Config{MH12Config: configMH12}
	return nil
}

func getMH12ConfigPath() string {
	return GetConfigPath() + "/yaml/configMH12.yaml"
	//return "/Users/vinodborole/Documents/personal/interview/go/final/mh12covid19/bin/yaml/configMH12.yaml"
}

//GetConfigPath get config path
func GetConfigPath() string {
	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}
	exPath := filepath.Dir(ex)
	return exPath
}
