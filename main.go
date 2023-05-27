package main

import (
	"fmt"
	"github.com/Uchijo/go-my-regex/eval"
)

func main() {
	regex := eval.MakeRegApp(
		[]eval.RegExp{
			eval.MakeRegChar("h"),
			eval.MakeRegChar("h"),
			eval.MakeRegChar("h"),
			eval.MakeRegChar("g"),
		},
	)
	input := "hhhg"
	fmt.Printf("regex.match(&input): %v\n", regex.Match(&input))
}
