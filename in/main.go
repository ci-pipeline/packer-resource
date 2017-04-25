package main

import (
	"io/ioutil"
	"os"

	"github.com/ci-pipeline/concourse-ci-resource/utils"
)

func main() {
	destination := os.Args[1]
	input := utils.GetInput()
	utils.Logln(input.Version)

	for k, v := range input.Version.(map[interface{}]interface{}) {
		err := ioutil.WriteFile(destination+"/"+k.(string), []byte(v.(string)), 0644)
		if err != nil {
			panic(err)
		}
	}

}
