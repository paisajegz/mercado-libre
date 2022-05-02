package test

import (
	"encoding/json"
	"io/ioutil"
	"log"
	api "mercado-level-1/internal"
	"mercado-level-1/internal/xmen"
	"testing"
)

// test function
func TestVerifyCadena(t *testing.T) {
	verification := new(xmen.Verifycation)
	verification.Item = "T"
	verification.Is_invalid = false
	*verification = xmen.VerifyCadena(*verification, "T")
	if verification.Is_invalid {
		t.Errorf("Expected String(%t) is not same as"+
			" actual string (%t)", false, verification.Is_invalid)
	}
}

func TestVerifyDna(t *testing.T) {
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
	isDna := xmen.VerifyDna(dna, "A", 0, 0)
	if !isDna {
		t.Errorf("Expected String(%t) is not same as"+
			" actual string (%t)", true, isDna)
	}
}

func TestVerifyDnaInvalid(t *testing.T) {
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
	isDna := xmen.VerifyDna(dna, "T", 0, 0)
	if isDna {
		t.Errorf("Expected String(%t) is not same as"+
			" actual string (%t)", true, isDna)
	}
}

func TestIsMutant(t *testing.T) {
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

	isMut := xmen.IsMutant(dna)
	if !isMut {
		t.Errorf("Expected String(%t) is not same as"+
			" actual string (%t)", true, isMut)
	}
}

func TestIsMutantInvalid(t *testing.T) {
	dna := new(xmen.DNA)
	file, err := ioutil.ReadFile("./data-invalid.json")
	if err != nil {
		log.Fatal(err)
	}
	text := string(file)
	errJS := json.Unmarshal([]byte(text), &dna)
	if errJS != nil {
		log.Fatal(errJS)
	}

	isMut := xmen.IsMutant(dna)
	if isMut {
		t.Errorf("Expected String(%t) is not same as"+
			" actual string (%t)", true, isMut)
	}
}

func TestValidateDna(t *testing.T) {
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
	validate := api.ValidateDna(dna.DNA)
	if !validate {
		t.Errorf("Expected String(%t) is not same as"+
			" actual string (%t)", true, validate)

	}
}

func TestValidateDnaInvalid(t *testing.T) {
	dna := new(xmen.DNA)
	file, err := ioutil.ReadFile("./data-invalidII-context.json")
	if err != nil {
		log.Fatal(err)
	}
	text := string(file)
	errJS := json.Unmarshal([]byte(text), &dna)
	if errJS != nil {
		log.Fatal(errJS)
	}
	validate := api.ValidateDna(dna.DNA)
	if validate {
		t.Errorf("Expected String(%t) is not same as"+
			" actual string (%t)", true, validate)

	}
}
