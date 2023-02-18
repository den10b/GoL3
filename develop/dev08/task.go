package main

import (
	"bufio"
	"fmt"
	goPs "github.com/mitchellh/go-ps"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
	"syscall"
)

func execute(line string, pipeInput string) (string, error) {
	commands := strings.Split(line, "|")
	var err error
	if len(commands) > 1 {
		var lastRes = pipeInput
		for _, commd := range commands {
			lastRes, err = execute(commd, lastRes)
			if err != nil {
				return "", err
			}
		}
		return lastRes, nil
	} else {
		cmdArgs := strings.Split(strings.Trim(line, " "), " ")
		target := strings.Join(cmdArgs[1:], " ")
		if pipeInput != "" {
			target = pipeInput
		}

		switch cmdArgs[0] {
		case "echo":
			return target, nil
		case "pwd":
			return os.Getwd()
		case "cd":

			if target == ".." {
				curr, _ := os.Getwd()
				return "", os.Chdir(filepath.Dir(curr))
			}
			curr, _ := os.Getwd()
			return "", os.Chdir(filepath.Join(curr, target))
		case "ps":
			prcs, _ := goPs.Processes()
			prcsSB := strings.Builder{}
			for _, prc := range prcs {
				prcsSB.WriteString(fmt.Sprintf("%v  -  %v\n", prc.Executable(), prc.Pid()))
			}
			return prcsSB.String(), nil
		case "kill":
			pid, err := strconv.Atoi(target)
			if err != nil {
				return "", err
			}
			process, err := os.FindProcess(pid)
			if err != nil {
				return "", err
			}
			return "", process.Kill()
		case "exec":
			// LookPath gets absolute address; the parameters can be absolute paths or relative paths.
			// binary should be an executable file, for the global command, you also need to find the location of its specific command file.
			binary, err := exec.LookPath(target)
			if err != nil {
				return "", err
			}
			// No parameters
			args := []string{""}
			// Environment variable using the current process
			env := os.Environ()
			// Execute and enter a new program
			return "", syscall.Exec(binary, args, env)
		default:
			return "", fmt.Errorf("команда %s не поддерживается", cmdArgs[0])
		}

	}
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
			fmt.Println(err)
		}
	}

	if scanner.Err() != nil {
		log.Fatal(scanner.Err())
	}

}
