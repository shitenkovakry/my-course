package apps

import (
	"curse/practice31/handle"
	"log"
	"net/http"
)

const addr string = "localhost:8080"

func main() {
	handler := handle.Handle1
	http.HandleFunc("/", handler)
	log.Fatalln(http.ListenAndServe(addr, nil))
}
