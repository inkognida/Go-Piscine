package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

// difference returns the elements in `a` that aren't in `b`.
func difference(a, b []string) []string {
	mb := make(map[string]struct{}, len(b))
	for _, x := range b {
		mb[x] = struct{}{}
	}
	var diff []string
	for _, x := range a {
		if _, found := mb[x]; !found {
			diff = append(diff, x)
		}
	}
	return diff
}

func readFile(fileName string) []string {
	file, err := os.Open(fileName)

	if err != nil {
		log.Fatalf("failed opening file: %s", err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var txtlines []string

	for scanner.Scan() {
		txtlines = append(txtlines, scanner.Text())
	}

	file.Close()
	return txtlines
}

func main() {
	switch {
	case len(os.Args) != 5 || os.Args[1] != "--old" || os.Args[3] != "--new":
		log.Fatalf("%v", "Wrong input\n")
	case strings.HasSuffix(os.Args[2], ".txt") && strings.HasSuffix(os.Args[4], ".txt"):
		old := readFile(os.Args[2])
		new := readFile(os.Args[4])
		for _, v := range difference(new, old) {
			fmt.Println("ADDED", v)
		}
		for _, v := range difference(old, new) {
			fmt.Println("REMOVED", v)
		}
	}

}
