package apis

import (
	"net/http"
)

func GetHealth(writer http.ResponseWriter, request *http.Request) {
	writer.Write([]byte("UP"))
}
