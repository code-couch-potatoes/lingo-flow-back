package yadict

type Word struct {
	Src      string
	POS      string
	Ts       string
	Tr       []string
	Examples []Example
}

type Example struct {
	Phrase string
	Tr     string
}
