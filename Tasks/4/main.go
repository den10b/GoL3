package main

import (
	"fmt"
	"sort"
	"strings"
)

func findAnagram(begin *[]string) *map[string]*[]string {
	tempMap := make(map[string][]string)
	checkMap := make(map[string]bool)
	for _, word := range *begin {
		if checkMap[strings.ToLower(word)] {
			continue //проверяем встречалось ли слово
		}
		var charArray []rune
		for _, char := range word {
			charArray = append(charArray, char)
		}
		var sortCharArray = make([]rune, len(charArray))
		copy(sortCharArray, charArray)
		sort.Slice(sortCharArray, func(i, j int) bool { //Сортируем массив рун
			if sortCharArray[i] < sortCharArray[j] {
				return false
			} else {
				return true
			}
		})
		tempMap[strings.ToLower(string(sortCharArray))] = append(tempMap[string(sortCharArray)], strings.ToLower(string(charArray)))
		charArray = make([]rune, 0)
		checkMap[strings.ToLower(word)] = true
	}
	finMap := make(map[string]*[]string)
	for key, val := range tempMap {
		if len(val) <= 1 {
			delete(tempMap, key)

		} else {
			tmpKey := val[0]
			tmpArray := new([]string)
			for _, word := range val {
				*tmpArray = append(*tmpArray, word)
			}
			sort.Strings(*tmpArray)
			finMap[tmpKey] = tmpArray
		}
	}
	return &finMap
}
func main() {
	beginStringArray := []string{"боб", "обб", "робб", "боб", "денис", "денси", "исден", "денис", "пятак", "пятка", "тяпка", "листок", "слиток", "столик"}
	mapp := findAnagram(&beginStringArray)
	fmt.Println(mapp)
	for key, arr := range *mapp {
		fmt.Printf("key: %s -> array: %s\n", key, arr)
	}
}
