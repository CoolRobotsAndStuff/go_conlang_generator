package modifiers

// Adds a string to the end of the base.
type Suffix struct {
	Suffix string
}

func (suffixModiffier *Suffix) Apply(baseWord string) (newWord string) {
	return baseWord + suffixModiffier.Suffix
}

// Adds a string at the beggining of the base.
type Prefix struct{
	Prefix string
}

func (prefixModdiffier *Prefix) Apply(baseWord string) (newWord string) {
	return prefixModdiffier.Prefix + baseWord
}