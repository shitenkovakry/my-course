package proxy

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
)

const proxyAddr string = "localhost:9090"

var (
	counter               = 0
	theFirstInstanceHost  = "http://localhost:8080"
	theSecondInstanceHost = "http://localhost:8081"
)

func main() {

}

func Proxy(w http.ResponseWriter, r *http.Request) {
	textBytes, err := io.ReadAll(r.Body)
	if err != nil {
		log.Fatalln(err)
	}

	text := string(textBytes)

	if counter == 0 {
		resp, err := http.Post(theFirstInstanceHost, "text/plain", bytes.NewBuffer([]byte(text)))
		if err != nil {
			log.Fatalln(err)
		}

		counter = 1

		textBytes, err = io.ReadAll(resp.Body)
		if err != nil {
			log.Fatalln(err)
		}
		defer resp.Body.Close()

		fmt.Println(string(textBytes))

		return
	}

	resp, err := http.Post(theSecondInstanceHost, "text/plain", bytes.NewBuffer([]byte(text)))
	if err != nil {
		log.Fatalln(err)
	}

	textBytes, err = io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()

	fmt.Println(string(textBytes))

	counter = 0
}
