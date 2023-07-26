package word

import "github.com/CoolRobotsAndStuff/go_conlang_generator/modifiers"

type Word struct {
	Base      string
	Modifiers []modifiers.Modifier
}

func (word Word) ApplyModifiers() string {
	finalWord := word.Base
	for _, mod := range word.Modifiers {
		finalWord = mod.Apply(finalWord)
	}
	return finalWord
}
