package main

import "fmt"

func main() {
    array := []int{1,2,3,4,5,6}
    item := 6
    fmt.Println(binarySearch(item, array))
}

func binarySearch(item int, array []int) []int {
    start := 0               // начинаем с 0
    end := len(array) - 1    // конец
    result := []int{-1, -1}  // дефолт результат

    if (item > array[end]) {
        return result        // выход
    }

    for start <= end {
        middle := (start + end) / 2                 // делим на 2 (таким образом получим середину)

        if (array[middle] == item) {                // начинаем проверять с середины
            result = []int{array[middle], middle}   // val, key
            break
        } else if (item > array[middle]) {          // если меньше, то добавляем к middle, именно middle а не start. если добавим к start то это линейный поиск O(n)
            start = middle + 1
        } else if (item < array[middle]) {          // если меньше, то отнимаем
            end = middle - 1
        }
    }
     return result
}

/*
время выполнение сортировки О(Log n) - логарифмические время (бинарный поиск)
*/