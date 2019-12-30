package readfiles

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

/**
 * Read file yml and fill a interface
 * @paramas path string, s interface *
 * @return map[interface{}]
 */
func YamlFileToStruct(path string) (map[string]map[string]interface{}, error) {
	s := make(map[string]map[string]interface{})
	yamlFile, err := ioutil.ReadFile(path)

	if err != nil {
		log.Fatalf("ReaderFile: %v", err)
		return nil, err
	}

	err = yaml.Unmarshal([]byte(yamlFile), &s)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}

	return s, nil
}
