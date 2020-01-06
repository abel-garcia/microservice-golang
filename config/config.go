package config

import (
	"framework/tools/readfiles"
	"os"
	"fmt"
	"log"
)

/**
 * Parameters to run server
 */
type Server struct {
	WriteTimeout int64  `yaml:"write_timeout"`
	ReadTimeout  int64  `yaml:"read_timeout"`
	IdleTimeout  int64  `yaml:"idle_timeout"`
	Port         string `yaml:"port"`
	Addr         string `yaml:"addr"`
}

/**
 * Method get data server and create enviroment server vars
 * @params main\main path
 * @return bool
 */
func (s *Server) GetServerConf(path string) bool {
	if server, err := readfiles.YamlFileToStruct(path); err != nil {
		return false
	} else if ok := setEnvFromStruct(server); ok {
		log.Println("The configuration vars server was loaded!")
		return true
	}

	return false
}

/**
 * Recive interface for generate enviroment vars
 * @params interface *
 * @return bool
 */
func setEnvFromStruct(data map[string]map[string]interface{}) (success bool) {

	for key, items := range data {
		for _key, item := range items {
			if key == "database" {
				os.Setenv(fmt.Sprintf("%s%s%s",items["dialect"],"_",_key), fmt.Sprintf("%v", item))
			} else {
				os.Setenv(_key, fmt.Sprintf("%v", item))
			}
		}
	}

	return true
}
