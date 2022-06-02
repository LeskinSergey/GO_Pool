package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

func printRes(fl []os.DirEntry, name string, sl, d, f *bool, ext *string) {
	for i := range fl {
		if fl[i].IsDir() {
			if *d {
				fmt.Println(name + fl[i].Name())
				myFind(sl, d, f, ext, name+fl[i].Name())
			}

		} else {
			link, err := os.Readlink(name + fl[i].Name())
			if err != nil && *f && strings.HasSuffix(fl[i].Name(), *ext) {
				fmt.Printf("%s\n", name+fl[i].Name())
			} else {
				if *sl && link != "" {
					fmt.Printf("%s -> %s\n", name+fl[i].Name(), "[broken]")
				} else if *sl {
					fmt.Printf("%s -> %s\n", name+fl[i].Name(), link)
				}
			}
		}
	}
}

func myFind(sl, d, f *bool, ext *string, name string) {
	_, err := os.Open(os.Args[len(os.Args)-1])
	if err != nil {
		log.Fatal(err)
	}
	if !strings.HasSuffix(name, "/") {
		name = name + "/"
	}
	fl, err := os.ReadDir(name)
	if err != nil {
		log.Fatal(err)
	}
	printRes(fl, name, sl, d, f, ext)
}
func main() {
	sl := flag.Bool("sl", false, "symlink")
	d := flag.Bool("d", false, "directory")
	f := flag.Bool("f", false, "file")
	ext := flag.String("ext", "", "extension")
	flag.Parse()
	if len(os.Args) < 2 {
		log.Fatal("please enter input")
	}
	if *sl && *d && *f == false && *ext == "" {
		*sl, *d, *f = true, true, true
		myFind(sl, d, f, ext, os.Args[1])
	}
	if *ext != "" && (*d == true || *sl == true || *f == false) {
		log.Fatal("wrong input param")
	}
	myFind(sl, d, f, ext, os.Args[len(os.Args)-1])
}
