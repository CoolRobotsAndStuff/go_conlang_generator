package main

import (
	"fmt"

	"github.com/CoolRobotsAndStuff/go_conlang_generator/modifiers"
	"github.com/CoolRobotsAndStuff/go_conlang_generator/word"
)


func main() {
	my_modifiers := []modifiers.Modifier{&modifiers.Suffix{Suffix: "o"}}
	my_word := word.Word{Base: "gat", Modifiers: my_modifiers}

	fmt.Printf("Word: %v \n", my_word.ApplyModifiers())


}