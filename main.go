package main

import (
	"fmt"
	"framework/app/model"
	"framework/config"
	"framework/database/postgresql"
	"log"
)

var (
	server     = config.Server{}
	serverConf = server.GetServerConf("config.yml")
)

type EnvConn struct {
	db model.Datastore
}

func main() {
	conn := &EnvConn{postgresql.Conn()}
	if serverConf {
		log.Println("The configuration vars server was loaded!")
	}
	fmt.Println(conn)
	initServer()

}
