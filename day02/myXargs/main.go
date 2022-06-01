package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"
)

func	main() {
	var args []string

	args = append(args, os.Args[2:]...)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		args = append(args, scanner.Text())
	}
	cmd := exec.Command(os.Args[1], args...)

	var out bytes.Buffer
	cmd.Stdout = &out

	err := cmd.Run()
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(out.String())
}