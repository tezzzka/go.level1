package main

import (
	"flag"
	"io/ioutil"

	"gopkg.in/yaml.v3"
)

type MyConfig struct {
	HOST string `yaml:"host"`
	PORT int32
	OS   string `yaml:"os"`
	ARCH string `yaml:"arch"`
}

var (
	host = flag.String("host", "localhost", "127.0.0.1")
	port = flag.String("port", "port", "3000")
	os   = flag.String("os", "unix", "unix")
	arch = flag.String("arch", "amd64", "amd64")
)

func GetConfigYaml() (*MyConfig, error) {
	buf, err := ioutil.ReadFile("config.yaml")
	if err != nil {
		return nil, err
	}

	data := &MyConfig{}
	err = yaml.Unmarshal(buf, data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func Config() *MyConfig {
	flag.Parse()
	//cli-code here. Реализуем в след.раз, ямл реализовали. Здесь сделаем так: если хоть один из параметров был передан в cli, то он будет перезаписан в структуре и всё равно будет передан в data

	if data, e := GetConfigYaml(); e == nil {
		return data
	} else {
		return nil
	}

}
