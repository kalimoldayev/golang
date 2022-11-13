package main

import "fmt"

func main() {
    list := linkedList{}
    list2 := linkedList{}
    list.prepend(18)
    list.prepend(19)

    list2.append(21)
    list2.append(22)
    list2.append(23)
    list2.append(24)
    list2.printList()

    list2.deleteNodeByValue(24)
    list2.printList()
}

type node struct { // описание node (узла)
    data int       // данные
    next *node     // указатель
}

type linkedList struct {
    head *node     // указатель на 1 элемент списка
    length int     // длина списка
}

func (l *linkedList) prepend (val int) {  // функция которая относятся к linkedList
    newNode := node{data: val} // создаю новую ноду
    if  l.length != 0 { // если список не пустой
        newNode.next = l.head // меняю у текущего нода next
    }

    l.head = &newNode  // меняю head
    l.length++         // длина для итерации
}

func (l *linkedList) append (val int) { // добавляет в конец
    newNode := node{data: val} // создаю новую ноду
    length := l.length // для итерации
    currentNode := l.head // текущая нода

    if l.head == nil { // если список пустой
        l.head = &newNode // head
        l.length++
    } else {
        for length > 0 {
            if currentNode.next == nil { // нахожу у какого пустой next
                currentNode.next = &newNode // добавляю
                l.length++ // длина списка
            }
            currentNode = currentNode.next
            length-- // для итерации
        }
    }
}

func (l *linkedList) printList() {
    length := l.length     // длина по которому буду итерировать
    currentNode := l.head  // стартовый node

    for length > 0 { // пока больше
        fmt.Printf("->%d", currentNode.data) // вывод
        currentNode = currentNode.next // меняю текущего на следующую ноду
        length-- // для итерации
    }
    fmt.Printf("\n")
}

func (l *linkedList) deleteNodeByValue(val int) {
    if l.length == 0 { // если список пустой
        return
    }

    currentNode := l.head
    if currentNode.data == val { // если head равен
        l.head = currentNode.next
        l.length-- // длина
        return
    }

    for currentNode.next.data != val {
        if currentNode.next.next == nil { // если нету в списке
            return
        }
        currentNode = currentNode.next // меняю текущего на следующую ноду
    }
    currentNode.next = currentNode.next.next
    l.length-- // длина
}