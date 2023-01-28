package main

import log "github.com/sirupsen/logrus"

func init() {
	log.SetFormatter(&log.JSONFormatter{})
}

func main() {
	a := 0
	i := 0

	for {
		if i >= 10 {
			break
		}

		a += i * i
		i++
	}

	log.WithFields(log.Fields{
		"i =": i,
		"a =": a,
	}).Info("miaw")
}
