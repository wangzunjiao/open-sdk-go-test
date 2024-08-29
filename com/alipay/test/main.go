package main

import (
	"fmt"
	"testing"
)

type IChildHandle interface {
	AddAge()
	HandleAfterInit(handler IChildHandle)
}

type Base struct {
	name string
	age  int
}

func (b *Base) AddAge() {
	fmt.Println("base add Age .....")
}

func (b *Base) reName() {
	fmt.Println("rename ing...")
}

func (b *Base) HandleAfterInit(handler IChildHandle) {
	fmt.Printf("before save base obj name:%s,age:%d\r\n", b.name, b.age)
	b.name = b.name + "aaaaa"
	handler.AddAge()
	fmt.Printf("real save base obj name:%s,age:%d\r\n", b.name, b.age)
}

type Child1 struct {
	*Base
	price int
}

func (c *Child1) AddAge() {
	fmt.Println("child add age ...")
	c.age += 2
}

func TestDuoTai(t *testing.T) {
	b := &Base{
		name: "base xx",
		age:  100,
	}
	c := &Child1{
		Base:  b,
		price: 333,
	}
	b.HandleAfterInit(c)
	c.AddAge()
	c.reName()
	fmt.Println(c.age)
}
