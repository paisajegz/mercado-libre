package xmen

type DNA struct {
	DNA []string `json:"dna" validate:"required,is-dna"`
}

type Verifycation struct {
	is_invalid bool
	item       string
}
