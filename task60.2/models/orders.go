package models

type Order struct {
	ID        int
	Address   string
	Telephone string
	Status    string
}

type Orders []*Order
