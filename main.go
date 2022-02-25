package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	f := flag.Bool("l", false, "count lines")
	var s, file string
	flag.StringVar(&s, "c", "", "find number of characters present")
	flag.StringVar(&file, "f", "", "find number of characters present")
	flag.Parse()
	if file == "" {
		fmt.Println(Counter(os.Stdin, *f, s))
	} else {
		_, err := ioutil.ReadFile(file)
		if err != nil {
			log.Fatal("failed to read file :", file)
			os.Exit(1)
		}
		fl, err := os.Open(file)
		if err != nil {
			log.Fatal("failed to open file: ", file)
			os.Exit(1)
		}
		fmt.Println(Counter(fl, *f, s))
	}

}

/*
Counter determines total number of words or lines based on flag provided
and return the count
*/
func Counter(r io.Reader, lines bool, char string) (int, int) {

	scanner := bufio.NewScanner(r) // scanner is used for reading text from io.Reader

	if !lines {
		scanner.Split(bufio.ScanWords) // defining that scanner split type is words
	} else {
		scanner.Split(bufio.ScanLines) // split lines
	}

	wc := 0
	chars := 0
	for scanner.Scan() {
		wc++
		if char != "" { // when charater is provided for a count as flag
			t := scanner.Text() //get text from scanned
			for _, text := range []byte(t) {
				if text == []byte(char)[0] {
					chars++
				}
			}
		}
	}
	return chars, wc
}
