package main

import "fmt"

func main() {
    nums := []int{0, 4, 5, 3, 1, 6, 2}  // создал слайс
    sortSelection(nums)                 // вызов функции
}

func sortSelection(nums []int) {
    fmt.Println("Before", nums) // до

    for i := 0; i<(len(nums)-1); i++ {
        min := i                         // берем индекс итерируемого элемента как индекс минимального. таким образом проходимся по массиву и сравниваем все элементы

        for j:=i+1; j<(len(nums)); j++{  // проходимся по второму разу и пропускаем итерируемого элемента i
            if nums[min] > nums[j] {     // если больше, то меняем значение минимального
                min = j
            }
        }

        if min != i {                    // если не совпадают
            temp := nums[i]
            nums[i] = nums[min]          // здесь меняем местами элементы
            nums[min] = temp
        }
    }

    fmt.Println("After", nums)           // после
}

/*
время выполнение сортировки O(n2). потому что проходимся 2 раза по массиву
*/