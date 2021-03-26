/*
 * Return wanted metrics in a structure that will be processed by other parts
 */

package client

import (
    "io/ioutil"
    "gopkg.in/yaml.v2"
)

type Config struct {
	Targets []struct {
		Host        string `yaml:"host"`
		Firstvalue  int    `yaml:"firstvalue"`
		Secondvalue int    `yaml:"secondvalue"`
		Thirdvalue  int    `yaml:"thirdvalue"`
		Fourthvalue int    `yaml:"fourthvalue"`
	} `yaml:"targets"`
}

// this will eventually be a method call, for now we test
const ConfigFile = "/home/osboxes/development/static-exporter/test/config/test.yml"

func GetThresholds() (*Config, error){
    yamlFile, err := ioutil.ReadFile(ConfigFile)

    if err != nil {
        return nil, err
    }

    var config Config

    err = yaml.Unmarshal(yamlFile, &config)
    if err != nil {
        return nil, err
    }

    return &config, nil

}
