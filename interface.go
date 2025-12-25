package main

import "fmt"

type Speaker interface{
	Speak() string
}

type Human struct{
	Name string
}

func (h Human) Speak() string{
	return "Hello, I'm" + h.Name
}

func interfaceDemo(){
	h := Human{Name: "Vansh"}
	fmt.Println(h.Speak())
}