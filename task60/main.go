package main

import "log"

type Order struct {
	ID        int
	Address   string
	Telephone string
	Status    string
}

type Orders []*Order

type Dispatcher struct {
	orders Orders
	log    log.Logger
}
