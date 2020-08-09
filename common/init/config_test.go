package init

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"testing"
)

func TestConfig(t *testing.T) {
	var yml Yaml
	file, err := ioutil.ReadFile("../conf/common.yml")
	if err != nil {
		panic(err)
	}
	yaml.Unmarshal(file, &yml)
	fmt.Println(yml)
}
