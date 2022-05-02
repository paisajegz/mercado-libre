package xmen

import (
	"github.com/ymzuiku/hit"
)

func VerifyDna(dna *DNA, item string, x int, y int) bool {
	var verify [6]Verifycation
	var valYPos = (y + 3) < len(dna.DNA)
	var valXPos = (x + 3) < len(dna.DNA[y])
	var valXMin = (x - 3) >= 0
	var valYMin = (y - 3) >= 0

	for i, _ := range verify {
		verify[i].Item = item
		verify[i].Is_invalid = false
	}

	for i := 1; i <= 3; i++ {
		verify[0] = hit.If(valYPos, func() Verifycation { return VerifyCadena(verify[0], string(dna.DNA[y+i][x])) }, Verifycation{Is_invalid: true}).(Verifycation)
		verify[1] = hit.If(valXPos, func() Verifycation { return VerifyCadena(verify[1], string(dna.DNA[y][x+i])) }, Verifycation{Is_invalid: true}).(Verifycation)
		verify[2] = hit.If(valXPos && valYPos, func() Verifycation { return VerifyCadena(verify[2], string(dna.DNA[y+i][x+i])) }, Verifycation{Is_invalid: true}).(Verifycation)
		verify[3] = hit.If(valXMin && valYPos, func() Verifycation { return VerifyCadena(verify[3], string(dna.DNA[y+i][x-i])) }, Verifycation{Is_invalid: true}).(Verifycation)
		verify[4] = hit.If(valXPos && valYMin, func() Verifycation { return VerifyCadena(verify[4], string(dna.DNA[y-i][x+i])) }, Verifycation{Is_invalid: true}).(Verifycation)
		verify[5] = hit.If(valXMin && valYMin, func() Verifycation { return VerifyCadena(verify[5], string(dna.DNA[y-i][x-i])) }, Verifycation{Is_invalid: true}).(Verifycation)
	}

	for _, validate := range verify {
		if !(validate.Is_invalid) {
			return true
		}
	}
	return false
}

func VerifyCadena(validate Verifycation, item string) Verifycation {
	// validate.item = item
	if !(validate.Is_invalid) && validate.Item != item {
		validate.Is_invalid = true
	}
	return validate
}
