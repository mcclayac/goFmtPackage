package main

import (
	"fmt"
	"poetry"
)

func main() {



	//p1 := poetry.NewPoem()
	p2 := poetry.Poem{{"The mortal fruit upon the bough",
		"Hands above the nuptial bed.",
		"The cat-bird in the tree returns",
		"The forfeit of his mutual vow.",
	}, {"The hard, untimely apple of",
		"The branch that feeds on watered rain,",
		"Takes the place upon her lips",
		"Of her late lamented love."},
		{"Many hands together press,",
			"Shaped within a static prayer",
			"Recall to one the chorister",
			"Docile in his sexless dress.",
		}}

	/*
	fmt.Printf("%v\n\n", p2)   // best view option type
	fmt.Printf("%#v\n\n", p2)  // The type of object and the values
	fmt.Printf("%T\n\n", p2)  // The type of object no values
	fmt.Printf("%T\n\n", p2[0])  // The type of object no values
	fmt.Printf("%T\n\n", p2[0][0])  // The type of object no values
	fmt.Printf("%s\n\n", p2[0][0])  // The type of object no values
	fmt.Printf("%q\n\n", p2[0][0])  // The type of object no values
	fmt.Printf("% x\n\n", p2[0][0])  // The type of object no values
	fmt.Printf("=========================\n\n")
	*/

	fmt.Printf("=========================\n")
	fmt.Printf("%s\n\n", p2)  // The type of object no values
	fmt.Printf("=========================\n\n")


}
