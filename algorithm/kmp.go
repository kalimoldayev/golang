package main

import "fmt"

func main() {
	haystack := " aabaabaaaaaabaabaabaabbaaab"
	needle := "aabaab"
	fmt.Println(searchBase(needle, haystack))
	fmt.Println(searchKmp(needle, haystack))
}

func searchBase(needle string, haystack string) []int {
	result := []int{}
	
	for i := range haystack {
		j := 0
		for j < len(needle) && i + j < len(haystack) && haystack[i+j] == needle[j] {
			j++
		}

		if j == len(needle) {
			result = append(result, i)
		}
	}
	return result
}

/*
 время выполнения O(n*m)
 линейное время
*/  


func getPrefix(text string) []int {
	prefix := make([]int, len(text))
	
	for i:=1; i<len(text); i++ {
		j := prefix[i - 1] // j равен к пред. элемента с префиксного массива
	
		for j > 0 && text[j] != text[i] { // до тех пор пока элементы по ключу j & i не равны или j не равен к 0
			j = prefix[j - 1] // то меняем значение к пред. элемента с префиксного массива
		}
		if text[i] == text[j] { // если равен, то добавляем единицу
			j++
		}
		prefix[i] = j // в конце устанавливаем новое значение
	}

	/*
	 первый элемент префиксной функции равен к 0, поэтому пропускаеи начинаем с 1 (i = 1)
	*/

	return prefix
}


func searchKmp(needle string, haystack string) []int {
	result := []int{}
	prefix := getPrefix(needle)
	i := 0 // ключ для строки
	j := 0 // ключ для подстроки

	for i < len(haystack) {
		if needle[j] == haystack[i] {
			j++
			i++
		}
		if j == len(needle) {
			result = append(result, i - j)
			j = prefix[j - 1]
		} else if i < len(haystack) && needle[j] != haystack[i] {
			if j != 0 {
				j = prefix[j - 1]
			} else {
				i++
			}
		}
	}

	return result
}

/*
 можно разделить на 2 этапа КМП:
 1) формирование массива (префикс)
 2) посик в строке
*/

/*
 префикс - подстрока, начинающаяся с первого символа строки
 суффукс - подстрока, заканчивающаяся на последний символ строки
 строка  = префикс + суффукс
*/