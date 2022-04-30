package xmen

func IsMutant(dna *DNA) bool {
	for y, cadena := range dna.DNA {
		for x, item := range cadena {
			if verifyDna(dna, string(item), x, y) {
				return true
			}

		}
	}
	return false
}
