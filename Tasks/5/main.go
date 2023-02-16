package main

import (
	"bufio"
	"container/list"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

func manGrep(lines []string, target string, after int, before int, context int, count bool, ignoreCase bool, invert bool, fixed bool, lineNum bool) error {
	counter := 0
	if context > after {
		after = context
	}
	if context > before {
		before = context
	}
	if ignoreCase {
		target = strings.ToLower(target)
	}
	oldStringsQueue := list.New()
	afterCounter := 0
	for index, line := range lines {
		newLine := line
		if ignoreCase {
			newLine = strings.ToLower(line)
		}
		if ((strings.Contains(newLine, target) && !fixed) || ((newLine == target) && fixed)) != invert {
			liner := index - oldStringsQueue.Len() + 1
			for e := oldStringsQueue.Front(); e != nil; e = e.Next() {
				if lineNum {
					fmt.Printf("%v ", liner)
				}
				fmt.Println(e.Value)
				liner++
			}

			if count {
				counter++
			} else {
				if lineNum {
					fmt.Printf("%v ", index+1)
				}
				fmt.Println(line)
			}
			oldStringsQueue.Init()
			afterCounter = after

		} else {

			oldStringsQueue.PushBack(fmt.Sprintf("%v", line))
			for oldStringsQueue.Len() > before {
				first := oldStringsQueue.Front()
				oldStringsQueue.Remove(first)
			}
			if afterCounter > 0 {
				if lineNum {
					fmt.Printf("%v ", index+1)
				}
				back := oldStringsQueue.Back()
				oldStringsQueue.Remove(back)
				fmt.Println(line)
				afterCounter--
			}
		}

	}
	if count {
		fmt.Printf("Количество строк с совпадениями - %v", counter)
	}
	return nil
}

func main() {
	after := flag.Int("A", 0, "печатать +N строк после совпадения")
	before := flag.Int("B", 0, "печатать +N строк до совпадения")
	context := flag.Int("C", 0, "печатать ±N строк вокруг совпадения")
	count := flag.Bool("c", false, "количество строк")
	ignoreCase := flag.Bool("i", false, "игнорировать регистр")
	invert := flag.Bool("v", false, "вместо совпадения, исключать")
	fixed := flag.Bool("F", false, "точное совпадение со строкой, не паттерн")
	lineNum := flag.Bool("n", false, "напечатать номер строки")
	flag.Parse()

	lastArg := os.Args[len(os.Args)-1]
	//lastArg := "Tasks\\5\\example.txt"
	targetString := os.Args[len(os.Args)-2]
	//targetString := "ing"
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
	err = file.Close()
	if err != nil {
		log.Fatal(err)
	}

	err = manGrep(stringsArr, targetString, *after, *before, *context, *count, *ignoreCase, *invert, *fixed, *lineNum)
	if err != nil {
		log.Fatal(err)
	}

}
