package main

import (
	"archive/tar"
	"compress/gzip"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
	"sync"
)

func addToArchive(tw *tar.Writer, filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	info, err := file.Stat()
	if err != nil {
		return err
	}
	header, err := tar.FileInfoHeader(info, info.Name())
	if err != nil {
		return err
	}
	header.Name = filename
	err = tw.WriteHeader(header)
	if err != nil {
		return err
	}
	_, err = io.Copy(tw, file)
	if err != nil {
		return err
	}
	return nil
}

func CreateArchive(files []string, buf io.Writer) error {
	gw := gzip.NewWriter(buf)
	defer gw.Close()
	tw := tar.NewWriter(gw)
	defer tw.Close()

	for _, file := range files {
		err := addToArchive(tw, file)
		if err != nil {
			return err
		}
	}

	return nil
}

func	MultiFiles(fileName string) {
	files := []string{fileName}
	info, err := os.Stat(fileName)
	if err != nil {
		log.Fatalln(err)
	}
	timeStamp := strconv.FormatInt(info.ModTime().Unix(), 10)
	tarName := strings.TrimSuffix(fileName, ".log") + timeStamp + ".tar.gz"

	err = os.MkdirAll(strings.TrimPrefix(os.Args[2], "/"), os.ModePerm)
	if err != nil {
		log.Fatalln(err)
	}

	var tarNamePath string
	if strings.HasSuffix(os.Args[2], "/") {
		tarNamePath = os.Args[2]
	} else {
		tarNamePath = os.Args[2] + "/"
	}
	fmt.Println(tarNamePath+tarName)
	out, err := os.Create(strings.TrimPrefix(tarNamePath, "/") + tarName)
	if err != nil {
		log.Fatalln(err)
	}
	defer out.Close()

	err = CreateArchive(files, out)
	if err != nil {
		log.Fatalln(err)
	}
}

func	SingleFile(fileName string) {
	files := []string{fileName}
	info, err := os.Stat(fileName)
	if err != nil {
		log.Fatalln(err)
	}

	timeStamp := strconv.FormatInt(info.ModTime().Unix(), 10)
	tarName := strings.TrimSuffix(fileName, ".log") + timeStamp + ".tar.gz"

	out, err := os.Create(tarName)
	if err != nil {
		log.Fatalln(err)
	}
	defer out.Close()

	err = CreateArchive(files, out)
	if err != nil {
		log.Fatalln(err)
	}
}

func	TagGzCreator(flag string, start int) {
	var wg sync.WaitGroup

	for _, file := range os.Args[start:] {
		if strings.HasSuffix(file, ".log") {
			wg.Add(1)
			file := file

			go func() {
				defer wg.Done()
				switch flag {
				case "a":
					MultiFiles(file)
				case "one":
					SingleFile(file)
				}
			}()
			wg.Wait()
		} else {
			log.Fatalln("Not .log")
		}
	}
}

func	main() {
	files := flag.Bool("a", false, "many files")
	flag.Parse()

	switch {
	case len(os.Args) < 2 || (len(os.Args) == 2 && *files):
		log.Fatalln("Wrong input")
	case *files:
		TagGzCreator("a", 3)
	case !*files:
		TagGzCreator("one", 1)
	}
}