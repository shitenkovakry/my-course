package handle

import (
	"io/ioutil"
	"log"
	"net/http"
)

func Handle1(w http.ResponseWriter, r *http.Request) {
	bodyBytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatalln(err)
	}
	defer r.Body.Close()

	text := string(bodyBytes)
	response := "1 instance" + text + "\n"

	if _, err := w.Write([]byte(response)); err != nil {
		log.Fatalln(err)
	}
}

func Handle2(w http.ResponseWriter, r *http.Request) {
	bodyBytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatalln(err)
	}
	defer r.Body.Close()

	text := string(bodyBytes)
	response := "2 instance" + text + "\n"

	if _, err := w.Write([]byte(response)); err != nil {
		log.Fatalln(err)
	}
}
