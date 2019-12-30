package controller

import (
	"net/http"
)

func CreateBoook(writer http.ResponseWriter, request *http.Request) {
	writer.Write([]byte("book saved!\n"))
}
