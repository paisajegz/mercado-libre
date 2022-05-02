package xmen

type DNA struct {
	DNA []string `json:"dna" validate:"required,is-dna"`
}

type Verifycation struct {
	Is_invalid bool
	Item       string
}

type ResponseStat struct {
	CountHuman  int     `json:"count_human_dna"`
	CountNutant int     `json:"count_mutant_dna"`
	Ratio       float64 `json:"ratio"`
}
