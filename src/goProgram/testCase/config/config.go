package config

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	yaml "gopkg.in/yaml.v2"
)

var BaseConfig = new(Yaml)

func init() {
	ReadYamlConfig()
	ReadCaseXML()
}


func ReadYamlConfig(){
	//conf := new(Yaml)
	dir,_ := os.Getwd()
	str := dir+`\config\config.yaml`
	fmt.Println("str:",str)
	yamlFile, err := ioutil.ReadFile(str)
	log.Println("yamlFile:\n",string(yamlFile))
	if err != nil {
		log.Printf("%v:",err)
	}
	err = yaml.Unmarshal(yamlFile,BaseConfig)
	fmt.Println("BaseConfig:",BaseConfig.Mysql)
}

type Yaml struct {
	Mysql struct {
		User 	string `yaml:"user"`
		Host	string `yaml:"host"`
		Password string `yaml:"password"`
		Port	string  `yaml:"port"`
		Name	string  `yaml:"name"`
	}
	Cache struct {
		Enable  bool `yaml:"enable"`
		List	[]string `yaml:"list,flow"`
	}
}

type Mysql struct {
	User 	string  	`yaml:"user"`
	Hose    string 		`yaml:"host"`
	Password  string 	`yaml:"password"`
	Port 	string 		`yaml:"port"`
	Name 	string  	`yaml:"name"`
}

type Cache struct {
	Enable 	bool 	`yaml:"enable"`
	List 	[]string `yaml:"list,flow"`
}
