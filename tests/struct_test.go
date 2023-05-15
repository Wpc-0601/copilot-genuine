package tests

import (
	"math/rand"
	"strconv"
	"testing"
)

type validate func(name string) string

type OrderLine struct {
	id   string
	line int
	name string
}
type Address struct {
	no string
}
type Order struct {
	Id        string
	Address   Address
	OrderLine OrderLine
	Valid     validate
}

func (order *Order) changeName(name string) string {
	if name != "" {
		order.OrderLine.name = name
		return "success"
	}
	return "Failed"
}

func TestOrder(test *testing.T) {
	order := New()
	order.Valid = func(name string) string {
		result := order.changeName(order.OrderLine.name)
		if result == "success" && name != "" {
			return "valid name"
		}
		return "invalid name, alternatively the order lien id validate failed"
	}
	order.Valid("sj")
}

func New() *Order {
	address := &Address{no: "14101090110"}
	line := &OrderLine{
		id:   strconv.Itoa(rand.Intn(100)),
		line: 17,
		name: "The bed seems to occupy the most of room",
	}
	order := &Order{
		Id:        strconv.Itoa(rand.Intn(2000)),
		Address:   *address,
		OrderLine: *line,
	}
	return order
}
