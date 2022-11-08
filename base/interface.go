package main

import (
	"fmt"
	"reflect"
)

func main() {
	var fullName FullNameInterface = &user{"alisher", "kalimoldayev", 19} // создаю переменную с типом интерфейс, которая является структурой #1
	fmt.Println(fullName.getFullName())                                   // #1

	var dog animalInterface = &dog{15} // #2
	dog.walk()
	dog.run()

	m := map[string]interface{}{} // пустой интерфейс. используем когда не знаем какого типа элемент для передачи в функцию/map
	m["one"] = 15
	m["two"] = "lorem"
	m["three"] = 12.5
	m["four"] = []int{5, 7}

	for i := range m {
		switch m[i].(type) {
		case int:
			fmt.Printf("%s is an integer\n", i)
		case float64:
			fmt.Printf("%s is a float\n", i)
		case []int:
			fmt.Printf("%s is an array\n", i)
		default:
			fmt.Printf("%s is %v\n", i, reflect.TypeOf(m[i]))
		}
	}
}

type user struct { // создание структуры #1
	name, surname string
	age           int
}

type FullNameInterface interface { // создание интерфейса для функции getUserName #1
	getFullName() string
}

func (u *user) getFullName() string { // создание функции для user #1
	return u.name + " " + u.surname
}

type animalInterface interface { // композитный интерфейс
	walkerInterface
	runnerInterface
}

type walkerInterface interface {
	walk()
}

type runnerInterface interface {
	run()
}

type dog struct {
	age int
}

func (d *dog) walk() {
	fmt.Println("dog is walking")
}

func (d *dog) run() {
	fmt.Println("dog is running")
}
