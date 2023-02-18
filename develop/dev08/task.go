package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func execute(line string, pipeInput string) (string, error) {
	commands := strings.Split(line, "|")
	if len(commands) > 1 {
		var lastArg = ""
		for _, commd := range commands {
			execute(commd, lastArg)
		}
		fmt.Println(lastArg)
	}
	//newDir, err := os.Getwd()
	//if err != nil {
	//}
	//fmt.Printf("Current Working Direcotry: %s\n", newDir)
	//target := "Tasks"
	//os.Chdir(filepath.Join(newDir, target))
	//newDir, err = os.Getwd()
	//if err != nil {
	//}
	//fmt.Printf("Current Working Direcotry: %s\n", newDir)
	//os.Chdir(filepath.Dir(newDir))
	//
	//newDir, err = os.Getwd()
	//if err != nil {
	//}
	//fmt.Printf("Current Working Direcotry: %s\n", newDir)
	return "OK", nil
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		if scanner.Text() == "\\quit" {
			return
		}
		res, err := execute(scanner.Text(), "")
		if res != "" {
			fmt.Println(res)
		}
		if err != nil {
			log.Fatal(err)
		}
	}

	if scanner.Err() != nil {
		log.Fatal(scanner.Err())
	}

}
