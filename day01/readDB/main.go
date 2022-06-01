package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"log"
	"os"
	"strings"
)

type DBReader interface {
	parseFile() Recipes
}

type Recipes struct {
	XMLName xml.Name `xml:"recipes" json:"-"`
	Cake    []struct {
		Name        string `xml:"name" json:"name"`
		Time        string `xml:"stovetime" json:"time"`
		Ingredients []struct {
			IngredientName  string `xml:"itemname" json:"ingredient_name"`
			IngredientCount string `xml:"itemcount" json:"ingredient_count"`
			IngredientUnit  string `xml:"itemunit" json:"ingredient_unit,omitempty"`
		} `xml:"ingredients>item" json:"ingredients"`
	} `xml:"cake" json:"cake"`
}

type JsonReader struct {
	FileName string
}

type XmlReader struct {
	FileName string
}

func (r *JsonReader) parseFile() Recipes {
	data, err := os.ReadFile(r.FileName)
	if err != nil {
		log.Fatalf("%v", err)
	}
	var Cakes Recipes
	if err := json.Unmarshal(data, &Cakes); err != nil {
		log.Fatal(err)
	}

	modified, err := xml.MarshalIndent(&Cakes, "", "    ")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", modified)
	return Cakes
}

func (r *XmlReader) parseFile() Recipes {
	data, _ := os.ReadFile(r.FileName)
	var Cakes Recipes
	if err := xml.Unmarshal(data, &Cakes); err != nil {
		log.Fatal(err)
	}

	modified, err := json.MarshalIndent(&Cakes, "", "    ")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", modified)
	return Cakes
}

func main() {
	if len(os.Args) != 3 || os.Args[1] != "-f" {
		log.Fatalf("%v", "Wrong input\n")
	}
	switch {
	case strings.HasSuffix(os.Args[2], ".json"):
		jsonR := new(JsonReader)
		jsonR.FileName = os.Args[2]
		jsonR.parseFile()
	case strings.HasSuffix(os.Args[2], ".xml"):
		xmlR := new(XmlReader)
		xmlR.FileName = os.Args[2]
		xmlR.parseFile()
	default:
		log.Fatalf("%v", "Wrong filename\n")
	}
}
