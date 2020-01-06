package main

import (
	"fmt"
	"framework/routes"
	"framework/tools/convertions"
	"os"

	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

/**
 *
 */
var (
	addr         = fmt.Sprintf("%s%s%s", os.Getenv("addr"), ":", os.Getenv("port"))
	writeTimeout = time.Duration(convertions.StringToInt64(os.Getenv("write_timeout")))
	readTimeout  = time.Duration(convertions.StringToInt64(os.Getenv("read_timeout")))
	idleTimeout  = time.Duration(convertions.StringToInt64(os.Getenv("idle_timeout")))
)

/**
 * Start server
 */
func initServer() {

	log.Printf("HTTP SERVICE LISTENING: %s", addr)

	if err := configServer(routes.AplicationV1Router()).ListenAndServe(); err != nil {
		log.Printf("%s", err)
		panic("Error: Server init")
	}

}

/**
 * Set configuration server data from enviroment vars
 * @param routes\AplicationV1Router router
 * @return http.Server
 */
func configServer(router *mux.Router) *http.Server {
	return &http.Server{
		Addr:         addr,
		WriteTimeout: time.Second * writeTimeout,
		ReadTimeout:  time.Second * readTimeout,
		IdleTimeout:  time.Second * idleTimeout,
		Handler:      router,
	}
}
