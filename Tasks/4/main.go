package main

import (
	"fmt"
	"sort"
)

func findAnagram(begin *[]string) *map[string]*[]string {
	finMap := make(map[string]*[]string)
	for _, word := range *begin {
		var charArray []rune
		for _, char := range word {
			fmt.Printf("%c ", char)
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
		*finMap[string(sortCharArray)] = append(*finMap[string(sortCharArray)], string(charArray))
		fmt.Println(finMap)
		charArray = make([]rune, 0)
	}
	println(finMap)
	//finMap2 := make(map[string][]string)
	//for key, val := range finMap {
	//	if len(val) <= 1 {
	//		delete(finMap, key)
	//
	//	} else {
	//		tmpKey:=val[0]
	//		for index,word:=range val[1:]{
	//			if index
	//		}
	//	}
	//}
	return &finMap
}
func main() {
	beginStringArray := []string{"боб", "обб", "робб", "денис", "денси", "исден", "денис"}
	mapp := findAnagram(&beginStringArray)
	fmt.Println(mapp)
}
