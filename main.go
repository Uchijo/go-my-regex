package main

import (
	"fmt"
	"github.com/Uchijo/go-my-regex/eval"
)

func main() {
	regex := eval.MakeRegApp(
		[]eval.RegExp{
			eval.MakeRegChar("a"),
			eval.MakeRegStar(
				eval.MakeRegUnion(
					eval.MakeRegChar("a"),
					eval.MakeRegChar("b"),
				),
			),
		},
	)

	input := "aaaaaa"
	fmt.Printf("regex.Match(&input): %v\n", regex.Match(&input))
}
