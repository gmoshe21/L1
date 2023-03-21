package main

import "fmt"

type Human struct {
	Name     string
	LastName string
}

func (h *Human) Speak() {
	fmt.Println(h.Name, h.LastName)
}

type Action struct {
	Human
	SitDown bool
}

func main() {
	h := Action{
		Human{
			Name:   "Grandfather",
			LastName: "AY",
		},
		false,
	}

	h.Speak()
}