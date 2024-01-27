package main

import (
	"fmt"
	"strings"
)

func findPresidentMentions(hymn, name string) (int, []int) {
	
	var nameMention []int

	
	index := 0
	for {
		
		mentionIndex := strings.Index(hymn[index:], name)

		
		if mentionIndex != -1 {
			nameMention = append(nameMention, index+mentionIndex)
			index = index + mentionIndex + 1
		} else {
			// Если не найдено больше вхождений, выходим из цикла
			break
		}
	}

	
	return len(nameMention), nameMention
}

func main() {
	
	var hymn, presidentName string
	fmt.Scan(&hymn)
	fmt.Scan(&presidentName)

	
	count, indices := findPresidentMentions(hymn, presidentName)
	fmt.Println(count)
	for i := 0; i < len(indices); i++ {
		fmt.Print(indices[i], " ")
	}
}
