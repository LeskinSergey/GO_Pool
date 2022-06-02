package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"sync"
	"unicode/utf8"
)

func MyWcL(name string, wg *sync.WaitGroup) {
	defer wg.Done()
	count := 0
	fl, err := os.Open(name)
	if err != nil {
		log.Fatal(err)
	}
	defer fl.Close()
	scan := bufio.NewScanner(fl)
	for scan.Scan() {
		count++
	}
	fmt.Printf("%d\t%s\n", count, name)
}
func MyWcM(name string, wg *sync.WaitGroup) {
	defer wg.Done()
	count := 0
	fl, err := os.Open(name)
	if err != nil {
		log.Fatal(err)
	}
	defer fl.Close()
	scan := bufio.NewScanner(fl)
	for scan.Scan() {
		count += utf8.RuneCountInString(scan.Text())
	}
	fmt.Printf("%d\t%s\n", count, name)
}
func MyWcW(name string, wg *sync.WaitGroup) {
	defer wg.Done()
	count := 0
	fl, err := os.Open(name)
	if err != nil {
		log.Fatal(err)
	}
	defer fl.Close()
	scan := bufio.NewScanner(fl)
	for scan.Scan() {
		count += len(strings.Split(scan.Text(), " "))
	}
	fmt.Printf("%d\t%s\n", count, name)
}

func CheckFlags(l, m, w *bool) bool {
	count := 0
	if *l {
		count++
	}
	if *m {
		count++
	}
	if *w {
		count++
	}
	if count != 1 {
		return false
	}
	return true
}
func main() {
	l := flag.Bool("l", false, "count lines")
	m := flag.Bool("m", false, "count characters")
	w := flag.Bool("w", false, "count words")
	flag.Parse()
	if CheckFlags(l, m, w) {
		fl := flag.Args()
		wg := new(sync.WaitGroup)
		for _, i := range fl {
			if *l {
				wg.Add(1)
				go MyWcL(i, wg)
			}
			if *m {
				wg.Add(1)
				go MyWcM(i, wg)
			}
			if *w {
				wg.Add(1)
				go MyWcW(i, wg)
			}
			wg.Wait()
		}
	} else {
		log.Fatal("error! more than one flag")
	}
}
