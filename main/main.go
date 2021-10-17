package main

import (
	"fmt"

	"github.com/rafaelpapastamatiou/iniciando-go-lang/blankidentifier"
	"github.com/rafaelpapastamatiou/iniciando-go-lang/conditionals"
	"github.com/rafaelpapastamatiou/iniciando-go-lang/functions"
	"github.com/rafaelpapastamatiou/iniciando-go-lang/pointers"
	"github.com/rafaelpapastamatiou/iniciando-go-lang/repeatloops"
	"github.com/rafaelpapastamatiou/iniciando-go-lang/scope"
	switchcondition "github.com/rafaelpapastamatiou/iniciando-go-lang/switch"
	"github.com/rafaelpapastamatiou/iniciando-go-lang/variables"
	"github.com/rafaelpapastamatiou/iniciando-go-lang/visibility"
)

func main() {
	variables.Describe()

	scope.Describe()

	fmt.Println("\n\n\nVisibility")
	visibility.VisibleFunc()
	// invisible visibility.invisibleFunc()

	blankidentifier.Describe()

	pointers.Describe()

	conditionals.Describe()

	repeatloops.Describe()

	switchcondition.Describe()

	functions.Describe()
}
