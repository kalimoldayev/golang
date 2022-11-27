package main

import "fmt"

func main() {
	nums := []int{15, -1, 19, 5, 6, 8, -9, 16}
	nums2 := []int{15, -1, 19, 5, 6, 8, -9, 16}
	fmt.Println(quickSort(nums), j)
	fmt.Println(quickSort2(nums2), j2)
}

var (
	j, j2 int
)

func quickSort(nums []int) []int {
	j++

	if len(nums) < 2 {
		return nums
	}

	pivotIndex := len(nums) / 2

	for i := range nums {
		j++
		if nums[i] > nums[pivotIndex] {
			nums[i], nums[pivotIndex] = nums[pivotIndex], nums[i]
		}

		if pivotIndex < len(nums)-1 {
			pivotIndex++
		}
	}

	quickSort(nums[:pivotIndex-1])
	quickSort(nums[pivotIndex+1:])
	return nums
}

func quickSort2(nums []int) []int {
	j2++

	if len(nums) < 2 {
		return nums
	}

	pivotIndex := len(nums) / 2
	less := []int{}
	more := []int{}

	for i := range nums {
		j2++
		if nums[i] > nums[pivotIndex] {
			more = append(more, nums[i])
		}
		if nums[i] < nums[pivotIndex] {
			less = append(less, nums[i])
		}
	}

	result := quickSort2(less)

	result = append(result, nums[pivotIndex])
	for _, v := range quickSort2(more) {
		result = append(result, v)
	}

	return result
}

/*
 быстрая сортировка работает по принципу "разделяй и властвуй" .мы сначала определяем опорный элемент,
 это очень важно какой элемент мы выберем. от этого зависит время выполнения алгоритма:
 средний случай: O(nLogn)
 худший случай: O(n2)

 а вот сортировка слиянием работает за O(nLogn). у быстрой сортировки меньше константы чем сортировки слиянием.
 быстрая сортировка не знает массив уже отсортирован или нет, и если мы дадим уже отсортированный массив
 и выберем первый элемент как опорный, стек вызовов получается длинным и это = O(n)

 если выберем опорный элемент = array[len(array) / 2] таким образом мы быстрее получим базовый случай и намного короче стек вызовов
 в этом примере существует O(Logn) ( технической точки правильнее сказать "высота стека вызовов равна O(Logn)" уровней).
 а так как каждый уровень занимает время O(n), то весь алгоритм займет O(n) * (Logn) = O(nLogn)
*/
