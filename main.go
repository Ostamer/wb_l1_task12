package main

import "fmt"

func main() {
	// Исходная последовательность строк
	strings := []string{"cat", "cat", "dog", "cat", "tree"}

	// Создаем множество
	set := createSet(strings)

	// Выводим результат
	fmt.Println("Множество:", set)
}

// Функция для создания множества из среза строк
func createSet(strings []string) map[string]struct{} {
	// Создаем карту для множества
	set := make(map[string]struct{})

	// Проходим по срезу строк и добавляем в карту
	for _, str := range strings {
		set[str] = struct{}{}
	}

	return set
}
