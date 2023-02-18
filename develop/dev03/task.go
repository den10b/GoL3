package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	str "strings"
)

func manSort(strings []string, sortColoumn int, num bool, reverse bool, nonRepeat bool) error {
	if sortColoumn >= 0 {
		sort.Slice(strings, func(i, j int) bool {
			wordsI := str.Split(strings[i], " ")[sortColoumn]
			wordsJ := str.Split(strings[j], " ")[sortColoumn]
			var res bool
			if num {
				intI, err := strconv.Atoi(wordsI)
				if err != nil {
					log.Fatal(err)
				}
				intJ, err := strconv.Atoi(wordsJ)
				if err != nil {
					log.Fatal(err)
				}
				res = intI <= intJ
			} else {
				res = wordsI <= wordsJ
			}
			return res != reverse
		})
	} else {
		sort.Slice(strings, func(i, j int) bool {
			res := strings[i] <= strings[j]
			return res != reverse
		})
	}
	if nonRepeat {
		for i := 1; i < len(strings); i++ {
			if strings[i] == strings[i-1] {
				strings = append(strings[:i], strings[i+1:]...)
			}
		}
	}
	return nil
}

func main() {
	sortColoumn := flag.Int("k", 0, "указание колонки для сортировки")
	num := flag.Bool("n", false, "сортировать по числовому значению")
	reverse := flag.Bool("r", false, "сортировать в обратном порядке")
	nonRepeat := flag.Bool("u", false, "не выводить повторяющиеся строки")
	flag.Parse()

	lastArg := os.Args[len(os.Args)-1]
	lastArg = "Tasks\\3\\unsort.txt"
	filePath := fmt.Sprintf("%s", lastArg)
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	stringsArr := []string{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		stringsArr = append(stringsArr, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	file.Close()
	fmt.Println(stringsArr)
	err = manSort(stringsArr, *sortColoumn, *num, *reverse, *nonRepeat)
	if err != nil {
		log.Fatal(err)
	}
	file, _ = os.Create("Tasks\\3\\sort.txt")

	for _, line := range stringsArr {
		fmt.Fprintln(file, line)
	}
	file.Close()
}
