package main

import (
	"fmt"
	"strconv"
)

type Human struct {
	 name string
	age int
	phone string
}

type Student struct {
	Human
	school string
	loan float32
}

type Employee struct {
	Human
	company string
	money float32
}

func (h Human)sayHi() {
	fmt.Println("Hi i am")
}

func (s Student)sayHi() {
	fmt.Printf("hi, I am %s , a student\n", s.name)
}

func (h Human)sing(lyrics string) {
	fmt.Println("la la la la..." , lyrics)
}

func (e Employee)sayHi() {
	fmt.Printf("Hi , im an %s ,work at %s ,you can call on %s\n" , e.name, e.company, e.phone)
}

func (h Employee)Stringer() string {
	return "<"+h.name+"-"+strconv.Itoa(h.age) +"Years >"
}


type Men interface  {
	sayHi()
	sing(lyrics string)

}

func main() {
	mike := Student{ Human{"Mike", 25,"222-222-xxx1"} , "MIT", 0.0}
	paul := Student{ Human{"Paul", 26, "111-111-yyy2"} , "Harvard", 1000 }
	sam  := Employee{ Human{"Sam", 36, "444-222-xxx2" },"Golang Inc.", 1000}
	Tom  := Employee{ Human{"Tom", 37,"222-444-xxx4"}, "Things Ltd.", 5000}

	var i Men

	i = mike
	fmt.Println("This is Mike, a student")
	i.sayHi()
	i.sing("November rain")

	i = Tom

	fmt.Println("This is Tom,Employee")
	i.sayHi()
	i.sing("jje")

	x := make([]Men, 4)

	x[0], x[1], x[2],x[3] = mike, sam, Tom ,paul

	for _, value := range x {
		value.sayHi()

		fmt.Println(value)
	}
}

