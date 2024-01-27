package main

import (
	"fmt"
	"strings"
)

func findPresidentMentions(hymn, name string) (int, []int) {
	// Инициализация списка для хранения индексов упоминаний имени президента
	var nameMention []int

	// Проход по тексту гимна для поиска упоминаний имени президента
	index := 0
	for {
		// Поиск индекса первого вхождения имени президента в оставшейся части гимна
		mentionIndex := strings.Index(hymn[index:], name)

		// Если вхождение найдено, добавляем индекс в список и обновляем индекс для следующего поиска
		if mentionIndex != -1 {
			nameMention = append(nameMention, index+mentionIndex)
			index = index + mentionIndex + 1
		} else {
			// Если не найдено больше вхождений, выходим из цикла
			break
		}
	}

	// Возвращаем количество упоминаний и список индексов
	return len(nameMention), nameMention
}

func main() {
	// Ввод данных
	var hymn, presidentName string
	fmt.Scan(&hymn)
	fmt.Scan(&presidentName)

	// Получение результатов и вывод
	count, indices := findPresidentMentions(hymn, presidentName)
	fmt.Println(count)
	for i := 0; i < len(indices); i++ {
		fmt.Print(indices[i], " ")
	}
}
