package goft

import (
	"io/ioutil"
	"os"
)

func LoadConfigFile() []byte {
	dir, _ := os.Getwd()
	file := dir + "/application.yaml"
	b, err := ioutil.ReadFile(file)
	Error(err)
	return b
}
