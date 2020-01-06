package main

import (
	"os"
	"framework/config"
	"framework/database/psql"
	"log"
)

var (
	server     = config.Server{}
	serverConf = server.GetServerConf("config.yml")
)

func main() {
	if !serverConf {
		log.Println("The configuration vars server didnt loaded!")
	} else {
		dataBaseInit()
		initServer()
	}
}

/*
* range all dialect and create news vars
*/
func dataBaseInit() {
	var database = os.Getenv("postgresql_dialect")
	switch database {
	case "postgresql":

		if err := psql.Connection.Set(psql.Conn()); err == nil {
			log.Println("The database postgresql was loaded!")
		} 
	}
}
