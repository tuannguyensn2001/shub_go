package errpkg

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)


var General *general
var Auth *auth

func LoadError() {
	root := rootErr{}

	file, err := ioutil.ReadFile("error.yaml")

	if err != nil {
		log.Fatalln("error load error", err)
	}

	err = yaml.Unmarshal(file, &root)

	if err != nil {
		log.Fatalln("error unmarshal file", err)
	}

	//General = root.General
	
    General = root.General 
    Auth = root.Auth 
}
