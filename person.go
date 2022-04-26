package main

import "fmt"

type person struct {
    name string
    age  int
}

func newPerson(name string) person {
	p := person{name: name}
    p.age = 42
    return p
}

func itsMyBirthday(pers person){
	pers.age = pers.age +1
	fmt.Println(pers.age)
}