package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"sync"
)

func	CountWords(file string) {
	fileHandle, err := os.Open(file)
	if err != nil {
		log.Fatalln(err)
	}
	defer fileHandle.Close()
	fileScanner := bufio.NewScanner(fileHandle)
	fileScanner.Split(bufio.ScanWords)
	count := 0
	for fileScanner.Scan() {
		count++
	}
	if err := fileScanner.Err(); err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("%d %s\n", count, file)
}

func	CountChars(file string) {
	fileHandle, err := os.Open(file)
	if err != nil {
		log.Fatalln(err)
	}
	scanner := bufio.NewScanner(fileHandle)
	scanner.Split(bufio.ScanRunes)
	count := 0
	for scanner.Scan() {
		count ++
	}
	if err := scanner.Err(); err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("%d %s\n", count, file)
}

func	CountLines(file string) {
	fileHandle, err := os.Open(file)
	if err != nil {
		log.Fatalln(err)
	}
	defer fileHandle.Close()
	scanner := bufio.NewScanner(fileHandle)
	scanner.Split(bufio.ScanLines)
	count := 0
	for scanner.Scan() {
		count++
	}
	if err := scanner.Err(); err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("%d %s\n", count, file)
}

func	CountSpecs(flag string, start int) {
	var wg sync.WaitGroup

	for _, file := range os.Args[start:] {
		wg.Add(1)
		file := file

		go func() {
			defer wg.Done()
			switch flag {
				case "words":
					CountWords(file)
				case "chars":
					CountChars(file)
				case "lines":
					CountLines(file)
			}
		}()
		wg.Wait()
	}

}

func main(){
	lines := flag.Bool("l", false, "lines counting")
	chars := flag.Bool("m", false, "chars counting")
	words := flag.Bool("w", false, "words counting")

	flag.Parse()
	switch {
		case len(os.Args) == 1 || (*lines && *chars) || (*lines && *words) || (*chars && *words):
			log.Fatalln("Wrong input")
		case !*words && !*lines && !*chars:
			CountSpecs("words", 1)
		case *words:
			CountSpecs("words", 2)
		case *chars:
			CountSpecs("chars", 2)
		case *lines:
			CountSpecs("lines", 2)
	}
}