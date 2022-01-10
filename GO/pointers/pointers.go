package main

import "fmt"

type Creature struct {
	Species string
}

func main() {
	var creature string = "shark"
	var pointer *string = &creature

	// shark
	fmt.Println("Creature =", creature)
	// 0xc000010230
	fmt.Println("Pointer =", pointer)
	// shark
	fmt.Println("Pointer Value=", *pointer)

	*pointer = "jellyfish"
	// jellyfish
	fmt.Println("*pointer =", *pointer)

	var newCreature Creature = Creature{Species: "shark"}

	fmt.Printf("1) %+v\n", newCreature)
	changeCreature(&newCreature)
	fmt.Printf("3) %+v\n", newCreature)
}

func changeCreature(creature *Creature) {
	creature.Species = "jellyfish"
	fmt.Printf("2) %+v\n", creature)
}
