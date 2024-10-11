package main

import "fmt"

type USB interface {
	Connect()
}

type USBConnecter struct {
}

func (c *USBConnecter) Connect(u USB) {
	u.Connect()
}

type USB3 struct {
}

func (u USB3) Connect() {
	fmt.Println("USB 3.0 подключен")
}

type TypeC struct {
}

func (tc TypeC) PlugIn() {
	fmt.Println("Type-C подключен")
}

type Adapter struct {
	TypeC
}

func (a Adapter) Connect() {
	fmt.Println("Адаптер присоединен к Type-C")
	a.TypeC.PlugIn()
}

func main() {
	c := USBConnecter{}
	u := USB3{}
	tc := TypeC{}
	a := Adapter{tc}
	c.Connect(u)
	c.Connect(a)
}
