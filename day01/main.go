package main

import (
	"encoding/xml"
	"fmt"
	"log"
	"os"
)

type DBReader interface {
	parseFile() Recipes
}

// type Recipes struct {
// 	XMLName xml.Name `xml:"recipes"`
// 	Text    string   `xml:",chardata"`
// 	Cake    []struct {
// 		Text        string `xml:",chardata"`
// 		Name        string `xml:"name"`
// 		Stovetime   string `xml:"stovetime"`
// 		Ingredients struct {
// 			Text string `xml:",chardata"`
// 			Item []struct {
// 				Text      string `xml:",chardata"`
// 				Itemname  string `xml:"itemname"`
// 				Itemcount string `xml:"itemcount"`
// 				Itemunit  string `xml:"itemunit"`
// 			} `xml:"item"`
// 		} `xml:"ingredients"`
// 	} `xml:"cake"`
// }

// type Recipes struct {
// 	XMLName xml.Name `xml:"recipes"`
// 	Cake    []struct {
// 		Name        string `xml:"name"`
// 		Stovetime   string `xml:"stovetime"`
// 		Ingredients struct {
// 			Item []struct {
// 				Itemname  string `xml:"itemname"`
// 				Itemcount string `xml:"itemcount"`
// 				Itemunit  string `xml:"itemunit"`
// 			} `xml:"item"`
// 		} `xml:"ingredients"`
// 	} `xml:"cake"`
// }

type Recipes struct {
	XMLName xml.Name `xml:"recipes" json:"recipes,omitempty"`
	Cake    []struct {
		Name        string `xml:"name" json:"name"`
		Time        string `xml:"stovetime" json:"time"`
		Ingredients []struct {
			Item []struct {
				Itemname  string `xml:"itemname"`
				Itemcount string `xml:"itemcount"`
				Itemunit  string `xml:"itemunit"`
			} `xml:"item" json:"item,omitempty"`

			IngredientName  string `json:"ingredient_name"`
			IngredientCount string `json:"ingredient_count"`
			IngredientUnit  string `json:"ingredient_unit,omitempty"`
		} `xml:"ingredients" json:"ingredients"`
	} `xml:"cake" json:"cake"`
}

func xmlA() {
	data, _ := os.ReadFile("recipes.xml")
	var recJson Recipes
	if err := xml.Unmarshal(data, &recJson); err != nil {
		log.Fatal(err)
	}

	modified, err := xml.MarshalIndent(&recJson, "", "    ")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s\n", modified)
}

func main() {
	// data, _ := os.ReadFile("recipes.json")
	// var recJson Recipes
	// if err := json.Unmarshal(data, &recJson); err != nil {
	// 	log.Fatal(err)
	// }

	// modified, err := json.MarshalIndent(&recJson, "", " ")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Printf("%s\n", modified)
	xmlA()
}
