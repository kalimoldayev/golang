package main

import "fmt"

func main() {
	qyeues := Qyeue{}
	qyeues.push(1)
	qyeues.push(2)
	qyeues.push(3)
	qyeues.push(4)
	fmt.Println(qyeues)
	qyeues.dequeue()
	fmt.Println(qyeues)
}

type Qyeue []int

func (q *Qyeue) push (val int) {
	*q = append(*q, val)
}

func (q *Qyeue) dequeue () int {
	if len(*q) != 0 {
		element := (*q)[0]
		*q = (*q)[1:]
		return element
	}
	return -1
}