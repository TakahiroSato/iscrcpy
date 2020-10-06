package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	optionKeys := []string{"--serial", "--window-title", "--max-size", "--bit-rate", "--max-fps"}
	options := map[string]string{
		optionKeys[0]: "",
		optionKeys[1]: "",
		optionKeys[2]: "",
		optionKeys[3]: "",
		optionKeys[4]: "",
	}
	scanner := bufio.NewScanner(os.Stdin)

	for _, k := range optionKeys {
		fmt.Print(k + " : ")
		scanner.Scan()
		options[k] = scanner.Text()
		if options[k] == "" {
			delete(options, k)
		}
	}
	var args []string
	for k, v := range options {
		args = append(args, k, v)
	}

	var extraOptions []string
	if len(os.Args) > 1 {
		for i := 1; i < len(os.Args); i++ {
			extraOptions = append(extraOptions, os.Args[i])
		}
	}

	args = append(args, extraOptions...)

	fmt.Println("execute :", "scrcpy "+strings.Join(args, " "))

	cmd := exec.Command("scrcpy", args...)
	stroud, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	cmd.Start()
	cmdScanner := bufio.NewScanner(stroud)
	for cmdScanner.Scan() {
		fmt.Println(cmdScanner.Text())
	}
	cmd.Wait()
}
