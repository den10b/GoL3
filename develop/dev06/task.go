package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

func manCut(line string, fields int, delimiter string, separated bool) error {
	if separated {
		if !strings.Contains(line, delimiter) {
			return nil
		}
	}
	words := strings.Split(line, delimiter)
	if len(words) < fields {
		return fmt.Errorf("количество колонок меньше, чем требуется вывести")
	}
	for index, word := range words {
		if index+1 == fields {
			fmt.Println(word)
		}
	}

	return nil
}

func main() {
	fields := flag.Int("f", 1, "выбрать поля (колонки)")
	delimiter := flag.String("d", "	", "использовать другой разделитель")
	separated := flag.Bool("s", false, "только строки с разделителем")

	flag.Parse()

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		if scanner.Text() == "" {
			return
		}
		err := manCut(scanner.Text(), *fields, *delimiter, *separated)
		if err != nil {
			log.Fatal(err)
		}
	}

	if scanner.Err() != nil {
		log.Fatal(scanner.Err())
	}

}
