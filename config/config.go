package config

import (
	"fmt"
	"framework/tools/readfiles"
	"os"
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
	} else if ok := setEnvFromStruct(server); !ok {
		return !ok
	}

	return true
}

/**
 * Recive interface for generate enviroment vars
 * @params interface *
 * @return bool
 */
func setEnvFromStruct(data map[string]map[string]interface{}) (success bool) {

	for _, items := range data {
		for key, item := range items {
			os.Setenv(key, fmt.Sprintf("%v", item))
		}
	}

	return true
}
