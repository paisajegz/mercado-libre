package xmen

type DNA struct {
	DNA []string `json:"dna"`
}

type Verifycation struct {
	is_invalid bool
	item       string
}
