package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	fmt.Println(findTwoEntries())
	fmt.Println(findThreeEntries())
}

func findTwoEntries() int {
	var d []int

	file, err := os.Open("../dec1/data.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	s := bufio.NewScanner(file)
	for s.Scan() {
		txt, err := strconv.Atoi(s.Text())
		if err != nil {
			log.Fatal(err)
		}
		d = append(d, txt)
	}

	for _, x := range d {
		for _, y := range d {
			if x+y == 2020 {
				return x * y
			}
		}
	}
	return 0
}

func findThreeEntries() int {
	var d []int

	file, err := os.Open("../dec1/data.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	s := bufio.NewScanner(file)
	for s.Scan() {
		txt, err := strconv.Atoi(s.Text())
		if err != nil {
			log.Fatal(err)
		}
		d = append(d, txt)
	}

	for _, x := range d {
		for _, y := range d {
			for _, z := range d {
				if x+y+z == 2020 {
					return x * y * z
				}
			}
		}
	}
	return 0
}
