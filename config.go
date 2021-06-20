package main

import (
	"io/ioutil"
	"log"
	"os"
	"gopkg.in/yaml.v2"
)

type Config struct {
	Address string `yaml:"address"`
	Port    string    `yaml:"port"`
}

type Mimeconf struct {
	Mime []string `yaml:"types,flow"`
}


func getConf() (Config){
	file, err := os.Open("/etc/warp/config.yml")
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err = file.Close(); err !=nil {
			log.Fatal(err)
		}
	}()
	b, err := ioutil.ReadAll(file)
	var conf Config
	yaml.Unmarshal(b, &conf)
	return conf
}

func getMimeconf() (Mimeconf) {
	file, err := os.Open("/etc/warp/mime.yml")
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err = file.Close(); err !=nil {
			log.Fatal(err)
		}
	}()
	b, err := ioutil.ReadAll(file)
	var conf Mimeconf
	yaml.Unmarshal(b, &conf)
	return conf
}
