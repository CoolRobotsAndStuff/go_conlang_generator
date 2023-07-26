package modifiers

// Modifies a base word.
type Modifier interface {
	Apply(baseWord string) (newWord string)
}