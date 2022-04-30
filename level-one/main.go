package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"mercado-level-1/xmen"
)

func main() {
	dna := new(xmen.DNA)
	file, err := ioutil.ReadFile("./data.json")
	if err != nil {
		log.Fatal(err)
	}
	text := string(file)
	errJS := json.Unmarshal([]byte(text), &dna)
	if errJS != nil {
		log.Fatal(errJS)
	}

	if xmen.IsMutant(dna) {
		fmt.Println("is mutant")
	} else {
		fmt.Println("is not mutant")
	}
}
