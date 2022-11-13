package main

import "fmt"

func main() {
    list := linkedList{}
    list.append(0)
    list.append(1)
    list.append(2)
    list.print()
    list.printBack()
}

type node struct {
    data interface{}
    prev *node
    next *node
}

type linkedList struct {
    head *node
    tail *node
    length int
}


func (l * linkedList) append (val interface{}) {
    newNode := node{data: val}

    if l.head == nil { // если список пустой
        l.head = &newNode
        l.tail = &newNode
    } else {
        currentNode := l.head

        for currentNode.next != nil { // пока не равен к хвосту
            currentNode = currentNode.next
        }
        newNode.prev = currentNode // предыдущий элемент
        currentNode.next = &newNode // добавляю в список
        l.tail = &newNode // меняю хвост
    }
    l.length++
}

func (l *linkedList) print() {
    currentNode := l.head // начало

    for currentNode != l.tail { // пока не равен к хвосту
        fmt.Printf("->%v", currentNode.data) // вывожу
        currentNode = currentNode.next
    }
    fmt.Printf("->%v", currentNode.data) // хвост
    fmt.Printf("\n")
}

func (l *linkedList) printBack () {
    length := l.length
    currentNode := l.tail

    for length > 0 {
        fmt.Printf("->%v", currentNode.data) // вывод
        currentNode = currentNode.prev // меняю текущего на предыдущую ноду
        length-- // для итерации
    }
    fmt.Printf("\n")
}

func (l *linkedList) deleteByValue (val interface{}) {
    currentNode := l.head

    if currentNode.data == val { // если head равен
        l.head = currentNode.next
        currentNode.next.prev = nil
        l.length-- // длина
        return
    }

    if l.tail.data == val {
        l.tail = l.tail.prev
        l.tail.next = nil
        l.length--
        return
    }

    for currentNode.next.data != val {
        if currentNode.next.next == nil { // если нету в списке
            return
        }
        currentNode = currentNode.next
    }
    currentNode.next = currentNode.next.next
    currentNode.next.prev = currentNode
    l.length--
}

func (l *linkedList) toArray() interface{} {
    result := make([]interface{}, l.length) // инициализация массива
    currentNode := l.head // начало

    for length:=0; length<=l.length; length++ {
        if currentNode != nil {
            result[length] = currentNode.data // заполняю
            currentNode = currentNode.next
        }
    }

    return result
}

func (l *linkedList) deleteByIndex(index int) {
    if index > l.length - 1 {
        return
    }
    middle := l.length / 2 // как в бинарном поиске получаю середину

    switch {
    case index <= middle:
        currentNode := l.head
        for i := 0; i < index; i++ {
            currentNode = currentNode.next
        }
        if l.head == currentNode {
            if l.length == 1 {
                l.head = nil
                l.tail = nil
                l.length--
                return
            }
            l.head = currentNode.next
            currentNode.next.prev = nil
        } else {
            currentNode.prev.next = currentNode.next
        }
    case index > middle:
        currentNode := l.tail
        for i := l.length - index - 1; i > 0; i-- {
            currentNode = currentNode.prev
        }
        currentNode.prev.next = currentNode.next
        if l.tail == currentNode {
            l.tail = currentNode.prev
        }
    }
    l.length--
}