package apps

import (
	"curse/practice31/handle"
	"log"
	"net/http"
)

const addr1 string = "localhost:8081"

func main() {
	handler := handle.Handle2
	http.HandleFunc("/", handler)
	log.Fatalln(http.ListenAndServe(addr1, nil))
}
