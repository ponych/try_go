package main

import "fmt"

type person struct{
	name  string
	age  int
}

type student struct{
	person // 继承了 person 的所有属性
	teacher person
	subject string
}

func older(p1, p2 person) (p person , diff int) {
	if p1.age > p2.age {
		p = p1
		diff = p1.age - p2.age
	} else {
		p = p2
		diff = p2.age - p1.age
	}

	return
}

func main() {
	var p person
	p2 := person{"John", 55}
	p.name = "Eric"
	p.age = 23

	fmt.Println(p)
	fmt.Println(p2)
	var s student
	s.name = "Cate"
	s.age = 16
	s.subject = "math"
	s.teacher = p
	fmt.Println(s)
	fmt.Printf("my teacher is %s, %d years old\n", s.teacher.name ,s.teacher.age)
	p.age += 5
	// did not changed
	fmt.Printf("after 5 years,my teacher is %d years old\n" ,s.teacher.age)
	s2 := student{ person{"Eric.Chen", 18} , p2, "Physics"}
//
	fmt.Println(s2)

	fmt.Printf("student 2 is :%s at age :%d studing: %s\n" , s2.name, s2.age, s2.subject)
	s2.age += 13
	older, diff := older(s2.person, p)

	fmt.Printf("who is older: %s, %d years older", older.name, diff)
}