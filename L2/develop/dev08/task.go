package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"math"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"sync"

	"github.com/shirou/gopsutil/load"
)

/*
=== Взаимодействие с ОС ===
Необходимо реализовать собственный шелл
встроенные команды: cd/pwd/echo/kill/ps
поддержать fork/exec команды
Программа должна проходить все тесты. Код должен проходить проверки go vet и go lint.
*/

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("$ ")
		cmdString, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
		err = runCommand(cmdString, os.Stdout)
	}
}

func runCommand(commandStr string, w io.Writer) error {
	commandStr = strings.TrimSuffix(commandStr, "\n")
	arrCommandStr := strings.Fields(commandStr)
	if len(arrCommandStr) == 0 {
		return errors.New("incorrect input")
	}
	switch arrCommandStr[0] {
	case "exit":
		os.Exit(0)
	case "cd":
		cd(arrCommandStr)
	case "pwd":
		pwd(w)
	case "echo":
		echo(arrCommandStr, w)
	case "kill":
		kill(arrCommandStr, w)
	case "ps":
		ps(w)
	case "exec":
		ownExec(arrCommandStr, w)
	case "fork":
		fork(arrCommandStr, w)
	default:
		fmt.Println("This command doesn't exists.")
	}
	cmd := exec.Command(arrCommandStr[0], arrCommandStr[1:]...)
	cmd.Stderr = os.Stderr
	cmd.Stdout = w
	return cmd.Run()
}

func cd(arr []string) {
	if len(arr) < 2 {
		fmt.Println("Not enough arguments for cd function")
	} else {
		os.Chdir("/" + arr[1])
	}
}

func pwd(w io.Writer) {
	_, err := os.Getwd()
	if !check(err, w) {
		return
	}
	fmt.Fprintln(w, "Current Working Direcotry: ")
}

func echo(arr []string, w io.Writer) {
	if len(arr) < 2 {
		fmt.Fprintln(w, "Not enough arguments for cd function")
	}
}

func kill(arr []string, w io.Writer) {
	if len(arr) != 2 {
		return
	}

	pid, err := strconv.Atoi(arr[1])
	check(err, w)
	fmt.Fprintf(w, "PID: %d will be killed.\n", pid)
	proc, err := os.FindProcess(pid)

	if !check(err, w) {
		return
	}
	proc.Kill()
}

func ps(w io.Writer) {
	miscStat, _ := load.Misc()
	fmt.Fprintf(w, "Running processes: %d\n", miscStat.ProcsRunning)
}

func fork(arr []string, w io.Writer) {
	if len(arr) < 2 {
		return
	}
	var wg sync.WaitGroup
	var mux sync.Mutex
	var forkCounter int

	for _, str := range arr {
		if str == "fork" {
			forkCounter += 1
		}
	}
	str := strings.Join(arr[forkCounter:], " ")
	for i := 0; i < int(math.Pow(2, float64(forkCounter))); i++ {
		wg.Add(1)
		go func() {
			mux.Lock()
			err := runCommand(str, w)
			mux.Unlock()
			check(err, w)
			wg.Done()
		}()
	}
	wg.Wait()
}

func ownExec(arr []string, w io.Writer) {
	if len(arr) != 2 {
		return
	}
	dataFromFile, err := os.ReadFile(arr[1])
	if !check(err, w) {
		return
	}
	lines := strings.Split(string(dataFromFile), "\n")

	for _, command := range lines {
		err = runCommand(command, w)
		check(err, w)
	}
}

func check(err error, w io.Writer) bool {
	if err != nil {
		fmt.Fprintln(w, err)
		return false
	}
	return true
}
