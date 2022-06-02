package main

import (
	"encoding/json"
	"encoding/xml"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
)

type Recipes struct {
	Cake []Cake `json:"cake" xml:"cake"`
}
type Cake struct {
	Name        string `json:"name" xml:"name"`
	Time        string `json:"time" xml:"stovetime"`
	Ingredients []struct {
		XMLName         xml.Name `json:"-" xml:"item"`
		IngredientName  string   `json:"ingredient_name" xml:"itemname"`
		IngredientCount string   `json:"ingredient_count" xml:"itemcount"`
		IngredientUnit  string   `json:"ingredient_unit,omitempty" xml:"itemunit,omitempty"`
	} `json:"ingredients" xml:"ingredients>item"`
}
type DBReader interface {
	myRead(f *os.File) Recipes
}
type myJson struct {
}

type myXML struct {
}

func (m *myXML) myRead(f *os.File) Recipes {
	var res Recipes
	file, err := io.ReadAll(f)
	if err != nil {
		log.Fatal(err)
	}
	err = xml.Unmarshal(file, &res)
	if err != nil {
		log.Fatal(err)
	}
	return res
}
func (m *myJson) myRead(f *os.File) Recipes {
	var res Recipes
	file, err := io.ReadAll(f)
	if err != nil {
		log.Fatal(err)
	}
	err = json.Unmarshal(file, &res)
	if err != nil {
		log.Fatal(err)
	}
	return res
}
func main() {
	var cakeRes Recipes
	var fileValue string
	var j, x DBReader = &myJson{}, &myXML{}
	flag.StringVar(&fileValue, "f", "", "filename")
	flag.Parse()
	if fileValue == "" {
		log.Fatal("none file")
	}
	f, err := os.Open(fileValue)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	switch filepath.Ext(fileValue) {
	case ".xml":
		cakeRes = x.myRead(f)
		FinalRes, err := json.MarshalIndent(cakeRes, "", "    ")
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(string(FinalRes))
	case ".json":
		cakeRes = j.myRead(f)
		FinalRes, err := xml.MarshalIndent(cakeRes, "", "    ")
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(string(FinalRes))
	default:
		log.Fatal("incorrect format")
	}
}
