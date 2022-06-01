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

type XmlReader struct {
	FileName string
}

type JsonReader struct {
	FileName string
}

type Recipes struct {
	XMLName xml.Name `xml:"recipes" json:"-"`
	Cake    []struct {
		Name        string `xml:"name" json:"name"`
		Time        string `xml:"stovetime" json:"time"`
		Ingredients []struct {
			Itemname  string `xml:"itemname" json:"ingredient_name"`
			Itemcount string `xml:"itemcount" json:"ingredient_count"`
			Itemunit  string `xml:"itemunit" json:"ingredient_unit,omitempty"`
		} `xml:"ingredients>item" json:"ingredients"`
	} `xml:"cake" json:"cake"`
}

func (r *JsonReader) parseFile() Recipes {
	data, _ := os.ReadFile(r.FileName)
	var recipes Recipes
	if err := json.Unmarshal(data, &recipes); err != nil {
		log.Fatal(err)
	}
	return recipes
}

func (r *XmlReader) parseFile() Recipes {
	data, _ := os.ReadFile(r.FileName)
	var recipes Recipes
	if err := xml.Unmarshal(data, &recipes); err != nil {
		log.Fatal(err)
	}
	return recipes
}

func getReader(fileName string) DBReader {
	switch {
	case strings.HasSuffix(fileName, ".json"):
		reader := new(JsonReader)
		reader.FileName = fileName
		return reader
	case strings.HasSuffix(fileName, ".xml"):
		reader := new(XmlReader)
		reader.FileName = fileName
		return reader
	default:
		log.Fatalf("%v", "Wrong filename\n")
	}
	return nil
}

func main() {
	if len(os.Args) != 5 || os.Args[1] != "--old" || os.Args[3] != "--new" {
		log.Fatalf("%v", "Wrong input\n")
	}
	oldFile := getReader(os.Args[2]).parseFile()
	newFile := getReader(os.Args[4]).parseFile()

	for i := 0; i < len(newFile.Cake); i++ {
		check := false
		for j := 0; j < len(oldFile.Cake); j++ {
			if newFile.Cake[i].Name == oldFile.Cake[j].Name {
				check = true
				break
			}
		}
		if !check {
			fmt.Printf("ADDED cake \"%s\"\n", newFile.Cake[i].Name)
		}
	}
	for i := 0; i < len(oldFile.Cake); i++ {
		check := false
		for j := 0; j < len(newFile.Cake); j++ {
			if newFile.Cake[i].Name == oldFile.Cake[j].Name {
				check = true
				break
			}
		}
		if !check {
			fmt.Printf("REMOVED cake \"%s\"\n", oldFile.Cake[i].Name)
		}
	}
	for i := 0; i < len(newFile.Cake); i++ {
		for j := 0; j < len(oldFile.Cake); j++ {
			if newFile.Cake[i].Name == oldFile.Cake[j].Name {
				if newFile.Cake[i].Time != oldFile.Cake[j].Time {
					fmt.Printf("CHANGED cooking time for cake \"%s\" - \"%s\" instead of \"%s\"\n",
						newFile.Cake[i].Name, newFile.Cake[i].Time, oldFile.Cake[j].Time)
					break
				}
			}
		}
	}
	for i := 0; i < len(newFile.Cake); i++ {
		for j := 0; j < len(oldFile.Cake); j++ {
			if newFile.Cake[i].Name == oldFile.Cake[j].Name {
				for k := 0; k < len(newFile.Cake[i].Ingredients); k++ {
					check := false
					for z := 0; z < len(oldFile.Cake[j].Ingredients); z++ {
						if newFile.Cake[i].Ingredients[k].Itemname == oldFile.Cake[j].Ingredients[z].Itemname {
							check = true
							break
						}
					}
					if !check {
						fmt.Printf("ADDED ingredient \"%s\" for cake \"%s\"\n", newFile.Cake[i].Ingredients[k].Itemname, newFile.Cake[i].Name)
					}
				}
				break
			}
		}
	}
	for i := 0; i < len(newFile.Cake); i++ {
		for j := 0; j < len(oldFile.Cake); j++ {
			if newFile.Cake[i].Name == oldFile.Cake[j].Name {
				for k := 0; k < len(oldFile.Cake[i].Ingredients); k++ {
					check := false
					for z := 0; z < len(newFile.Cake[j].Ingredients); z++ {
						if newFile.Cake[i].Ingredients[z].Itemname == oldFile.Cake[j].Ingredients[k].Itemname {
							check = true
							break
						}
					}
					if !check {
						fmt.Printf("REMOVED ingredient \"%s\" for cake \"%s\"\n", oldFile.Cake[i].Ingredients[k].Itemname, oldFile.Cake[i].Name)
					}
				}
				break
			}
		}
	}
	for i := 0; i < len(newFile.Cake); i++ {
		for j := 0; j < len(oldFile.Cake); j++ {
			if newFile.Cake[i].Name == oldFile.Cake[j].Name {
				for k := 0; k < len(oldFile.Cake[i].Ingredients); k++ {
					for z := 0; z < len(newFile.Cake[j].Ingredients); z++ {
						if newFile.Cake[i].Ingredients[z].Itemname == oldFile.Cake[j].Ingredients[k].Itemname {
							if newFile.Cake[i].Ingredients[z].Itemcount != oldFile.Cake[j].Ingredients[k].Itemcount {
								fmt.Printf("CHANGED unit count for ingredient \"%s\" for cake \"%s\" - \"%s\" instead of \"%s\"\n",
									newFile.Cake[i].Ingredients[z].Itemname, newFile.Cake[i].Name, newFile.Cake[i].Ingredients[z].Itemcount,
									oldFile.Cake[j].Ingredients[k].Itemcount)
							}
							if newFile.Cake[i].Ingredients[z].Itemunit != oldFile.Cake[j].Ingredients[k].Itemunit {
								if newFile.Cake[i].Ingredients[z].Itemunit == "" {
									fmt.Printf("REMOVED unit \"%s\" for ingredient \"%s\" for cake \"%s\"\n",
										oldFile.Cake[j].Ingredients[k].Itemunit, oldFile.Cake[j].Ingredients[k].Itemname,
										oldFile.Cake[j].Name)
								} else if oldFile.Cake[j].Ingredients[k].Itemunit == "" {
									fmt.Printf("ADDED unit \"%s\" for ingredient \"%s\" for cake \"%s\"\n",
										newFile.Cake[i].Ingredients[z].Itemunit, oldFile.Cake[j].Ingredients[k].Itemname,
										oldFile.Cake[j].Name)
								} else {
									fmt.Printf("CHANGED unit for ingredient \"%s\" for cake \"%s\" - \"%s\" instead of \"%s\"\n",
										oldFile.Cake[j].Ingredients[k].Itemname, oldFile.Cake[j].Name, newFile.Cake[i].Ingredients[z].Itemunit,
										oldFile.Cake[j].Ingredients[k].Itemunit)
								}
							}
							break
						}
					}
				}
				break
			}
		}
	}
}
