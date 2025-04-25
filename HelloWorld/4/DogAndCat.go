package main

import "fmt"

type Speaker interface {
	Speak() string
}
type Dog struct {
	name string
}

func (d Dog) Speak() string {
	return fmt.Sprintf("%s says :woof", d.name)
}

type Cat struct {
	name string
}

func (c Cat) Speak() string {
	return fmt.Sprintf("%s says :meow", c.name)
}
func main() {
	dog := Dog{"a"}
	cat := Cat{"b"}
	var speaker Speaker
	speaker = &dog
	fmt.Println(speaker.Speak())
	speaker = &cat
	fmt.Println(speaker.Speak())

}
