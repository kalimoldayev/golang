package main

import "fmt"

const ArraySize = 2 // длина таблицы

// структура хеш таблицы
type HashTable struct {
	array [ArraySize]*bucket
}

// ячейки для таблицы, зависит от длины ArraySize
type bucket struct {
	head *bucketNode
}

// сами элементы внутри bucket
type bucketNode struct {
	key  string
	next *bucketNode
}

// инициализация таблицы
func Init() *HashTable {
	result := &HashTable{} // создал пустую таблицу
	for i := range result.array {
		result.array[i] = &bucket{} // вставляю ячейку
	}

	return result
}

// вставка
func (h *HashTable) Insert(key string) {
	index := hash(key) // получаю хеш
	fmt.Println(index, "hash")
	h.array[index].insert(key) // добавляю
}

// поиск
func (h *HashTable) Search(key string) bool {
	index := hash(key)
	return h.array[index].search(key)
}

// удаление
func (h *HashTable) Delete(key string) {
	index := hash(key)
	h.array[index].delete(key)
}

// втавка самих элементов
func (b *bucket) insert(k string) {
	if !b.search(k) { // если в таблице нету по такому индексу
		newNode := &bucketNode{key: k}
		newNode.next = b.head
		b.head = newNode
	} else { // (коллизия)
		for b.head.next != nil { // до тех пор пока не дойду до конца
			b.head.next = b.head.next.next
		}
		newNode := &bucketNode{key: k} // новая нода
		b.head.next = newNode          // добавляю в список
	}
}

func (b *bucket) search(k string) bool {
	currentNode := b.head

	for currentNode != nil {
		if currentNode.key == k {
			return true
		}
		currentNode = currentNode.next
	}
	return false
}

func (b *bucket) delete(k string) {

	if b.head.key == k {
		b.head = b.head.next
		return
	}

	prevNode := b.head
	for prevNode.next != nil {
		if prevNode.next.key == k {
			prevNode.next = prevNode.next.next
		}
		prevNode = prevNode.next
	}
}

func hash(key string) int {
	sum := 0
	for i := range key {
		sum += int(key[i])
	}
	return sum % ArraySize
}

func main() {
	hashTable := Init()
	list := []string{
		"ALISHER",
		"ALISHER",
		"vsvdv",
	}

	for i := range list {
		hashTable.Insert(list[i])
	}

	for i := range hashTable.array {
		fmt.Println(hashTable.array[i].head, hashTable.array[i].head.next)
	}
}

/*
 хеш таблица - это абстрактный тип данных, в go это map, в php ассоциативный массив.
 как работает хеш таблица:

 сначала создаем хеш функцию, на хеш функцию можно представит как черную коробку, куда мы передаем ключ и получаем индекс/хеш значение
 как только получили записываем в таблицу под эти индексом.
 таблица - это массив, если таблица заполнена 75% то нужно  сделать => rehash
 rehash - это когда мы переносим с одной таблицу на новую, именно в этот момент O(n), потому что мы не только переносим а заново для каждого
 элемента получаем новый индекс/хеш, потому что длина таблицы уже изменилась
 почему делаем релокацию? (rehash) для того чтобы хеш таблица работала по O(1)

 коллизия - это когда для разных ключей получаем один индекс/хеш
 чтобы решить можем использовать метод открытой адресации || релокацию || метод цепочек
*/
