package main

import "fmt"

func main() {
	fmt.Println(string(nextGreatestLetter([]byte{'c', 'f', 'j', 'h'}, 'a')))

}
func nextGreatestLetter(letters []byte, target byte) byte {
	// Если целевая буква больше или равна последней букве в массиве,
	// то наименьшей буквой, которая строго больше цели, будет первая буква в массиве
	if target >= letters[len(letters)-1] {
		return letters[0]
	}
	// Инициализируем две переменные: left и right, которые будут указывать на границы поиска
	left, right := 0, len(letters)-1
	// Пока left не больше right, выполняем цикл
	for left < right {
		// Вычисляем середину текущего диапазона
		mid := left + (right-left)/2
		// Если буквы в середине меньше или равно цели, то ищем в правой половине
		if letters[mid] <= target {
			left = mid + 1
		} else {
			// Иначе ищем в левой половине
			right = mid
		}
	}
	// Возвращаем букву по индексу left, которая строго больше цели
	return letters[left]
}
