package main

import "fmt"

func main() {
	stacks := Stack{}
	stacks.push(15)
	stacks.push(30)
	fmt.Println(stacks)
	last := stacks.pop()
	fmt.Println(last, stacks)
}

type Stack []int


func (s  *Stack) push (val int) { // добавляет в конец
	*s = append(*s, val)
}

func (s *Stack) pop () int { // извлекает самый последний элемент и return
	if len(*s) != 0 {
		index := len(*s) - 1
		element := (*s)[index]
		*s = (*s)[:index]
		return element
	}
	return -1
}