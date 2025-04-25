package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func main() {
	p1 := newPerson("a", 23)
	fmt.Println(p1)
	fmt.Println(p1.getName())
	fmt.Println(p1.getAge())
	p1.setName("w")
	p1.setAge(22)
	fmt.Println(p1.getName())
	fmt.Println(p1.getAge())

}

func newPerson(name string, age int) *Person {
	return &Person{name, age}
}
func (p *Person) setName(name string) {
	p.Name = name
}
func (p *Person) setAge(age int) {
	p.Age = age
}
func (p *Person) getName() string {
	return p.Name
}
func (p *Person) getAge() int {
	return p.Age
}
