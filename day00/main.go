package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/montanaflynn/stats"
)

var (
	nums = make([]float64, 0)
)

func Mean() {
	mean, _ := stats.Mean(nums)
	fmt.Printf("Mean: %.2f\n", mean)
}

func Median() {
	median, _ := stats.Median(nums)
	fmt.Printf("Median: %.2f\n", median)
}

func Mode() {
	mode, _ := stats.Mode(nums)
	if len(mode) == 0 {
		mode_, _ := stats.Min(nums)
		fmt.Printf("Mode: %.2f\n", mode_)
	} else {
		fmt.Printf("Mode: %.2f\n", mode[0])
	}
}

func StandardDeviation() {
	sd, _ := stats.StandardDeviation(nums)
	fmt.Printf("SD: %.2f\n", sd)
}

func main() {
	flag.Parse()

	var d float64
loop:
	for {
		if _, err := fmt.Scan(&d); err != nil {
			switch {
			case err.Error() == "EOF":
				break loop
			default:
				log.Fatalf("%v", "Wrong input\n")
			}
		}
		if d > 100000 || d < -100000 {
			log.Fatalf("%v", "Out of range\n")
		}
		nums = append(nums, d)
	}
	if len(os.Args) == 1 {
		Mean()
		Median()
		Mode()
		StandardDeviation()
		os.Exit(0)
	}
	for _, cmd := range os.Args[1:] {
		switch cmd {
		case "Mean":
			Mean()
		case "Median":
			Median()
		case "Mode":
			Mode()
		case "SD":
			StandardDeviation()
		default:
			log.Fatalf("%v", "Wrong metric\n")
		}
	}
}
