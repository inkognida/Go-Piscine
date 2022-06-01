package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

func myFind(path string, f, d, l bool, extension string) {
	if !strings.HasSuffix(path, "/") {
		path = path + "/"
	}
	files, err := os.ReadDir(path)
	if err != nil {
		log.Fatalf("%v", err)
	}
	for i := range files {
		if files[i].IsDir() {
			if d {
				fmt.Println(path + files[i].Name())
			}
			myFind(path+files[i].Name(), f, d, l, extension)
		} else {
			link, err := os.Readlink(path + files[i].Name())
			if err == nil {
				_, errLink := os.Open(path + link)
				if l && errLink != nil {
					fmt.Printf("%s -> %s\n", path+files[i].Name(), "[broken]")
				} else if l {
					fmt.Printf("%s -> %s\n", path+files[i].Name(), link)
				}
			} else if f && strings.HasSuffix(files[i].Name(), extension) {
				fmt.Println(path + files[i].Name())
			}
		}
	}
}

func main() {
	if len(os.Args) == 1 {
		log.Fatalf("%v", "Wrong input\n")
	}
	files := flag.Bool("f", false, "files")
	dirs := flag.Bool("d", false, "directories")
	links := flag.Bool("sl", false, "symbol links")
	ext := flag.String("ext", "", "files extension")

	flag.Parse()
	if !*files && !*dirs && !*links {
		myFind(os.Args[1], true, true, true, *ext)
	} else {
		myFind(os.Args[len(os.Args)-1], *files, *dirs, *links, *ext)
	}
}
